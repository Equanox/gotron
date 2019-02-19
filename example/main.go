package main

import (
	"log"

	"github.com/Equanox/gotron"
)

func main() {

	window, err := gotron.New("ui/build")
	if err != nil {
		log.Println(err)
		return
	}

	window.WindowOptions.Width = 1200
	window.WindowOptions.Height = 600
	window.WindowOptions.Title = "Gotron"

	done, err := window.Start()
	if err != nil {
		log.Println(err)
		return
	}

	onEvent := gotron.Event{Event: "hello"}

	window.On(&onEvent, func(bin []byte) {
		log.Println("received hello")
		log.Println(bin)
		window.Send(&gotron.Event{Event: "hello From backend"})
	})

	//window.OpenDevTools()

	<-done
}
