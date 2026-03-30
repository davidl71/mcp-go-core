//go:build windows

package filelock

import (
	"fmt"
	"syscall"

	"golang.org/x/sys/windows"
)

func (fl *FileLock) tryLockOS() error {
	var ol windows.Overlapped
	err := windows.LockFileEx(
		windows.Handle(fl.file.Fd()),
		windows.LOCKFILE_EXCLUSIVE_LOCK|windows.LOCKFILE_FAIL_IMMEDIATELY,
		0,
		1,
		0,
		&ol,
	)
	if err != nil {
		// Treat common “already locked” cases consistently with unix.
		if err == syscall.ERROR_LOCK_VIOLATION {
			return fmt.Errorf("lock is held by another process")
		}
		return fmt.Errorf("acquire lock: %w", err)
	}
	return nil
}

func (fl *FileLock) unlockOS() error {
	var ol windows.Overlapped
	if err := windows.UnlockFileEx(windows.Handle(fl.file.Fd()), 0, 1, 0, &ol); err != nil {
		return fmt.Errorf("release lock: %w", err)
	}
	return nil
}

