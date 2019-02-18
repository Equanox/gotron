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

	window.On("hello", func(i interface{}) {
		log.Println("received hello")
		log.Println(i)
		window.Send("hello From backend", i)
	})

	//window.OpenDevTools()

	<-done
}
