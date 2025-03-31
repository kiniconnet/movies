package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kiniconnet/react-go-tutorial/internal/config"
	"github.com/kiniconnet/react-go-tutorial/internal/repository"
	dbrepo "github.com/kiniconnet/react-go-tutorial/internal/repository/db_repo"
)

type application struct {
	Domain       string
	DSN          string
	DB           repository.DatabaseRepository
	auth         Auth
	JWTScret     string
	JWTIssuer    string
	JWTAudience  string
	CookieDomain string
	Config       config.Config
}

func main() {
	// set application configuration
	var app application

	// read from the command line
	flag.StringVar(&app.DSN, "dsn", "mongodb+srv://kiniconnet:kiniconnet2025@cluster0.at1fb.mongodb.net/golang_db?retryWrites=true&w=majority&appName=Cluster0", "MongoDB connection string")
	flag.StringVar(&app.JWTScret, "jwt-secret", "verysecret", "signing secret for JWT")
	flag.StringVar(&app.JWTIssuer, "jwt-issuer", "example.com", "issuer for JWT")
	flag.StringVar(&app.JWTAudience, "jwt-audience", "example.com", "audience for JWT")
	flag.StringVar(&app.CookieDomain, "cookie-domain", "localhost", "domain for cookie")
	flag.StringVar(&app.Domain, "domain", "example.com", "domain for the application")
	flag.Parse()

	// Connect to the database
	client, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the repository
	repo := dbrepo.NewMongoDBRepo(client)
	app.DB = repo

	// close the database connection
	defer app.DB.Connection().Disconnect(context.Background())

	port := os.Getenv("PORT")

	app.auth = Auth{
		Issuer:        app.JWTIssuer,
		Audience:      app.JWTAudience,
		Secret:        app.JWTScret,
		TokenExpiry:   15 * time.Minute,
		RefreshExpiry: 24 * time.Hour,
		CookiePath:    "/",
		CookieDomain:  app.CookieDomain,
		CookieName:    "__HOST-refresh_token",
	}

	log.Println("Starting application on port", port)

	// start the server
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), app.routes())
	if err != nil {
		log.Println("Error starting server:", err)
	}
}
