# User Management Service

A clean, modular RESTful API built with **Go**, **Gin** (Web Framework), and **GORM** (ORM). This service provides user management capabilities including registration, retrieval, and password hashing, using a layered clean architecture approach.

## Key Features
- **User Management**: Registration, profile retrieval (by ID), and listing all users.
- **Security**: Secure password hashing using `bcrypt`.
- **Modular Design**: Implements clean layering (Handler → Service → Repository) with interface-based dependency injection.
- **Data Validation**: Struct-based validation using Gin binding tags.
- **Database**: Automatic schema migration via GORM and MySQL.

## Architecture (Clean Architecture Layers)
- **`api/internal/config`**: Configuration management using `.env` via `godotenv`.
- **`api/internal/database`**: MySQL connectivity and GORM setup.
- **`api/internal/model`**: Domain entities with GORM tags (`User`).
- **`api/internal/repository`**: Interface-based data access layer (GORM implementation).
- **`api/internal/service`**: Business logic layer — bcrypt hashing, duplicate email validation.
- **`api/internal/handler`**: HTTP handlers with Gin binding validation.
- **`api/internal/router`**: Gin router with `/api/v1` prefix and route grouping.

**Dependency Flow:**
```
main.go → config → database (AutoMigrate) → repository → service → handler → router → Run()
```

## Getting Started

### Prerequisites
- Go 1.25+
- MySQL database

### Installation
1. Clone the repository.
2. Create a `.env` file in the root directory:
   ```env
   DB_HOST=127.0.0.1
   DB_PORT=3306
   DB_USER=root
   DB_PASSWORD=yourpassword
   DB_NAME=your_db
   APP_PORT=8080
   ```
3. Run the application:
   ```bash
   go run api/main.go
   ```

## Development
- **Build**: `go build -o bin/api ./api`
- **Dependencies**: Uses Go modules.
