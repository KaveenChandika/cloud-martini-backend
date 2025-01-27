package db_test

import (
	"cloud-martini-backend/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectMongo(t *testing.T) {
	// Test cases
	tests := []struct {
		name        string
		uri         string
		expectError bool
	}{
		{
			name:        "Valid MongoDB URI",
			uri:         "mongodb+srv://Kaveen:qX10lodLpHHEDFLg@cluster1.i6vai.mongodb.net/cloud-martini", // Use your valid local MongoDB URI
			expectError: false,
		},
		{
			name:        "Invalid MongoDB URI",
			uri:         "mongodb://invalid:27017",
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := db.ConnectMongo(tc.uri)

			if tc.expectError {
				assert.Error(t, err, "Expected an error but got none")
			} else {
				assert.NoError(t, err, "Expected no error but got one")
				assert.NotNil(t, db.MongoClient, "MongoClient should not be nil")
			}

			// Disconnect after successful connection
			if db.MongoClient != nil {
				db.DisconnectMongo()
				assert.Nil(t, db.MongoClient, "MongoClient should be nil after disconnect")
			}
		})
	}
}
