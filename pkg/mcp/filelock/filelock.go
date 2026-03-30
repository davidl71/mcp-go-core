// Package filelock provides simple cross-process file-based locking.
//
// This is intentionally small and dependency-light so MCP tools and CLIs can
// coordinate on repo-level locks (e.g. “only one sync at a time”).
package filelock

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// FileLock provides cross-process file-based locking using OS-level locks.
type FileLock struct {
	file    *os.File
	path    string
	timeout time.Duration
}

// New creates a new file lock at path. The lock file is created if needed.
func New(path string, timeout time.Duration) (*FileLock, error) {
	if path == "" {
		return nil, fmt.Errorf("lock path is empty")
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, fmt.Errorf("create lock directory: %w", err)
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0o644)
	if err != nil {
		return nil, fmt.Errorf("open lock file: %w", err)
	}

	return &FileLock{file: f, path: path, timeout: timeout}, nil
}

// Path returns the lock file path.
func (fl *FileLock) Path() string {
	if fl == nil {
		return ""
	}
	return fl.path
}

// TryLock attempts to acquire the lock non-blocking.
func (fl *FileLock) TryLock() error {
	if fl == nil || fl.file == nil {
		return fmt.Errorf("lock file not open")
	}
	return fl.tryLockOS()
}

// Lock acquires the lock, blocking until available or timeout.
func (fl *FileLock) Lock() error {
	if fl == nil {
		return fmt.Errorf("lock is nil")
	}

	if fl.timeout == 0 {
		return fl.TryLock()
	}

	start := time.Now()
	pollInterval := 50 * time.Millisecond
	for {
		if err := fl.TryLock(); err == nil {
			return nil
		}
		if time.Since(start) >= fl.timeout {
			return fmt.Errorf("lock timeout after %v", fl.timeout)
		}
		time.Sleep(pollInterval)
	}
}

// Unlock releases the lock.
func (fl *FileLock) Unlock() error {
	if fl == nil || fl.file == nil {
		return nil
	}
	return fl.unlockOS()
}

// Close closes the lock file and attempts to release the lock first.
func (fl *FileLock) Close() error {
	if fl == nil || fl.file == nil {
		return nil
	}
	_ = fl.Unlock()
	err := fl.file.Close()
	fl.file = nil
	return err
}

