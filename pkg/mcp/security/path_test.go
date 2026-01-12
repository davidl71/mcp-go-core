package security

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetProjectRoot(t *testing.T) {
	// Get the actual project root (where this test is running)
	expectedRoot, err := filepath.Abs("../../..")
	if err != nil {
		t.Fatalf("Failed to get expected root: %v", err)
	}

	// Test from current directory
	root, err := GetProjectRoot(".")
	if err != nil {
		t.Fatalf("GetProjectRoot failed: %v", err)
	}

	// Normalize paths for comparison
	expectedRoot = filepath.Clean(expectedRoot)
	root = filepath.Clean(root)

	if root != expectedRoot {
		t.Errorf("GetProjectRoot() = %v, want %v", root, expectedRoot)
	}
}

func TestGetProjectRootFromSubdirectory(t *testing.T) {
	// Test from a subdirectory
	root, err := GetProjectRoot("pkg/mcp/security")
	if err != nil {
		t.Fatalf("GetProjectRoot failed: %v", err)
	}

	// Should still find the project root
	expectedRoot, _ := filepath.Abs("../../..")
	expectedRoot = filepath.Clean(expectedRoot)
	root = filepath.Clean(root)

	if root != expectedRoot {
		t.Errorf("GetProjectRoot() = %v, want %v", root, expectedRoot)
	}
}

func TestValidatePath(t *testing.T) {
	// Create a temporary directory structure
	tmpDir, err := os.MkdirTemp("", "mcp-go-core-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a subdirectory
	subDir := filepath.Join(tmpDir, "subdir")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("Failed to create subdir: %v", err)
	}

	tests := []struct {
		name        string
		path        string
		projectRoot string
		wantErr     bool
	}{
		{
			name:        "valid relative path",
			path:        "subdir",
			projectRoot: tmpDir,
			wantErr:     false,
		},
		{
			name:        "valid absolute path within root",
			path:        subDir,
			projectRoot: tmpDir,
			wantErr:     false,
		},
		{
			name:        "directory traversal attempt",
			path:        "../outside",
			projectRoot: tmpDir,
			wantErr:     true,
		},
		{
			name:        "empty project root",
			path:        "subdir",
			projectRoot: "",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ValidatePath(tt.path, tt.projectRoot)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidatePathExists(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "mcp-go-core-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a file
	testFile := filepath.Join(tmpDir, "test.txt")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Test existing path
	_, err = ValidatePathExists("test.txt", tmpDir)
	if err != nil {
		t.Errorf("ValidatePathExists() for existing file error = %v", err)
	}

	// Test non-existing path
	_, err = ValidatePathExists("nonexistent.txt", tmpDir)
	if err == nil {
		t.Error("ValidatePathExists() for non-existing file should return error")
	}
}
