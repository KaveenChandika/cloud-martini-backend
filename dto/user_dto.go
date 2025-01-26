package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string
	Designation string
	Email       string
	Projects    []string
}
