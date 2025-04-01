package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {

	// define routes
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	mux.Use(app.enableCORS)

	mux.Get("/", app.Home)
	mux.Get("/api/movies", app.AllMovies)
	mux.Post("/api/authenticate", app.Authenticate)
	mux.Post("/api/signup", app.Signup)

	return mux
}
