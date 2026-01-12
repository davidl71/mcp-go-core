package gosdk

import "github.com/davidl71/mcp-go-core/pkg/mcp/logging"

// AdapterOption configures a GoSDKAdapter
type AdapterOption func(*GoSDKAdapter)

// WithLogger sets a custom logger for the adapter
// If not provided, a default logger will be created.
// The logger is used for tool registration, tool calls, and other adapter operations.
func WithLogger(logger *logging.Logger) AdapterOption {
	return func(a *GoSDKAdapter) {
		if logger != nil {
			a.logger = logger
		}
	}
}

// WithMiddleware adds middleware to the adapter
// This is a placeholder for future middleware support
func WithMiddleware(middleware interface{}) AdapterOption {
	return func(a *GoSDKAdapter) {
		// TODO: Implement middleware support
		// This would allow users to add request/response middleware
	}
}
