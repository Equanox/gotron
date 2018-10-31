// +build prod

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// FrontendPath Sets path to call electron in production environment
func FrontendPath() (dir string, exe string, args string) {
	fmt.Println("production mode")

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	return exPath + "/electron/", "electron-rampup", ""
}
