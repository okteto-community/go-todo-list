# Go Todo List Application - Okteto AI Agent Fleet Demo

🚀 **Live Demo**: https://todo-app-agent-quh1cdulyz9i.rberrelleza.fleets.okteto.ai

A demonstration of how the **Okteto AI Agent Fleet** transforms cloud-native development by automatically migrating a simple in-memory todo application to use PostgreSQL with full Docker Compose and Kubernetes deployment support.

## 🤖 Okteto AI Agent Fleet Capabilities Demonstrated

This repository showcases how the Okteto AI Agent Fleet excels at cloud-native development:

### ✅ **Infrastructure-Aware Development**
- Automatically integrated PostgreSQL database with proper connection handling
- Created production-ready Docker Compose configuration with health checks
- Generated Okteto manifest for seamless Kubernetes deployment

### ✅ **Full Development Lifecycle Automation**
- **Build**: Updated Dockerfile with proper dependency management (`go mod tidy`)
- **Deploy**: Deployed live application with `okteto deploy --wait`
- **Test**: Created comprehensive testing scripts and validated all endpoints

### ✅ **Live Infrastructure Integration**
- **Real-time Deployment**: Changes deployed instantly to Kubernetes
- **Database Integration**: PostgreSQL running in production with persistent storage
- **Public Endpoints**: Automatically configured ingress with SSL certificates

### ✅ **Cloud-Native Best Practices**
- Environment variable configuration for database connections
- Health checks and service dependencies in Docker Compose
- Proper error handling and database connection retry logic
- Volume persistence for data storage and backup

---

## Application Overview

A REST API for managing todo items that demonstrates the evolution from simple in-memory storage to production-ready PostgreSQL integration - all automated by the Okteto AI Agent Fleet.

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

## 🎬 Try More Demos

This repository includes **5 additional demo prompts** that showcase different Okteto AI Agent Fleet capabilities:

1. **🏗️ Microservices Architecture** - Split into user service, todo service, and API gateway
2. **📊 Observability & Monitoring** - Add Prometheus, Grafana, and distributed tracing  
3. **🔒 Security & Compliance** - Implement JWT auth, RBAC, and security scanning
4. **🚀 CI/CD & GitOps** - Automated pipelines with multi-environment deployments
5. **📈 Performance & Scaling** - Auto-scaling, caching, and performance optimization

**[View All Demo Prompts →](./DEMO_PROMPTS.md)**

---

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
