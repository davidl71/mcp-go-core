package gosdk

import (
	"testing"

	"github.com/davidl71/mcp-go-core/pkg/mcp/logging"
)

func TestNewGoSDKAdapter_WithOptions(t *testing.T) {
	// Test that adapter can be created with options (even if they're placeholders)
	adapter := NewGoSDKAdapter("test-server", "1.0.0", WithLogger(nil), WithMiddleware(nil))

	if adapter == nil {
		t.Fatal("NewGoSDKAdapter() returned nil")
	}
	if adapter.GetName() != "test-server" {
		t.Errorf("adapter.GetName() = %q, want %q", adapter.GetName(), "test-server")
	}
}

func TestNewGoSDKAdapter_WithoutOptions(t *testing.T) {
	// Test backward compatibility - adapter should work without options
	adapter := NewGoSDKAdapter("test-server", "1.0.0")

	if adapter == nil {
		t.Fatal("NewGoSDKAdapter() returned nil")
	}
	if adapter.GetName() != "test-server" {
		t.Errorf("adapter.GetName() = %q, want %q", adapter.GetName(), "test-server")
	}
}

func TestAdapterOption_WithLogger(t *testing.T) {
	// Test with nil logger (should use default)
	adapter1 := NewGoSDKAdapter("test", "1.0.0", WithLogger(nil))
	if adapter1 == nil {
		t.Fatal("NewGoSDKAdapter() returned nil")
	}
	if adapter1.logger == nil {
		t.Error("Expected default logger to be created")
	}

	// Test with custom logger
	customLogger := logging.NewLogger()
	customLogger.SetLevel(logging.LevelDebug)
	adapter2 := NewGoSDKAdapter("test", "1.0.0", WithLogger(customLogger))
	if adapter2 == nil {
		t.Fatal("NewGoSDKAdapter() returned nil")
	}
	if adapter2.logger != customLogger {
		t.Error("Expected custom logger to be set")
	}
	if adapter2.logger != customLogger {
		t.Error("Logger should be the same instance")
	}
}

func TestAdapterOption_WithMiddleware(t *testing.T) {
	// Test that WithMiddleware option can be applied (even if it's a placeholder)
	adapter := NewGoSDKAdapter("test", "1.0.0", WithMiddleware(nil))

	if adapter == nil {
		t.Fatal("NewGoSDKAdapter() returned nil")
	}
	// Middleware integration is not yet implemented, so we just verify it doesn't crash
}
