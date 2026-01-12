package framework

import (
	"context"
	"testing"
)

func TestStdioTransport(t *testing.T) {
	transport := &StdioTransport{}

	// Test Type
	if transport.Type() != "stdio" {
		t.Errorf("StdioTransport.Type() = %q, want %q", transport.Type(), "stdio")
	}

	// Test Start
	ctx := context.Background()
	if err := transport.Start(ctx); err != nil {
		t.Errorf("StdioTransport.Start() error = %v, want nil", err)
	}

	// Test Stop
	if err := transport.Stop(ctx); err != nil {
		t.Errorf("StdioTransport.Stop() error = %v, want nil", err)
	}
}

func TestSSETransport(t *testing.T) {
	transport := &SSETransport{}

	// Test Type
	if transport.Type() != "sse" {
		t.Errorf("SSETransport.Type() = %q, want %q", transport.Type(), "sse")
	}

	// Test Start
	ctx := context.Background()
	if err := transport.Start(ctx); err != nil {
		t.Errorf("SSETransport.Start() error = %v, want nil", err)
	}

	// Test Stop
	if err := transport.Stop(ctx); err != nil {
		t.Errorf("SSETransport.Stop() error = %v, want nil", err)
	}
}

func TestTransportInterface(t *testing.T) {
	// Verify that both transports implement the Transport interface
	var _ Transport = (*StdioTransport)(nil)
	var _ Transport = (*SSETransport)(nil)
}

func TestStdioTransport_WithCancelledContext(t *testing.T) {
	transport := &StdioTransport{}

	// Create cancelled context
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	// Start should still work (stdio doesn't check context)
	if err := transport.Start(ctx); err != nil {
		t.Errorf("StdioTransport.Start() with cancelled context error = %v, want nil", err)
	}

	// Stop should still work
	if err := transport.Stop(ctx); err != nil {
		t.Errorf("StdioTransport.Stop() with cancelled context error = %v, want nil", err)
	}
}
