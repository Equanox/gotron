// +build !prod

package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// FrontendPath Sets path to call electron in development environment
func FrontendPath() (dir string, exe string, args string, err error) {
	fmt.Println("development mode")
	switch runtime.GOOS {
	case "darwin":
		appDirectory, err := filepath.Abs("app")
		if err != nil {
			return "", "", "", err
		}
		return "./app/node_modules/electron/dist/", "Electron.app", appDirectory, nil
	}

	return "./app/node_modules/electron/dist/", "electron", "app", nil
}
