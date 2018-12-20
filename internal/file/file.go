package file

import (
	"os"
)

// Exists return true when a file exists, false otherwise.
func Exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
