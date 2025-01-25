package handler

import (
	"cloud-martini-backend/queries"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUsers(ctx *gin.Context) {
	MONGO_URI := os.Getenv("MONGO_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		fmt.Printf("Error connecting to MongoDB: %v", err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			fmt.Printf("Error disconnecting MongoDB: %v", err)
		}
	}()

	fmt.Println("Connected to MongoDB!")

	// Get the collection
	collection := client.Database("cloud-martini").Collection("users")

	// Fetch all users
	users, err := queries.GetUsers(collection)
	if err != nil {
		log.Printf("Error fetching all users: %v", err)
	} else {
		fmt.Println("Users found:")
		for _, user := range users {
			fmt.Printf("%v\n", user)
		}
	}
}
