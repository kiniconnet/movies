package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title        string             `json:"title" bson:"title"`
	ReleasedDate time.Time          `json:"released_date" bson:"released_date"`
	RunTime      int                `json:"run_time" bson:"run_time"`
	MPPARating   string             `json:"mppa_rating" bson:"mppa_rating"`
	Description  string             `json:"description" bson:"description"`
	Image        string             `json:"image,omitempty" bson:"image,omitempty"`
	CreatedAt    time.Time          `json:"-" bson:"-"`
	UpdatedAt    time.Time          `json:"-" bson:"-"`
}








