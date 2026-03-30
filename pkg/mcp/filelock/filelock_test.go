package filelock

import (
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

func TestNew_CreatesParentDirAndFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "locks", "a.lock")

	fl, err := New(path, 0)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	defer fl.Close()

	if fl.Path() != path {
		t.Fatalf("Path() = %q, want %q", fl.Path(), path)
	}
}

func TestLock_Timeout(t *testing.T) {
	// Windows and Darwin semantics vary for same-process locking behavior;
	// keep this as a simple “timeout returns error” test on Linux only.
	if runtime.GOOS != "linux" {
		t.Skip("contention timeout test is linux-only")
	}

	dir := t.TempDir()
	path := filepath.Join(dir, "a.lock")

	holder, err := New(path, 0)
	if err != nil {
		t.Fatalf("New(holder): %v", err)
	}
	defer holder.Close()
	if err := holder.Lock(); err != nil {
		t.Fatalf("holder.Lock: %v", err)
	}

	contender, err := New(path, 50*time.Millisecond)
	if err != nil {
		t.Fatalf("New(contender): %v", err)
	}
	defer contender.Close()

	if err := contender.Lock(); err == nil {
		t.Fatalf("contender.Lock: expected timeout error")
	}
}

