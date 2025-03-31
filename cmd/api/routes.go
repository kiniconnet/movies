package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	// define routes
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	if os.Getenv("ENV") != "production" {
		mux.Use(app.enableCORS)
	}

	mux.Get("/", app.Home)
	mux.Get("/api/movies", app.AllMovies)
	mux.Post("/api/authenticate", app.Authenticate)
	mux.Post("/api/signup", app.Signup)

	dir, err := filepath.Abs("./client/dist")
	if err != nil {
		log.Fatalf("Failed to resolve static files directory: %v", err)
	}

	fileServer := http.FileServer(http.Dir(dir))
	http.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
