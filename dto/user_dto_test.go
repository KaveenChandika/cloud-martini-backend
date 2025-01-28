package dto_test

import (
	"cloud-martini-backend/dto"
	"reflect"
	"testing"
)

func TestUsersJSONTags(t *testing.T) {
	// Define the expected JSON tags for each field
	expectedTags := map[string]string{
		"ID":          "_id",
		"Name":        "name",
		"Designation": "designation",
		"Email":       "email",
		"Projects":    "projects",
	}

	// Use reflection to iterate through the fields of the Users struct
	userType := reflect.TypeOf(dto.Users{})
	for fieldName, expectedTag := range expectedTags {
		field, found := userType.FieldByName(fieldName)
		if !found {
			t.Errorf("Field %s is missing in the Users struct", fieldName)
			continue
		}

		actualTag := field.Tag.Get("json")
		if actualTag != expectedTag {
			t.Errorf(
				"Field %s has incorrect json tag. Expected '%s', got '%s'",
				fieldName, expectedTag, actualTag,
			)
		}
	}
}
