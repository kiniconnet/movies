package dbrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/kiniconnet/react-go-tutorial/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBRepo struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewMongoDBRepo(client *mongo.Client) *MongoDBRepo {
	return &MongoDBRepo{
		Client:     client,
		Collection: client.Database("golang_db").Collection("users"),
	}
}

const dbTimeout = 3 * time.Second

func (r *MongoDBRepo) Connection() *mongo.Client {
	return r.Client
}

func (r *MongoDBRepo) AllMovies() ([]*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	cursor, err := r.Client.Database("golang_db").Collection("movies").Find(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var movies []*models.Movie

	for cursor.Next(ctx) {
		var movie models.Movie

		err := cursor.Decode(&movie)
		if err != nil {
			return nil, err
		}

		movies = append(movies, &movie)
	}

	return movies, nil
}

func (r *MongoDBRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user models.User
	err := r.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil // No user found
	} else if err != nil {
		return nil, fmt.Errorf("failed to fetch user: %v", err)
	}

	return &user, nil
}

func (r *MongoDBRepo) InsertUser(user models.User) error {
	_, err := r.Collection.InsertOne(context.Background(), user)
	if err != nil {
		fmt.Printf("unable to insert user: %S", err)
	}
	return err
}
