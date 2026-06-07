// internal/commands/logger.go
package commands

import (
	"fmt"
	"os"
)

var isVerbose bool

// SetVerbose enables or disables verbose logging
func SetVerbose(enabled bool) {
	isVerbose = enabled
}

// Logf prints a message if verbose mode is enabled
func Logf(format string, args ...interface{}) {
	if isVerbose {
		fmt.Printf("[DEBUG] "+format+"\n", args...)
	}
}

// Log prints a message if verbose mode is enabled
func Log(msg string) {
	if isVerbose {
		fmt.Printf("[DEBUG] %s\n", msg)
	}
}

// CheckFileExists checks if a file exists and returns true if it does
func CheckFileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// CheckDirectoryExists checks if a directory exists and returns true if it does
func CheckDirectoryExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}
