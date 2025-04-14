# My DDD Go App

This project is a simple implementation of Domain-Driven Design (DDD) using Go. It focuses on user management and demonstrates how to structure a Go application following DDD principles.

## Project Structure

```
ddd-go
├── cmd
│   └── main.go                # Entry point of the application
├── internal
│   ├── domain
│   │   ├── entities
│   │   │   └── user.go        # User entity definition
│   │   ├── repositories
│   │   │   └── user_repository.go # User repository interface
│   │   └── services
│   │       └── user_service.go # User service for business logic
│   ├── infrastructure
│   │   ├── database
│   │   │   └── db.go          # Database connection and migration logic
│   │   └── http
│   │       └── handler.go     # HTTP handlers for user-related endpoints
│   └── application
│       └── usecases
│           └── user_usecase.go # Use cases for user management
├── go.mod                      # Module definition
├── go.sum                      # Module dependency checksums
└── README.md                   # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone https://github.com/yourusername/my-ddd-go-app.git
   cd my-ddd-go-app
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Run the application:**
   ```
   go run cmd/main.go
   ```

## Usage

Once the application is running, you can interact with the user-related endpoints. For example, you can create a new user or retrieve user information through the defined HTTP handlers.

## Domain-Driven Design Approach

This application follows the principles of Domain-Driven Design by organizing the code into distinct layers:

- **Domain Layer:** Contains the core business logic, including entities, repositories, and services.
- **Application Layer:** Orchestrates the use cases and interacts with the domain layer.
- **Infrastructure Layer:** Handles external concerns such as database connections and HTTP requests.

By adhering to these principles, the application remains modular, maintainable, and scalable.