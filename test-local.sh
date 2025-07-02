#!/bin/bash

echo "Testing local Docker Compose deployment..."

# Start the services
echo "Starting services with Docker Compose..."
docker-compose up -d

# Wait for services to be ready
echo "Waiting for services to be ready..."
sleep 10

# Test health endpoint
echo "Testing health endpoint..."
curl -f http://localhost:8080/healthz && echo " - Health check OK" || echo " - Health check FAILED"

# Test getting todos (should be empty initially)
echo "Testing GET /todo (should be empty)..."
curl -s http://localhost:8080/todo | jq .

# Test creating a todo
echo "Testing POST /todo..."
RESPONSE=$(curl -s -X POST http://localhost:8080/todo -d "task=Test local deployment")
echo $RESPONSE | jq .
TODO_ID=$(echo $RESPONSE | jq -r .id)

# Test getting todos (should have one item)
echo "Testing GET /todo (should have one item)..."
curl -s http://localhost:8080/todo | jq .

# Test deleting the todo
echo "Testing DELETE /todo/$TODO_ID..."
curl -s -X DELETE http://localhost:8080/todo/$TODO_ID

# Test getting todos (should be empty again)
echo "Testing GET /todo (should be empty again)..."
curl -s http://localhost:8080/todo | jq .

echo "Local testing complete!"
echo "To stop the services, run: docker-compose down"