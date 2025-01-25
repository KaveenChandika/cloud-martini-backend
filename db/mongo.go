package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongo(uri string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("MongoDB ping failed: %v", err)
	}

	MongoClient = client
	fmt.Println("Connected to MongoDB!")
}

func DisconnectMongo() {
	if MongoClient != nil {
		err := MongoClient.Disconnect(context.TODO())
		if err != nil {
			log.Fatalf("Error disconnecting from MongoDB: %v", err)
		}
		fmt.Println("Disconnected from MongoDB!")
	}
}

func GetCollection(collectionName string) (*mongo.Collection, error) {
	var MONGO_URI string = os.Getenv("MONGO_URI")
	var MONGO_DB string = os.Getenv("MONGO_DB")

	fmt.Println(MONGO_DB, collectionName)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		fmt.Printf("Error connecting to MongoDB: %v", err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			fmt.Printf("Error disconnecting MongoDB: %v", err)
		}
	}()

	return client.Database(MONGO_DB).Collection(collectionName), err

}
