// +build !prod

package main

import "fmt"

// FrontendPath Sets path to call electron in development environment
func FrontendPath() (string, string, string) {
	fmt.Println("!prod")
	return "./app/node_modules/electron/dist/", "electron", "app"
}
