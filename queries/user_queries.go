package queries

import (
	"cloud-martini-backend/dto"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUsers(collection *mongo.Collection) ([]dto.Users, error) {
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var users []dto.Users

	for cursor.Next(context.TODO()) {
		var user dto.Users
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func InsertUser(collection *mongo.Collection, users dto.Users) (*mongo.InsertOneResult, error) {
	insertResult, err := collection.InsertOne(context.TODO(), users)
	if err != nil {
		fmt.Println("Error Inserting Values")
		return nil, err
	}

	return insertResult, nil

}

func DeleteUser(collection *mongo.Collection, objectId primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": objectId}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter, options.Delete())
	if err != nil {
		fmt.Printf("Error deleting user: %v\n", err)
		return nil, err
	}
	return deleteResult, nil
}

func UpdateUsers(collection *mongo.Collection, objectId primitive.ObjectID, updateUsers dto.Users) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": objectId}
	// Define the update operation
	update := bson.M{
		"$set": bson.M{
			"designation": updateUsers.Designation,
			"email":       updateUsers.Email,
			"name":        updateUsers.Name,
			"projects":    updateUsers.Projects,
		},
	}

	// Perform the update operation
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Printf("Error updating user: %v\n", err)
		return nil, err
	}

	return updateResult, nil
}
