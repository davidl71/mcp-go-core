# mcp-go-core

**Core MCP Infrastructure Library for Go**

A shared library containing common MCP (Model Context Protocol) infrastructure code extracted from `exarp-go` and `devwisdom-go` projects.

## Purpose

This library provides reusable components for building MCP servers in Go, including:

- Framework abstraction layer
- Security utilities (path validation, project root detection)
- Structured logging
- CLI utilities (TTY detection)
- JSON-RPC protocol types
- Common types and interfaces
- Platform detection
- Build system patterns
- CI/CD workflow templates

## Installation

```bash
go get github.com/davidl71/mcp-go-core
```

## Usage

```go
import (
    "github.com/davidl71/mcp-go-core/pkg/mcp/framework"
    "github.com/davidl71/mcp-go-core/pkg/mcp/security"
    "github.com/davidl71/mcp-go-core/pkg/mcp/logging"
)
```

## Project Structure

```
mcp-go-core/
├── pkg/
│   └── mcp/
│       ├── framework/     # Framework abstraction
│       ├── cli/           # CLI utilities
│       ├── config/        # Base configuration
│       ├── security/      # Security utilities
│       ├── logging/       # Structured logging
│       ├── protocol/      # JSON-RPC types
│       ├── platform/      # Platform detection
│       └── types/         # Common types
├── scripts/
│   └── dev-common.sh     # Shared development utilities
├── .github/
│   └── workflows/
│       └── common-go.yml  # Reusable CI/CD workflow
├── docs/
│   └── TEMPLATES/         # Documentation templates
├── Makefile.common        # Shared Makefile patterns
├── go.mod
└── README.md
```

## Status

✅ **v0.1.0 Released** - Core components extracted and published

This library contains shared MCP infrastructure code extracted from `exarp-go` and `devwisdom-go` projects. All high-priority components have been extracted, tested, and refactored.

## Contributing

This library is extracted from existing projects. See the source projects for contribution guidelines:

- [exarp-go](https://github.com/davidl71/exarp-go)
- [devwisdom-go](https://github.com/davidl71/devwisdom-go)

## License

TBD - Will match source project licenses
