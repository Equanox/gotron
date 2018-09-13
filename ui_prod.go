// +build prod

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func FrontendPath() (string, string, string) {
	fmt.Println("prod")

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	return exPath + "/", "electron-rampup", ""
}
