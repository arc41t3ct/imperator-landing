package main

import (
	"net/http"

	chi "github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {
	// middleware must come before any routes using aliases
	// a.use(a.Middlware.Remember)
	// routes go here using the aloases
	a.get("/", a.Handlers.Home)
	// static routes
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
