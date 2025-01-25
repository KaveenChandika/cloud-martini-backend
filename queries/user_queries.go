package queries

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUsers(collection *mongo.Collection) ([]map[string]interface{}, error) {
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var users []map[string]interface{}
	// Iterate through the cursor and decode each document
	for cursor.Next(context.TODO()) {
		var user map[string]interface{}
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// Check if the cursor encountered any errors
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
