// Package platform provides platform detection and utilities for MCP servers.
//
// This package provides utilities for detecting the operating system, architecture,
// and platform-specific path handling. This is useful for cross-platform compatibility
// and platform-specific optimizations.
//
// Example usage:
//
//	// Detect platform
//	os := platform.OS()
//	arch := platform.Architecture()
//
//	// Platform-specific path handling
//	path := platform.NormalizePath("/path/to/file")
//
//	// Check platform
//	if platform.IsWindows() {
//		// Windows-specific code
//	}
package platform

import (
	"runtime"
	"strings"
)

// OS represents the operating system type
type OS string

const (
	// OSWindows represents Microsoft Windows
	OSWindows OS = "windows"
	// OSLinux represents Linux
	OSLinux OS = "linux"
	// OSDarwin represents macOS (Darwin)
	OSDarwin OS = "darwin"
	// OSUnknown represents an unknown operating system
	OSUnknown OS = "unknown"
)

// Architecture represents the CPU architecture
type Architecture string

const (
	// ArchAMD64 represents x86-64 (64-bit Intel/AMD)
	ArchAMD64 Architecture = "amd64"
	// ArchARM64 represents ARM64 (64-bit ARM)
	ArchARM64 Architecture = "arm64"
	// Arch386 represents x86 (32-bit Intel/AMD)
	Arch386 Architecture = "386"
	// ArchARM represents ARM (32-bit ARM)
	ArchARM Architecture = "arm"
	// ArchUnknown represents an unknown architecture
	ArchUnknown Architecture = "unknown"
)

// PlatformInfo contains information about the current platform
type PlatformInfo struct {
	OS           OS
	Architecture Architecture
	GOOS         string
	GOARCH       string
}

// Detect returns the current platform information
func Detect() *PlatformInfo {
	return &PlatformInfo{
		OS:           OS(runtime.GOOS),
		Architecture: Architecture(runtime.GOARCH),
		GOOS:         runtime.GOOS,
		GOARCH:       runtime.GOARCH,
	}
}

// OS returns the current operating system
func OS() OS {
	goos := runtime.GOOS
	switch goos {
	case "windows":
		return OSWindows
	case "linux":
		return OSLinux
	case "darwin":
		return OSDarwin
	default:
		return OSUnknown
	}
}

// Architecture returns the current CPU architecture
func Architecture() Architecture {
	goarch := runtime.GOARCH
	switch goarch {
	case "amd64", "x86_64":
		return ArchAMD64
	case "arm64", "aarch64":
		return ArchARM64
	case "386", "i386":
		return Arch386
	case "arm":
		return ArchARM
	default:
		return ArchUnknown
	}
}

// IsWindows returns true if running on Windows
func IsWindows() bool {
	return OS() == OSWindows
}

// IsLinux returns true if running on Linux
func IsLinux() bool {
	return OS() == OSLinux
}

// IsDarwin returns true if running on macOS (Darwin)
func IsDarwin() bool {
	return OS() == OSDarwin
}

// IsUnix returns true if running on a Unix-like system (Linux, macOS, etc.)
func IsUnix() bool {
	return IsLinux() || IsDarwin()
}

// Is64Bit returns true if running on a 64-bit architecture
func Is64Bit() bool {
	arch := Architecture()
	return arch == ArchAMD64 || arch == ArchARM64
}

// Is32Bit returns true if running on a 32-bit architecture
func Is32Bit() bool {
	arch := Architecture()
	return arch == Arch386 || arch == ArchARM
}

// NormalizePath normalizes a path for the current platform.
// On Windows, converts forward slashes to backslashes.
// On Unix-like systems, ensures forward slashes.
func NormalizePath(path string) string {
	if IsWindows() {
		// Windows uses backslashes, but Go's filepath package handles this
		// We'll just ensure consistent forward slashes for cross-platform compatibility
		// The filepath package will handle the conversion when needed
		return strings.ReplaceAll(path, "\\", "/")
	}
	// Unix-like systems use forward slashes
	return path
}

// PathSeparator returns the path separator for the current platform
func PathSeparator() string {
	if IsWindows() {
		return "\\"
	}
	return "/"
}

// PathListSeparator returns the path list separator for the current platform
// (used in PATH environment variable)
func PathListSeparator() string {
	if IsWindows() {
		return ";"
	}
	return ":"
}

// String returns a string representation of the platform
func (p *PlatformInfo) String() string {
	return string(p.OS) + "/" + string(p.Architecture)
}

// IsCompatible checks if the current platform is compatible with the given OS and architecture
func (p *PlatformInfo) IsCompatible(os OS, arch Architecture) bool {
	return p.OS == os && p.Architecture == arch
}
