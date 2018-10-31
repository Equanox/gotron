// +build !prod

package main

import (
	"fmt"
	"runtime"
)

// FrontendPath Sets path to call electron in development environment
func FrontendPath() (dir string, exe string, args string) {
	fmt.Println("development mode")
	switch runtime.GOOS {
	case "darwin":
		return "./app/node_modules/electron/dist/", "Electron.app", "app"
	}

	return "./app/node_modules/electron/dist/", "electron", "app"
}
