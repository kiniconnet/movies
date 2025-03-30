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

	if app.Config.InProduction { 
		mux.Use(app.enableCORS)
		 fs := http.FileServer(http.Dir("./client/dist"))
		 http.Handle("/client/dist", http.StripPrefix("client/dist", fs))
	}



	mux.Get("/", app.Home)
	mux.Get("/api/movies", app.AllMovies)
	mux.Post("/api/authenticate", app.Authenticate)
	mux.Post("/api/signup", app.Signup)

	return mux
}
