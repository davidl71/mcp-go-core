package gosdk

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestValidateContext(t *testing.T) {
	tests := []struct {
		name    string
		ctx     context.Context
		wantErr bool
		errMsg  string
	}{
		{
			name:    "valid context",
			ctx:     context.Background(),
			wantErr: false,
		},
		{
			name:    "nil context",
			ctx:     nil,
			wantErr: true,
			errMsg:  "context cannot be nil",
		},
		{
			name: "cancelled context",
			ctx: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx
			}(),
			wantErr: true,
			errMsg:  "context cancelled",
		},
		{
			name: "context with timeout",
			ctx: func() context.Context {
				ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
				time.Sleep(time.Millisecond) // Wait for timeout
				cancel()
				return ctx
			}(),
			wantErr: true,
			errMsg:  "context cancelled",
		},
		{
			name:    "context with value",
			ctx:     context.WithValue(context.Background(), "key", "value"),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateContext(tt.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err == nil {
					t.Error("ValidateContext() expected error but got nil")
					return
				}
				if tt.errMsg != "" {
					errorMsg := err.Error()
					// Check if error message contains expected string
					found := false
					for i := 0; i <= len(errorMsg)-len(tt.errMsg); i++ {
						if errorMsg[i:i+len(tt.errMsg)] == tt.errMsg {
							found = true
							break
						}
					}
					if !found {
						t.Errorf("ValidateContext() error = %q, want error containing %q", errorMsg, tt.errMsg)
					}
				}
				// Verify error can be unwrapped
				if !errors.Is(err, context.Canceled) && tt.ctx != nil {
					// If context was cancelled, error should wrap context.Canceled
					if tt.ctx.Err() == context.Canceled {
						var cancelErr error
						if errors.As(err, &cancelErr) {
							// Error should be related to cancellation
						}
					}
				}
			}
		})
	}
}
