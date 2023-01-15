package main

import "github.com/go-chi/chi/v5"

type App struct {
	Router chi.Router
}

func (app *App) initializeRoutes() {

	app.Router.Get("/", app.Index)
	// Name("Index").

	app.Router.Get("/health", Health)
	// Name("Health").

	app.Router.Get("/version", Version)
	// Name("Version").

	app.Router.Post("/tffiles", app.CreateTFFiles)
	// Name("TF Files").
}
