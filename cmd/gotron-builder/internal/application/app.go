package application

import (
	"fmt"
)

type App struct {
}

func (app *App) Run() (err error) {
	fmt.Print("All your bases are belong to us!\n")
	return nil
}

func New() *App {
	return &App{}
}
