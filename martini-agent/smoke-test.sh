#!/bin/bash

# Set API base URL
BASE_URL="http://cloud-martini-backend-api-preview.default.svc.cluster.local"

# Test GET /healthCheck
echo "Testing GET /health"
curl -X GET "$BASE_URL/health"
echo -e "\n"

# Test GET /users
echo "Testing GET /users"
curl -X GET "$BASE_URL/users"
echo -e "\n"

# Test POST /users
echo "Testing POST /users"
curl -X POST "$BASE_URL/users" \
  -H "Content-Type: application/json" \
  -d '{
        "Id": "123",
        "Designation": "Software Engineer",
        "Email": "test@gmail.com",
        "Name": "Test User",
        "Projects": ["Project1"]
      }'
echo -e "\n"

# Test DELETE /user/:id
echo "Testing DELETE /user/:id"
USER_ID="6795e8a095678ebea11b8174" # Replace with a valid ID if available
curl -X DELETE "$BASE_URL/user/$USER_ID"
echo -e "\n"

# Test PUT /user/:id
echo "Testing PUT /user/:id"
USER_ID="6795b0d378ed545e7dde0c16" # Replace with a valid ID if available
curl -X PUT "$BASE_URL/user/$USER_ID" \
  -H "Content-Type: application/json" \
  -d '{
        "Name": "Updated User",
        "Designation": "Updated Designation",
        "Email": "updateduser@example.com",
        "Projects": ["Project A", "Project B"]
      }'
echo -e "\n"

# Test POST /order
echo "Testing POST /order"
curl -X POST "$BASE_URL/order" \
  -H "Content-Type: application/json" \
  -d '{
        "orderId": "order123",
        "productName": "Product 1",
        "quantity": 2
      }'
echo -e "\n"

# Test GET /orders
echo "Testing GET /orders"
curl -X GET "$BASE_URL/orders"
echo -e "\n"