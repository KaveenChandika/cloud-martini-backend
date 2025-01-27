package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongo(uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("error connecting to MongoDB: %w", err)
	}

	// Test the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("MongoDB ping failed: %w", err)
	}

	MongoClient = client
	fmt.Println("Connected to MongoDB!")
	return nil
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

func GetCollection(collectionName string) *mongo.Collection {
	MONGO_URI := "mongodb+srv://Kaveen:qX10lodLpHHEDFLg@cluster1.i6vai.mongodb.net/cloud-martini"
	ConnectMongo(MONGO_URI)

	var databaseName string = "cloud-martini"
	if MongoClient == nil {
		log.Fatalf("MongoClient is not initialized. Call ConnectMongo first.")
	}
	return MongoClient.Database(databaseName).Collection(collectionName)
}
