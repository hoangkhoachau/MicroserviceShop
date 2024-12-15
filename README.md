# E-Shop Microservices

This project is a comprehensive e-commerce platform designed with event-driven microservice architecture. It demonstrates best practices for building scalable, maintainable, and loosely coupled services.

## Overview

The system consists of several autonomous microservices that communicate asynchronously via events, providing resilience and flexibility. Each service is responsible for a specific business capability and can be developed, deployed, and scaled independently.

## Services

### API Gateway
- Functions as the single entry point for all client requests
- Routes requests to appropriate microservices
- Handles cross-cutting concerns like authentication
- Retrieves service addresses dynamically from Consul service discovery

### Identity Service
- Manages user authentication and authorization
- Handles user registration, login, and role management
- Issues JWT tokens for authentication
- Uses MySQL for persistent storage

### Catalog Service
- Manages product inventory and catalog
- Provides product search, filtering, and browsing capabilities
- Handles categories, subcategories, and brands
- Uses MySQL for relational data storage

### Basket Service
- Maintains user shopping baskets
- Provides basket manipulation operations (add/remove items)
- Verifies baskets before order creation
- Uses MongoDB for document storage

### Order Service
- Processes and manages customer orders
- Handles order fulfillment and status tracking
- Provides order history for users
- Uses MongoDB for document storage

### Email Service
- Handles all transactional email communications
- Sends verification emails, order confirmations, etc.
- Operates purely in response to events from other services

## Event-Driven Architecture

The services communicate with each other via an event bus implemented with RabbitMQ, enabling loose coupling and scalability.

![event bus architecture diagram](https://user-images.githubusercontent.com/64227421/208158347-a457a3f0-1bbd-4dd0-b064-ec76270f29fb.png)

### Event Flow

|         Event       |     Sender      |        Sender Description     |          Receivers            | Receiver Description
| ------------------- | --------------- | ----------------------------- | ----------------------------- | --------------------------------
| UserCreatedEvent    | identityService | Sent when new user is created | emailService, basketService   | A verification code is sent to the email of the created user and a basket is created for the user
| BasketVerifiedEvent | basketService   | Sent when basket is verified  | orderService                  | An order is created with the products in the verified basket.
| OrderCompletedEvent | orderService    | Sent when order is complete   | catalogService, basketService | After the order is completed, the number of products is reduced from stock and the basket is emptied   

## Service Discovery

Consul is used for service discovery:
- Each service registers itself with Consul on startup
- API Gateway queries Consul to locate services
- Provides health checking and service location transparency
- Enables dynamic scaling of services

## Database Strategy

The project employs a polyglot persistence approach:
- **MySQL**: Used by Identity and Catalog services for structured relational data
- **MongoDB**: Used by Order and Basket services for flexible document storage
- **Redis**: Used for caching to improve performance where applicable

## Technologies

* Go - Primary programming language
* RabbitMQ - Event bus for asynchronous communication
* MongoDB - Document database for Order and Basket services
* MySQL - Relational database for Identity and Catalog services
* Redis - In-memory data structure store for caching
* Consul - Service discovery and registration
* Docker - Containerization and deployment

## Prerequisites

- Docker and Docker Compose
- Go 1.16+ (for development)
- Make (optional, for development)

## Setup and Installation

### Running with Docker Compose

To start all services with their dependencies:

```
docker-compose up -d
```

This will:
1. Build all service images
2. Set up required networks
3. Start all containers in detached mode
4. Initialize databases if needed

### Configuration

Each service has its own configuration in the `config` directory. Main configuration parameters:
- Database connections
- RabbitMQ settings
- Service discovery settings
- JWT token configuration

### Development Setup

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/e-shop-microservices.git
   cd e-shop-microservices
   ```

2. Install dependencies for each service

3. For local development without Docker:
   - Set up local instances of MySQL, MongoDB, RabbitMQ, Redis, and Consul
   - Configure each service to connect to your local instances
   - Start each service individually

## API Documentation

The API Gateway exposes RESTful endpoints for all services. Main endpoints include:

- `/api/identity/` - Authentication and user management
- `/api/catalog/` - Product catalog operations
- `/api/basket/` - Shopping basket operations
- `/api/order/` - Order processing

Detailed API documentation can be generated using Swagger.
