package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func (app *application) routes() http.Handler {
	// Load environment variable
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}

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

	return mux
}
