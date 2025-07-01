# Todo List Application with PostgreSQL Database

This application has been updated to use PostgreSQL as the database backend instead of in-memory storage.

## Changes Made

### 1. Application Updates
- **Database Integration**: Replaced in-memory map storage with PostgreSQL database operations
- **Environment Variables**: Added support for database configuration via environment variables
- **Database Schema**: Automatically creates a `todos` table with `id` and `task` columns
- **Error Handling**: Added proper error handling for database operations

### 2. New Files Created
- `docker-compose.yml`: Orchestrates both the application and PostgreSQL database
- `okteto.yml`: Okteto manifest for cloud-native deployment
- `.env`: Environment variables for local development

### 3. Dependencies Added
- `github.com/lib/pq`: PostgreSQL driver for Go

## Environment Variables

The application uses the following environment variables for database configuration:

- `DB_HOST`: Database host (default: localhost)
- `DB_PORT`: Database port (default: 5432)
- `DB_USER`: Database username (default: postgres)
- `DB_PASSWORD`: Database password (default: password)
- `DB_NAME`: Database name (default: todoapp)

## Deployment Options

### Option 1: Docker Compose (Local Development)

1. **Prerequisites**: Docker and Docker Compose installed
2. **Deploy**: Run the following command in the project directory:
   ```bash
   docker-compose up --build
   ```
3. **Access**: The application will be available at http://localhost:8080

### Option 2: Okteto Cloud Deployment

1. **Prerequisites**: Okteto CLI installed and authenticated
2. **Deploy**: Run the following command in the project directory:
   ```bash
   okteto deploy --wait
   ```
3. **Access**: Use `okteto endpoints` to get the public URL

### Option 3: Manual Deployment

1. **Start PostgreSQL**: 
   ```bash
   docker run -d \
     --name postgres \
     -e POSTGRES_DB=todoapp \
     -e POSTGRES_USER=postgres \
     -e POSTGRES_PASSWORD=password \
     -p 5432:5432 \
     postgres:15-alpine
   ```

2. **Build and Run Application**:
   ```bash
   # Set environment variables
   export DB_HOST=localhost
   export DB_PORT=5432
   export DB_USER=postgres
   export DB_PASSWORD=password
   export DB_NAME=todoapp
   
   # Build and run (requires Go installed)
   go build -o app
   ./app
   ```

## Database Schema

The application automatically creates the following table:

```sql
CREATE TABLE IF NOT EXISTS todos (
    id VARCHAR(255) PRIMARY KEY,
    task TEXT NOT NULL
);
```

## API Endpoints

- `GET /healthz`: Health check endpoint
- `GET /todo`: Get all todo items
- `POST /todo`: Create a new todo item (form data: `task`)
- `DELETE /todo/{id}`: Delete a todo item by ID
- `GET /`: Serve static files (web interface)

## Development

For development with Okteto:

```bash
okteto up
```

This will start a development environment with:
- Live code synchronization
- Port forwarding (8080:8080)
- Database connection configured

## Troubleshooting

### Database Connection Issues
- Ensure PostgreSQL is running and accessible
- Check environment variables are set correctly
- Verify network connectivity between application and database

### Application Logs
- Docker Compose: `docker-compose logs app`
- Okteto: `kubectl logs -n ${OKTETO_NAMESPACE} <pod-name>`

### Database Logs
- Docker Compose: `docker-compose logs postgres`
- Okteto: `kubectl logs -n ${OKTETO_NAMESPACE} <postgres-pod-name>`