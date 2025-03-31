package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
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
	// Load environment variable
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}

	// set application configuration
	var app application


	// read from the command line
	flag.StringVar(&app.DSN, "dsn", os.Getenv("MONGODB_URI"), "MongoDB connection string")
	flag.StringVar(&app.JWTScret, "jwt-secret", "verysecret", "signing secret for JWT")
	flag.StringVar(&app.JWTIssuer, "jwt-issuer", "example.com", "issuer for JWT")
	flag.StringVar(&app.JWTAudience, "jwt-audience", "example.com", "audience for JWT")
	flag.StringVar(&app.CookieDomain, "cookie-domain", "localhost", "domain for cookie")
	flag.StringVar(&app.Domain, "domain", "example.com", "domain for the application")
	flag.BoolVar(&app.Config.LoadStatic, "loadStatic", true, "This is use to load Static file to the server")
	flag.BoolVar(&app.Config.InProduction, "production", true, "This is to be in production")
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
	if port == "" {
		port = "9000"
	}

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
