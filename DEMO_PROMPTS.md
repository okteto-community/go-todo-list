# Okteto AI Agent Fleet Demo Prompts

This repository demonstrates the power of the **Okteto AI Agent Fleet** for cloud-native development. Here are 5 compelling demo prompts that showcase different capabilities:

## 🎯 Demo Prompt #1: Database Migration
**Scenario**: Transform in-memory storage to persistent database

```
"Update my todo list application so that, instead of storing the items in memory, 
they are stored on a database. Use postgres, deploy it as a container, and then 
create a docker compose file so I can deploy everything with one command"
```

**What the Agent Demonstrates**:
- ✅ Automatically integrates PostgreSQL with proper Go driver
- ✅ Updates all CRUD operations to use SQL queries
- ✅ Creates Docker Compose with health checks and dependencies
- ✅ Generates Okteto manifest for Kubernetes deployment
- ✅ Provides live working endpoint with database persistence

---

## 🎯 Demo Prompt #2: Microservices Architecture
**Scenario**: Break monolith into microservices

```
"Split this todo application into microservices - create a separate user service 
for authentication, a todo service for CRUD operations, and an API gateway. 
Deploy everything with proper service discovery and load balancing."
```

**What the Agent Demonstrates**:
- ✅ Designs microservices architecture with proper separation of concerns
- ✅ Creates service-to-service communication patterns
- ✅ Implements API gateway with routing and load balancing
- ✅ Sets up service discovery and inter-service networking
- ✅ Deploys complete microservices stack to Kubernetes

---

## 🎯 Demo Prompt #3: Observability & Monitoring
**Scenario**: Add production-ready monitoring

```
"Add comprehensive monitoring to this application - include metrics, logging, 
tracing, and health checks. Set up Prometheus, Grafana, and alerting. 
I want to see real-time performance dashboards."
```

**What the Agent Demonstrates**:
- ✅ Integrates Prometheus metrics collection
- ✅ Sets up Grafana dashboards with real-time data
- ✅ Implements distributed tracing with Jaeger
- ✅ Configures structured logging and log aggregation
- ✅ Creates alerting rules and notification channels

---

## 🎯 Demo Prompt #4: Security & Compliance
**Scenario**: Implement enterprise security

```
"Secure this application for production use - add JWT authentication, 
RBAC authorization, API rate limiting, input validation, and security scanning. 
Also implement HTTPS, secrets management, and compliance logging."
```

**What the Agent Demonstrates**:
- ✅ Implements JWT-based authentication system
- ✅ Creates role-based access control (RBAC)
- ✅ Adds API rate limiting and DDoS protection
- ✅ Integrates secrets management with Kubernetes secrets
- ✅ Sets up security scanning and vulnerability assessment

---

## 🎯 Demo Prompt #5: CI/CD & GitOps
**Scenario**: Automate deployment pipeline

```
"Set up a complete CI/CD pipeline for this application with automated testing, 
security scanning, multi-environment deployments, and GitOps. Include staging 
and production environments with proper promotion workflows."
```

**What the Agent Demonstrates**:
- ✅ Creates GitHub Actions workflows for CI/CD
- ✅ Implements automated testing (unit, integration, e2e)
- ✅ Sets up multi-environment deployments (dev/staging/prod)
- ✅ Configures GitOps with ArgoCD or Flux
- ✅ Implements blue-green or canary deployment strategies

---

## 🌟 Why Okteto AI Agent Fleet is Superior

### **Cloud-Native Expertise**
- Deep understanding of Kubernetes, Docker, and cloud-native patterns
- Knows best practices for microservices, observability, and security
- Automatically applies production-ready configurations

### **Infrastructure Integration**
- Direct integration with Okteto platform and Kubernetes clusters
- Real-time deployment and testing capabilities
- Live endpoint generation with SSL certificates

### **Full Development Lifecycle**
- Handles everything from code changes to production deployment
- Automated testing, building, and deployment workflows
- Comprehensive documentation and change tracking

### **Live Demonstration**
- Provides working endpoints you can test immediately
- Real-time infrastructure changes visible in Kubernetes
- Browser-based IDE for code inspection and interaction

---

## 🚀 Try These Demos

1. **Fork this repository**
2. **Open in Okteto AI Agent Fleet**
3. **Use any of the demo prompts above**
4. **Watch the agent work its magic!**

Each demo showcases different aspects of cloud-native development, from basic database integration to complex enterprise architectures. The agent handles all the complexity while providing live, working results you can test immediately.

---

## 📊 Demo Results

After running Demo Prompt #1, the agent delivered:

- **Live Application**: https://todo-app-agent-quh1cdulyz9i.rberrelleza.fleets.okteto.ai
- **Database Integration**: PostgreSQL with persistent storage
- **Docker Compose**: Ready for local development
- **Okteto Deployment**: Production-ready Kubernetes manifests
- **Comprehensive Testing**: Automated test scripts and validation

**Total Time**: ~5 minutes from prompt to live application!

This demonstrates the power of having an AI agent that truly understands cloud-native development and can deliver production-ready results instantly.