package handler

import (
	"cloud-martini-backend/db"
	"cloud-martini-backend/dto"
	"cloud-martini-backend/queries"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUsers(ctx *gin.Context, getUserFunc func(collection *mongo.Collection) ([]map[string]interface{}, error)) {

	collection, err := db.GetCollection("users")
	if err != nil {
		panic(err)
	}

	users, err := queries.GetUsers(collection)
	if err != nil {
		log.Printf("Error fetching all users: %v", err)
	}

	jsonContent, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Error converting data ")
		panic(err)
	}

	ctx.JSON(200, gin.H{
		"data": json.RawMessage(jsonContent),
	})
}

func AddUsers(ctx *gin.Context, insertUserFunc func(collection *mongo.Collection, users dto.Users) (*mongo.InsertOneResult, error)) {
	collection, err := db.GetCollection("users")
	if err != nil {
		panic(err)
	}

	var users dto.Users
	jsonError := ctx.ShouldBindBodyWithJSON(&users)
	if jsonError != nil {
		panic(jsonError)
	}
	_, err2 := queries.InsertUser(collection, users)
	if err2 != nil {
		ctx.JSON(404, gin.H{
			"status":  false,
			"message": "Error Adding Data",
		})
	}
	ctx.JSON(200, gin.H{
		"status":  true,
		"message": "Data Added Successfully",
	})
}

func DeleteUsers(ctx *gin.Context, deleteUserFunc func(collection *mongo.Collection, objectID primitive.ObjectID) (*mongo.DeleteResult, error)) {
	collection, err := db.GetCollection("users")
	if err != nil {
		panic(err)
	}

	userId := ctx.Param("id")
	objectID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Invalid ObjectID"})
		return
	}
	fmt.Println(objectID)
	deleteResult, err := queries.DeleteUser(collection, objectID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Error deleting user"})
		return
	}

	if deleteResult.DeletedCount == 0 {
		ctx.JSON(404, gin.H{"message": "No user found with the given ID"})
	} else {
		ctx.JSON(200, gin.H{"message": "User deleted successfully"})
	}
}

func UpdateUsers(ctx *gin.Context, updateUserFunc func(collection *mongo.Collection, objectID primitive.ObjectID, updateUsers dto.Users) (*mongo.UpdateResult, error)) {
	collection, err := db.GetCollection("users")
	if err != nil {
		panic(err)
	}

	userId := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Invalid ObjectID"})
		return
	}
	fmt.Println(objectId)
	var updateUsers dto.Users
	userUpdateErr := ctx.ShouldBindJSON(&updateUsers)
	if userUpdateErr != nil {
		ctx.JSON(200, gin.H{"error": "Invalid user data"})
		return
	}

	updateResult, err := queries.UpdateUsers(collection, objectId, updateUsers)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	if updateResult.MatchedCount == 0 {
		ctx.JSON(404, gin.H{"message": "No user found with the given ID"})
		return
	}
	ctx.JSON(200, gin.H{"message": "User updated successfully"})
}
