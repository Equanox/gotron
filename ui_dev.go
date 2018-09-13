// +build !prod

package main

import "fmt"

func FrontendPath() (string, string, string) {
	fmt.Println("!prod")
	return "./app/node_modules/electron/dist/", "electron", "app"
}
