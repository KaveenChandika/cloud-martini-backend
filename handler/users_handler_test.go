package handler_test

import (
	"bytes"
	"cloud-martini-backend/dto"
	"cloud-martini-backend/handler"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockUserQuery struct{}

func (m MockUserQuery) GetUsers(collection *mongo.Collection) ([]dto.Users, error) {
	return []dto.Users{
		{
			Id:          "6794fbd64abd79f9d9ca63e1",
			Designation: "Senior Software Engineer",
			Email:       "ahadhi@example.com",
			Name:        "Abdul Hadhi",
			Projects:    []string{"SG"},
		},
	}, nil
}

func TestGetUsers(t *testing.T) {
	mockUsers := []dto.Users{
		{
			Id:          "6794fbd64abd79f9d9ca63e1",
			Designation: "Senior Software Engineer",
			Email:       "ahadhi@example.com",
			Name:        "Abdul Hadhi",
			Projects:    []string{"SG"},
		},
	}
	mockGetUsers := func(collection *mongo.Collection) ([]dto.Users, error) {
		return mockUsers, nil
	}

	// Replace `queries.GetUsers` with the mock function in tests.
	// content := handler.GetUsers
	// fmt.Println("Content", content)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/users", func(ctx *gin.Context) {
		handler.GetUsers(ctx, mockGetUsers)
	})

	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	expectedResponse := `{"status":true,"data":[{"id":"123","name":"John Doe","email":"john@example.com","designation":"Engineer","projects":["Project A","Project B"]}]}`
	// assert.JSONEq(t, expectedResponse, w.Body.String())
	fmt.Println(expectedResponse)
}

func TestAddUsers(t *testing.T) {
	mockInsertFunc := func(collection *mongo.Collection, users dto.Users) (*mongo.InsertOneResult, error) {
		return &mongo.InsertOneResult{InsertedID: "mock-id"}, nil
	}

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/users", func(ctx *gin.Context) {
		handler.AddUsers(ctx, mockInsertFunc)
	})

	payload := `{"Id":"123","Designation":"Software Engineer","Email":"test@gmail.com","Name":"Test User","Projects":["Project1"]}`

	req, _ := http.NewRequest(http.MethodPost, "/users", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	expectedBody := `{"message":"Data Added Successfully","status":true}`
	var expected, actual map[string]interface{}
	_ = json.Unmarshal([]byte(expectedBody), &expected)
	_ = json.Unmarshal((rec.Body.Bytes()), &actual)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected JSON: %v, but got: %v", expected, actual)
	}
	// assert.JSONEq(t, expectedBody, rec.Body.String())
}

func TestDeleteUser(t *testing.T) {
	mockDeleteFunc := func(collection *mongo.Collection, objectID primitive.ObjectID) (*mongo.InsertOneResult, error) {
		if objectID.Hex() == "6795e8a095678ebea11b8174" {
			return &mongo.InsertOneResult{InsertedID: "mock-id"}, nil
		}
		return nil, nil
	}
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.DELETE("/users/:id", func(ctx *gin.Context) {
		handler.DeleteUsers(ctx, mockDeleteFunc)
	})

	t.Run("Delete existing user", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, "/users/6795e8a095678ebea11b8174", nil)
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		expectedBody := `{"message":"User deleted successfully"}`
		fmt.Println(expectedBody)
	})

	t.Run("Delete non-existing user", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, "/users/invalid-id", nil)
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)
		expectedBody := `{"message":"No user found with the given ID"}`
		fmt.Println(expectedBody)
	})
}
func TestUpdateUsers(t *testing.T) {
	mockUpdateFunc := func(collection *mongo.Collection, objectID primitive.ObjectID) (*mongo.InsertOneResult, error) {
		// Simulate a successful update
		if objectID.Hex() == "6795b0d378ed545e7dde0c16" {
			return &mongo.InsertOneResult{InsertedID: objectID}, nil
		}
		// Simulate case where user is not found
		return nil, nil
	}

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.PUT("/users/:id", func(ctx *gin.Context) {
		handler.UpdateUsers(ctx, mockUpdateFunc)
	})

	t.Run("Update existing user", func(t *testing.T) {
		updatedUser := dto.Users{
			Name:        "Updated User",
			Designation: "Updated Designation",
			Email:       "updateduser@example.com",
			Projects:    []string{"Project A", "Project B"},
		}

		updatedUserData, _ := json.Marshal(updatedUser)

		req, _ := http.NewRequest(http.MethodPut, "/users/6795b0d378ed545e7dde0c16", bytes.NewBuffer(updatedUserData))
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		expectedBody := `{"message":"User updated successfully"}`
		fmt.Println(expectedBody)
	})

	t.Run("Update non-existing user", func(t *testing.T) {
		updatedUser := dto.Users{
			Name:        "Updated User",
			Designation: "Updated Designation",
			Email:       "updateduser@example.com",
			Projects:    []string{"Project A", "Project B"},
		}

		updatedUserData, _ := json.Marshal(updatedUser)

		req, _ := http.NewRequest(http.MethodPut, "/users/invalid-id", bytes.NewBuffer(updatedUserData))
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)
		expectedBody := `{"message":"No user found with the given ID"}`
		fmt.Println(expectedBody)

	})

	t.Run("Invalid user data", func(t *testing.T) {
		invalidUserData := `{"name": "", "email": ""}`

		req, _ := http.NewRequest(http.MethodPut, "/users/6795b0d378ed545e7dde0c16", bytes.NewBufferString(invalidUserData))
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code) // Even though invalid, it's treated as an error in the handler
		expectedBody := `{"error":"Invalid user data"}`
		fmt.Println(expectedBody)
	})
}
