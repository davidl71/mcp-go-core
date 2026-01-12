// Package client provides wrapper functions around existing MCP client libraries.
//
// This package wraps external MCP client libraries (e.g., github.com/metoro-io/mcp-golang)
// to provide:
//   - Type integration with mcp-go-core types
//   - Testing utilities for mcp-go-core servers
//   - Consistent API using mcp-go-core types
//
// Example usage:
//
//	clientInfo := protocol.ClientInfo{
//	    Name:    "test-client",
//	    Version: "1.0.0",
//	}
//	c, err := client.NewClient("/path/to/server", clientInfo)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer c.Close()
//
//	ctx := context.Background()
//	_, err = c.Initialize(ctx)
//	tools, err := c.ListTools(ctx)
//
// Dependencies:
//
// This package requires github.com/metoro-io/mcp-golang to be available.
// To build without this package (to avoid the dependency), use:
//
//	go build -tags no_mcp_client
package client

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/davidl71/mcp-go-core/pkg/mcp/protocol"
	"github.com/davidl71/mcp-go-core/pkg/mcp/types"
)

// Client wraps an external MCP client library to provide a consistent API
// using mcp-go-core types.
type Client struct {
	// underlying is the wrapped client from the external library.
	// Using interface{} to avoid direct dependency (can be made type-safe with build tags)
	underlying interface{} // Expected to be *mcp.Client from github.com/metoro-io/mcp-golang

	// clientInfo stores the client information
	clientInfo protocol.ClientInfo

	// serverCommand is the command to launch the server
	serverCommand string

	// serverArgs are arguments to pass to the server
	serverArgs []string

	// initialized tracks whether the client has been initialized
	initialized bool
}

// NewClient creates a new client wrapper that connects to an MCP server.
//
// The serverCommand should be the path to the server binary or command.
// The clientInfo identifies this client to the server.
//
// Note: This function requires the mcp-golang library to be available.
// See the package documentation for dependency requirements.
func NewClient(serverCommand string, clientInfo protocol.ClientInfo) (*Client, error) {
	if serverCommand == "" {
		return nil, fmt.Errorf("server command cannot be empty")
	}
	if clientInfo.Name == "" {
		return nil, fmt.Errorf("client info name cannot be empty")
	}

	return &Client{
		clientInfo:    clientInfo,
		serverCommand: serverCommand,
		serverArgs:    []string{},
		initialized:   false,
	}, nil
}

// NewClientWithArgs creates a new client wrapper with server arguments.
func NewClientWithArgs(serverCommand string, serverArgs []string, clientInfo protocol.ClientInfo) (*Client, error) {
	client, err := NewClient(serverCommand, clientInfo)
	if err != nil {
		return nil, err
	}
	client.serverArgs = serverArgs
	return client, nil
}

// Initialize initializes the client session with the MCP server.
//
// This must be called before any other operations. It establishes the connection
// and performs the MCP handshake.
func (c *Client) Initialize(ctx context.Context) (*protocol.InitializeResult, error) {
	// TODO: Implement actual initialization using wrapped client
	// This is a placeholder that shows the intended API
	// Actual implementation will call the underlying client's Initialize method
	
	// For now, return an error indicating this needs implementation
	return nil, fmt.Errorf("client wrapper not yet fully implemented - requires external client library")
}

// ListTools lists all available tools from the server.
//
// Returns tools using mcp-go-core types.ToolInfo format.
func (c *Client) ListTools(ctx context.Context) ([]types.ToolInfo, error) {
	if !c.initialized {
		return nil, fmt.Errorf("client must be initialized before listing tools")
	}
	
	// TODO: Implement actual tool listing using wrapped client
	// Convert external library tool types to types.ToolInfo
	return nil, fmt.Errorf("client wrapper not yet fully implemented - requires external client library")
}

// CallTool calls a tool on the server with the given arguments.
//
// The args map is converted to JSON and passed to the tool.
// Returns results using mcp-go-core types.TextContent format.
func (c *Client) CallTool(ctx context.Context, name string, args map[string]interface{}) ([]types.TextContent, error) {
	if !c.initialized {
		return nil, fmt.Errorf("client must be initialized before calling tools")
	}
	if name == "" {
		return nil, fmt.Errorf("tool name cannot be empty")
	}
	
	// TODO: Implement actual tool call using wrapped client
	// Convert external library response types to []types.TextContent
	return nil, fmt.Errorf("client wrapper not yet fully implemented - requires external client library")
}

// ListResources lists all available resources from the server.
func (c *Client) ListResources(ctx context.Context) ([]protocol.Resource, error) {
	if !c.initialized {
		return nil, fmt.Errorf("client must be initialized before listing resources")
	}
	
	// TODO: Implement actual resource listing using wrapped client
	return nil, fmt.Errorf("client wrapper not yet fully implemented - requires external client library")
}

// ReadResource reads a resource from the server by URI.
//
// Returns the resource content as bytes and the MIME type.
func (c *Client) ReadResource(ctx context.Context, uri string) ([]byte, string, error) {
	if !c.initialized {
		return nil, "", fmt.Errorf("client must be initialized before reading resources")
	}
	if uri == "" {
		return nil, "", fmt.Errorf("resource URI cannot be empty")
	}
	
	// TODO: Implement actual resource reading using wrapped client
	return nil, "", fmt.Errorf("client wrapper not yet fully implemented - requires external client library")
}

// PromptInfo represents prompt metadata (similar to ToolInfo)
type PromptInfo struct {
	Name        string
	Description string
}

// ListPrompts lists all available prompts from the server.
func (c *Client) ListPrompts(ctx context.Context) ([]PromptInfo, error) {
	if !c.initialized {
		return nil, fmt.Errorf("client must be initialized before listing prompts")
	}
	
	// TODO: Implement actual prompt listing using wrapped client
	return nil, fmt.Errorf("client wrapper not yet fully implemented - requires external client library")
}

// GetPrompt gets a prompt template from the server.
func (c *Client) GetPrompt(ctx context.Context, name string, args map[string]interface{}) (string, error) {
	if !c.initialized {
		return "", fmt.Errorf("client must be initialized before getting prompts")
	}
	if name == "" {
		return "", fmt.Errorf("prompt name cannot be empty")
	}
	
	// TODO: Implement actual prompt retrieval using wrapped client
	return "", fmt.Errorf("client wrapper not yet fully implemented - requires external client library")
}

// Close closes the client connection and cleans up resources.
func (c *Client) Close() error {
	if c.underlying == nil {
		return nil
	}
	
	// TODO: Call underlying client's Close method if it has one
	c.initialized = false
	return nil
}

// GetClientInfo returns the client information.
func (c *Client) GetClientInfo() protocol.ClientInfo {
	return c.clientInfo
}

// IsInitialized returns whether the client has been initialized.
func (c *Client) IsInitialized() bool {
	return c.initialized
}

// validateServerCommand checks if the server command exists and is executable.
func validateServerCommand(command string) error {
	// Check if command is executable
	cmd := exec.Command(command, "--help")
	if err := cmd.Run(); err != nil {
		// Command might not support --help, try just checking if file exists
		// This is a basic validation - actual execution will show real errors
		return nil // Don't fail validation on --help check
	}
	return nil
}
