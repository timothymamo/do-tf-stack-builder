package main

import (
	"github.com/go-chi/chi/v5"
)

const (
	version = "v0.2.0"
)

func main() {

	app := App{
		Router: chi.NewRouter(),
	}

	app.initializeRoutes()

	app.Run(":8010")
}
