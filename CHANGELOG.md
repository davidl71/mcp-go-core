# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial project structure
- Project root detection (`pkg/mcp/security/GetProjectRoot()`)
- Path validation utilities (`pkg/mcp/security/ValidatePath()`)
- Basic Makefile with common targets
- GitHub Actions CI workflow
- Project documentation (README, verification guide)

### Changed
- N/A (initial release)

### Deprecated
- N/A (initial release)

### Removed
- N/A (initial release)

### Fixed
- N/A (initial release)

### Security
- Path validation prevents directory traversal attacks
- Project root detection ensures operations stay within project boundaries
