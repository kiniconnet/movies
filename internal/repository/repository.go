package repository

import (
	"github.com/kiniconnet/react-go-tutorial/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type DatabaseRepository interface {
	Connection() *mongo.Client
	GetUserByEmail(email string) (*models.User, error)
	AllMovies() ([]*models.Movie, error)
	InsertUser(models.User) error
}
