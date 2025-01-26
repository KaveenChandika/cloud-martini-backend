package handler_test

import (
	"cloud-martini-backend/dto"
	"cloud-martini-backend/handler"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
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
	// mockUsers := []dto.Users{
	// 	{
	// 		Id:          "6794fbd64abd79f9d9ca63e1",
	// 		Designation: "Senior Software Engineer",
	// 		Email:       "ahadhi@example.com",
	// 		Name:        "Abdul Hadhi",
	// 		Projects:    []string{"SG"},
	// 	},
	// }

	// mockGetUsers := func(collection string) ([]dto.Users, error) {
	// 	return mockUsers, nil
	// }

	// // Replace `queries.GetUsers` with the mock function in tests.
	// content := handler.GetUsers
	// fmt.Println("Content", content)

	// gin.SetMode(gin.TestMode)
	// router := gin.Default()
	// router.GET("/users", handler.GetUsers)

	// req, _ := http.NewRequest("GET", "/users", nil)
	// w := httptest.NewRecorder()

	// router.ServeHTTP(w, req)

	// assert.Equal(t, http.StatusOK, w.Code)

	// // Verify the response
	// expectedResponse := `{"status":true,"data":[{"id":"123","name":"John Doe","email":"john@example.com","designation":"Engineer","projects":["Project A","Project B"]}]}`
	// assert.JSONEq(t, expectedResponse, w.Body.String())
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
