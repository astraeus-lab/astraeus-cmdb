package util

import (
	"errors"
	"os"
)

// IsFileExist check whether the file exists.
func IsFileExist(path string) bool {
	_, err := os.Stat(path)

	return !errors.Is(err, os.ErrNotExist)
}

// IsFileExistOrCreate check whether the file exists.
// If the file does not exist, it is created with mode 0666 (before umask).
func IsFileExistOrCreate(path string) error {
	if IsFileExist(path) {
		return nil
	}

	if _, err := os.Create(path); err != nil {
		return err
	}

	return nil
}

// IsFilePermission check whether the file has permission.
func IsFilePermission(path string) bool {
	_, err := os.Stat(path)

	return !errors.Is(err, os.ErrPermission)
}
