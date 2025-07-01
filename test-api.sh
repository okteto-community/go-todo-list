#!/bin/bash

# Test script for the Todo API
BASE_URL="http://localhost:8080"

echo "Testing Todo API..."
echo "==================="

# Test health endpoint
echo "1. Testing health endpoint..."
curl -s "$BASE_URL/healthz" && echo " ✓ Health check passed" || echo " ✗ Health check failed"

# Test getting todos (should be empty initially)
echo "2. Getting all todos..."
curl -s "$BASE_URL/todo" | jq '.' || echo "No jq installed, raw response:"

# Test creating a todo
echo "3. Creating a new todo..."
RESPONSE=$(curl -s -X POST -d "task=Test todo item" "$BASE_URL/todo")
echo "Response: $RESPONSE"
TODO_ID=$(echo "$RESPONSE" | jq -r '.id' 2>/dev/null)

# Test getting todos again
echo "4. Getting all todos after creation..."
curl -s "$BASE_URL/todo" | jq '.' || curl -s "$BASE_URL/todo"

# Test deleting the todo (if we got an ID)
if [ "$TODO_ID" != "null" ] && [ -n "$TODO_ID" ]; then
    echo "5. Deleting todo with ID: $TODO_ID"
    curl -s -X DELETE "$BASE_URL/todo/$TODO_ID" && echo " ✓ Delete successful" || echo " ✗ Delete failed"
    
    echo "6. Getting all todos after deletion..."
    curl -s "$BASE_URL/todo" | jq '.' || curl -s "$BASE_URL/todo"
else
    echo "5. Skipping delete test (no valid ID received)"
fi

echo "==================="
echo "API testing complete!"