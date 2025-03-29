package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func openConnection(dsn string) (*mongo.Client, error) {
	// Open a connection to the database
	clientOptions := options.Client().ApplyURI(dsn)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (app *application) connectToDB() (*mongo.Client, error) {
	// Open a connection to the database
	client, err := openConnection(app.DSN)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB!")

	return client, nil

}
