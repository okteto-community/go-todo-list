# Go Todo App with PostgreSQL

A simple todo application written in Go that uses PostgreSQL as the database backend.

## Features

- Create, read, and delete todo items
- PostgreSQL database for persistent storage
- Docker Compose setup for easy deployment
- Okteto support for cloud-native development

## Local Development with Docker Compose

### Prerequisites

- Docker and Docker Compose installed

### Running the Application

1. Clone the repository and navigate to the project directory
2. Start the application with Docker Compose:

```bash
docker-compose up --build
```

This will:
- Start a PostgreSQL database container
- Build and start the todo application container
- The application will be available at http://localhost:8080

### Environment Variables

The application uses the following environment variables for database configuration:

- `DB_HOST`: Database host (default: localhost)
- `DB_PORT`: Database port (default: 5432)
- `DB_USER`: Database user (default: postgres)
- `DB_PASSWORD`: Database password (default: password)
- `DB_NAME`: Database name (default: todoapp)

## Okteto Deployment

### Prerequisites

- Okteto CLI installed and configured

### Deploying to Okteto

1. Deploy the application:

```bash
okteto deploy --wait
```

2. Access your application using the public endpoint provided by Okteto

### Development Mode

To start development mode in Okteto:

```bash
okteto up
```

This will sync your local code changes to the remote development environment.

## API Endpoints

- `GET /healthz` - Health check endpoint
- `GET /todo` - Get all todo items
- `POST /todo` - Create a new todo item (form data: `task`)
- `DELETE /todo/{id}` - Delete a todo item by ID

## Testing

### Local Testing

A test script is provided to verify the application functionality:

```bash
./test-local.sh
```

This script will:
1. Start the services with Docker Compose
2. Test all API endpoints
3. Verify CRUD operations work correctly

### Manual Testing

You can also test the API manually:

```bash
# Health check
curl http://localhost:8080/healthz

# Get all todos
curl http://localhost:8080/todo

# Create a todo
curl -X POST http://localhost:8080/todo -d "task=My new task"

# Delete a todo (replace {id} with actual ID)
curl -X DELETE http://localhost:8080/todo/{id}
```

## Database Schema

The application creates a `todos` table with the following structure:

```sql
CREATE TABLE todos (
    id VARCHAR(36) PRIMARY KEY,
    task TEXT NOT NULL
);
```
