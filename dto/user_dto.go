package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name        string             `json:"name"`
	Designation string             `json:"Designation"`
	Email       string             `json:"email"`
	Projects    []string           `json:"projects"`
}
