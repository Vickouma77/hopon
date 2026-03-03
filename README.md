# HopOn 🚗

> Get in loser, we're riding

A modern, scalable rideshare platform built with microservices architecture, enabling seamless connections between riders and drivers.

## 📋 Table of Contents

- [Overview](#overview)
- [Architecture](#architecture)
- [Tech Stack](#tech-stack)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Development](#development)
- [Services](#services)
- [API Documentation](#api-documentation)
- [Deployment](#deployment)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## 🎯 Overview

HopOn is a comprehensive rideshare platform that provides:

- **Real-time Matching**: Connect riders with nearby drivers instantly
- **Live Tracking**: Track rides in real-time with GPS integration
- **Payment Processing**: Secure payment handling with multiple payment methods
- **Rating System**: Two-way rating system for riders and drivers
- **Route Optimization**: Intelligent routing for efficient pickups and drop-offs
- **Push Notifications**: Real-time updates for ride status changes
- **Admin Dashboard**: Complete oversight and management capabilities

## 🏗️ Architecture

HopOn follows a microservices architecture pattern with the following key components:

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Web App   │     │ Mobile Apps │     │   Admin     │
│ (TypeScript)│     │(iOS/Android)│     │  Dashboard  │
└──────┬──────┘     └──────┬──────┘     └──────┬──────┘
       │                   │                   │
       └───────────────────┼───────────────────┘
                          │
                   ┌──────▼──────┐
                   │  API Gateway │
                   │     (Go)     │
                   └──────┬───────┘
                          │
       ┌──────────────────┼──────────────────┐
       │                  │                  │
   ┌───▼───┐         ┌────▼────┐       ┌────▼────┐
   │ User  │         │  Ride   │       │ Payment │
   │Service│         │ Service │       │ Service │
   │  (Go) │         │   (Go)  │       │   (Go)  │
   └───┬───┘         └────┬────┘       └────┬────┘
       │                  │                  │
       └──────────────────┼──────────────────┘
                          │
                   ┌──────▼──────┐
                   │   MongoDB   │
                   │   Cluster   │
                   └─────────────┘
```

### Key Architectural Decisions

- **Microservices**: Independent services for better scalability and maintainability
- **Event-Driven**: Asynchronous communication using message queues
- **Cloud-Native**: Kubernetes orchestration for container management
- **API-First**: RESTful APIs with OpenAPI/Swagger documentation
- **Database per Service**: Each microservice owns its data

## 🛠️ Tech Stack

### Backend
- **Go (Golang)**: High-performance backend services
- **gRPC**: Internal service communication
- **REST APIs**: External client communication
- **MongoDB**: Primary database for all services
- **Redis**: Caching and session management
- **RabbitMQ/Kafka**: Message queue for event-driven architecture

### Frontend
- **TypeScript**: Type-safe frontend development
- **React**: Web application framework
- **React Native**: Cross-platform mobile apps
- **Redux/Context API**: State management
- **WebSocket**: Real-time communication

### Infrastructure & DevOps
- **Docker**: Containerization
- **Kubernetes**: Container orchestration
- **Tilt.dev**: Local Kubernetes development
- **Helm**: Kubernetes package management
- **Terraform**: Infrastructure as Code
- **GitHub Actions**: CI/CD pipelines

### Observability
- **Prometheus**: Metrics collection
- **Grafana**: Metrics visualization
- **Jaeger**: Distributed tracing
- **ELK Stack**: Centralized logging

### Additional Tools
- **Nginx**: Reverse proxy and load balancing
- **JWT**: Authentication and authorization
- **Stripe/daraja/buni**: Payment processing
- **Google Maps API**: Location and routing services
- **Firebase**: Push notifications

## 📦 Prerequisites

Before you begin, ensure you have the following installed:

- **Go**
- **Node.js** 
- **Docker**
- **Kubernetes** (Docker Desktop, minikube, or kind)
- **kubectl**
- **Tilt**
- **MongoDB** 
- **Git**

### Optional but Recommended
- **Helm**
- **Make**
- **k9s** (Kubernetes CLI manager)

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/hopon.git
cd hopon
```

### 2. Environment Setup

Copy the example environment files and configure them:

```bash
# Copy environment templates
cp .env.example .env
cp config/dev.yaml.example config/dev.yaml

# Edit the configuration files with your settings
nano .env
```

Required environment variables:
```env
# Database
MONGO_URI=mongodb://localhost:27017
MONGO_DB_NAME=hopon

# Redis
REDIS_HOST=localhost:6379

# JWT
JWT_SECRET=your-super-secret-key
JWT_EXPIRY=24h

# External APIs
GOOGLE_MAPS_API_KEY=your-google-maps-key
STRIPE_SECRET_KEY=your-stripe-key

# Service Ports
USER_SERVICE_PORT=8001
RIDE_SERVICE_PORT=8002
PAYMENT_SERVICE_PORT=8003
NOTIFICATION_SERVICE_PORT=8004
```

### 3. Local Development with Tilt

Tilt provides a seamless local Kubernetes development experience:

```bash
# Start all services with Tilt
tilt up

# Access the Tilt UI (opens automatically)
# Visit: http://localhost:10350
```

Tilt will:
- Build all Docker images
- Deploy services to your local Kubernetes cluster
- Set up hot-reloading for code changes
- Forward necessary ports
- Display logs from all services

### 4. Manual Setup (Without Tilt)

If you prefer to run services individually:

```bash
# Start MongoDB
docker-compose up -d mongodb redis

# Install Go dependencies
cd services/user-service
go mod download

# Run a service
go run cmd/server/main.go

# In another terminal, start the frontend
cd web-app
npm install
npm run dev
```

## 💻 Development

### Project Structure

```
hopon/
├── services/                 # Microservices
│   ├── user-service/        # User management
│   ├── ride-service/        # Ride matching and tracking
│   ├── payment-service/     # Payment processing
│   ├── notification-service/# Notifications
│   └── admin-service/       # Admin operations
├── web-app/                 # React web application
├── mobile-app/              # React Native mobile app
├── api-gateway/             # API Gateway
├── k8s/                     # Kubernetes manifests
│   ├── base/               # Base configurations
│   ├── dev/                # Development overlays
│   └── prod/               # Production overlays
├── helm/                    # Helm charts
├── scripts/                 # Utility scripts
├── docs/                    # Documentation
├── proto/                   # gRPC protocol definitions
├── Tiltfile                 # Tilt configuration
├── docker-compose.yml       # Local dependencies
└── README.md
```

### Go Service Structure

Each Go microservice follows this structure:

```
service-name/
├── cmd/
│   └── server/
│       └── main.go          # Entry point
├── internal/
│   ├── handler/            # HTTP/gRPC handlers
│   ├── service/            # Business logic
│   ├── repository/         # Data access layer
│   ├── model/              # Data models
│   └── middleware/         # Middleware functions
├── pkg/                     # Public packages
├── config/                  # Configuration files
├── Dockerfile
├── go.mod
└── go.sum
```

### TypeScript/React Structure

```
web-app/
├── src/
│   ├── components/         # React components
│   ├── pages/              # Page components
│   ├── services/           # API services
│   ├── store/              # State management
│   ├── hooks/              # Custom hooks
│   ├── types/              # TypeScript types
│   ├── utils/              # Utility functions
│   └── App.tsx
├── public/
├── package.json
└── tsconfig.json
```

### Development Workflow

1. **Create a feature branch**
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes**
   - Write clean, documented code
   - Follow the existing code style
   - Add tests for new functionality

3. **Run tests**
   ```bash
   # Go services
   go test ./...
   
   # Frontend
   npm test
   ```

4. **Commit and push**
   ```bash
   git add .
   git commit -m "feat: add your feature description"
   git push origin feature/your-feature-name
   ```

5. **Create a Pull Request**

### Code Style

- **Go**: Follow [Effective Go](https://golang.org/doc/effective_go) guidelines
- **TypeScript**: Use ESLint and Prettier configurations
- **Commits**: Follow [Conventional Commits](https://www.conventionalcommits.org/)

## 🔧 Services

### User Service
**Port**: 8001

Handles user authentication, registration, and profile management.

**Endpoints**:
- `POST /api/v1/auth/register` - Register new user
- `POST /api/v1/auth/login` - User login
- `GET /api/v1/users/:id` - Get user profile
- `PUT /api/v1/users/:id` - Update user profile

### Ride Service
**Port**: 8002

Manages ride requests, matching, and tracking.

**Endpoints**:
- `POST /api/v1/rides` - Request a ride
- `GET /api/v1/rides/:id` - Get ride details
- `PUT /api/v1/rides/:id/accept` - Driver accepts ride
- `PUT /api/v1/rides/:id/complete` - Complete ride
- `GET /api/v1/rides/driver/:id` - Get driver's rides
- `GET /api/v1/rides/rider/:id` - Get rider's rides

### Payment Service
**Port**: 8003

Processes payments and manages payment methods.

**Endpoints**:
- `POST /api/v1/payments` - Process payment
- `GET /api/v1/payments/:id` - Get payment details
- `POST /api/v1/payment-methods` - Add payment method
- `GET /api/v1/payment-methods` - List payment methods

### Notification Service
**Port**: 8004

Sends push notifications, emails, and SMS.

**Endpoints**:
- `POST /api/v1/notifications/push` - Send push notification
- `POST /api/v1/notifications/email` - Send email
- `POST /api/v1/notifications/sms` - Send SMS

## 📚 API Documentation

API documentation is available through Swagger UI:

```bash
# Start the API gateway
tilt up

# Access Swagger UI
open http://localhost:8000/swagger
```

For detailed API documentation, see [docs/api.md](docs/api.md).

## 🚢 Deployment

### Development Environment

Development deployments are handled automatically through Tilt:

```bash
tilt up
```

### Staging/Production Deployment

1. **Build and push Docker images**
   ```bash
   make docker-build-all
   make docker-push-all
   ```

2. **Deploy with Helm**
   ```bash
   # Add Helm repository
   helm repo add hopon https://charts.hopon.com
   
   # Install/Upgrade
   helm upgrade --install hopon hopon/hopon \
     --namespace production \
     --values helm/values-prod.yaml
   ```

3. **Apply Kubernetes manifests directly**
   ```bash
   kubectl apply -k k8s/prod
   ```

### CI/CD Pipeline

GitHub Actions automatically:
- Runs tests on every PR
- Builds Docker images on merge to main
- Deploys to staging automatically
- Deploys to production with manual approval

See [.github/workflows](.github/workflows) for pipeline definitions.

## 🧪 Testing

### Unit Tests

```bash
# Go services
cd services/user-service
go test ./... -v

# Frontend
cd web-app
npm test
```

### Integration Tests

```bash
# Run integration tests
make test-integration
```

### End-to-End Tests

```bash
# Start all services
tilt up

# Run E2E tests
make test-e2e
```

### Load Testing

```bash
# Using k6
k6 run tests/load/ride-booking.js
```

## 🤝 Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

### Quick Contribution Guide

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'feat: add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

### Code Review Process

- All PRs require at least one approval
- All tests must pass
- Code coverage must not decrease
- Follow the style guide

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Thanks to all contributors who have helped build HopOn
- Built with amazing open-source tools and frameworks
- Special thanks to the Go, Kubernetes, and TypeScript communities

## 📞 Support

- **Documentation**: [docs.hopon.com](https://docs.hopon.com)
- **Issues**: [GitHub Issues](https://github.com/yourusername/hopon/issues)
- **Discussions**: [GitHub Discussions](https://github.com/yourusername/hopon/discussions)
- **Email**: support@hopon.com

## 🗺️ Roadmap

- [ ] Multi-language support
- [ ] Advanced ride scheduling
- [ ] Carpool/ride sharing features
- [ ] Electric vehicle integration
- [ ] Carbon footprint tracking
- [ ] Corporate accounts
- [ ] API for third-party integrations
- [ ] Machine learning-based pricing

---

Made with ❤️ by zeroALT
