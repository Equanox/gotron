// +build !prod

package main

import (
	"fmt"
	"runtime"
)

// FrontendPath Sets path to call electron in development environment
func FrontendPath() (string, string, string) {
	fmt.Println("!prod")
	if runtime.GOOS == "darwin" {
		return "./app/node_modules/electron/dist/", "Electron.app", "app"
	}
	return "./app/node_modules/electron/dist/", "electron", "app"
}
