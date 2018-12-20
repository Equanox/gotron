package main

import (
	"log"

	"github.com/Equanox/gotron"
)

func main() {

	window, err := gotron.New()
	if err != nil {
		log.Println(err)
		return
	}

	window.WindowOptions.Width = 1200
	window.WindowOptions.Height = 600

	done, err := window.Start()
	if err != nil {
		log.Println(err)
		return
	}

	//window.OpenDevTools()

	<-done
}
