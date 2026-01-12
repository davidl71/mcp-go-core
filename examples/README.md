# MCP-Go-Core Examples

This directory contains example implementations using the `mcp-go-core` library.

## Examples

### 1. Basic Server (`basic_server/`)

A simple MCP server demonstrating:
- Server creation using factory pattern
- Tool registration (echo, math)
- Prompt registration
- Resource registration
- CLI/MCP dual mode support
- Stdio transport

**Run as MCP server:**
```bash
cd examples/basic_server
go run main.go
```

**Run as CLI:**
```bash
cd examples/basic_server
go run main.go list
go run main.go call echo --args '{"message":"Hello, World!"}'
go run main.go call math --args '{"operation":"add","a":10,"b":20}'
```

### 2. Advanced Server (`advanced_server/`)

An advanced MCP server demonstrating:
- Custom logger integration
- Middleware for request/response logging
- Adapter options pattern
- Error handling
- Performance monitoring

**Run:**
```bash
cd examples/advanced_server
go run main.go
```

## Building Examples

```bash
# Build all examples
make build-examples

# Build specific example
cd examples/basic_server
go build -o basic_server
```

## Using with Cursor

To use an example server with Cursor, add to `.cursor/mcp.json`:

```json
{
  "mcpServers": {
    "example-server": {
      "command": "/path/to/examples/basic_server/basic_server"
    }
  }
}
```

## Next Steps

1. Run the basic server example
2. Try the advanced server example
3. Modify examples to suit your needs
4. Create your own MCP server based on the examples
