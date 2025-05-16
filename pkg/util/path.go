package util

import (
    "path/filepath"
    "runtime"
)

// GetProjectRoot returns the absolute path to the project root directory.
// It works regardless of how or where the test is run.
func GetProjectRoot() string {
    _, filename, _, ok := runtime.Caller(0)
    if !ok {
        panic("failed to get caller info")
    }

    // filename /path/to/project/internal/test/testutil.go
    return filepath.Join(filepath.Dir(filename), "..", "..")
}
