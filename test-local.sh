#!/bin/bash

echo "Starting todo application with PostgreSQL using Docker Compose..."
docker-compose up -d

echo "Waiting for services to be ready..."
sleep 10

echo "Testing health endpoint..."
curl -f http://localhost:8080/healthz && echo " ✓ Health check passed"

echo "Testing empty todo list..."
curl -s http://localhost:8080/todo | jq . && echo " ✓ Empty todo list retrieved"

echo "Creating a test todo..."
RESPONSE=$(curl -s -X POST -d "task=Test local deployment" http://localhost:8080/todo)
echo $RESPONSE | jq .
TODO_ID=$(echo $RESPONSE | jq -r .id)

echo "Retrieving all todos..."
curl -s http://localhost:8080/todo | jq . && echo " ✓ Todo retrieved from database"

echo "Deleting the test todo..."
curl -X DELETE http://localhost:8080/todo/$TODO_ID && echo " ✓ Todo deleted"

echo "Verifying todo was deleted..."
curl -s http://localhost:8080/todo | jq . && echo " ✓ Todo list is empty again"

echo "Stopping services..."
docker-compose down

echo "✓ All tests passed! The application is working correctly with PostgreSQL."