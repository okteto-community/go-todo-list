# Todo List Application

This is a simple todo list application written in Go that uses PostgreSQL as the database backend. The application can be deployed using Docker Compose or on Okteto.

## Features

- Create, read, and delete todo items
- PostgreSQL database for persistent storage
- RESTful API endpoints
- Web interface for managing todos

## Architecture

- **Backend**: Go application with REST API
- **Database**: PostgreSQL 15
- **Frontend**: Static HTML/CSS/JS served by the Go application

## API Endpoints

- `GET /todo` - Get all todo items
- `POST /todo` - Create a new todo item (form data: `task`)
- `DELETE /todo/{id}` - Delete a todo item by ID
- `GET /healthz` - Health check endpoint

## Environment Variables

The application supports the following environment variables for database configuration:

- `DB_HOST` - Database host (default: localhost)
- `DB_PORT` - Database port (default: 5432)
- `DB_USER` - Database user (default: postgres)
- `DB_PASSWORD` - Database password (default: postgres)
- `DB_NAME` - Database name (default: todoapp)

## Local Development with Docker Compose

To run the application locally using Docker Compose:

```bash
docker-compose up --build
```

This will start both the PostgreSQL database and the todo application. The application will be available at http://localhost:8080.

## Deployment on Okteto

To deploy on Okteto:

```bash
okteto deploy --wait
```

This will build the application image and deploy both the database and application using the Docker Compose configuration.

## Database Schema

The application creates a `todos` table with the following structure:

```sql
CREATE TABLE todos (
    id VARCHAR(36) PRIMARY KEY,
    task TEXT NOT NULL
);
```
