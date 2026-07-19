# Gin Api

A clean, modular RESTful API built with **Go**, **Gin** (Web Framework), and **GORM** (ORM). This service provides user management capabilities including registration, retrieval, and password hashing, using a layered clean architecture approach.

## Key Features
- **User Management**: Registration, profile retrieval (by ID), and listing all users.
- **Security**: Secure password hashing using `bcrypt`.
- **Modular Design**: Implements clean layering (Handler → Service → Repository) with interface-based dependency injection.
- **Data Validation**: Struct-based validation using Gin binding tags.
- **Database**: Automatic schema migration via GORM and MySQL.
- **Firebase Authentication**: Token-based auth via Firebase Admin SDK (auto-register on first login).
- **Dual Auth Mode**: Supports both bcrypt registration and Firebase authentication.

## Architecture (Clean Architecture Layers)
- **`api/internal/config`**: Configuration management using `.env` via `godotenv`.
- **`api/internal/database`**: MySQL connectivity and GORM setup.
- **`api/internal/firebase`**: Firebase Admin SDK initialization (Auth client).
- **`api/internal/model`**: Domain entities with GORM tags (`User` — includes `FirebaseUID`).
- **`api/internal/repository`**: Interface-based data access layer (GORM implementation).
- **`api/internal/service`**: Business logic layer — bcrypt hashing, duplicate email validation, Firebase token verification.
- **`api/internal/handler`**: HTTP handlers with Gin binding validation.
- **`api/internal/middleware`**: Gin middleware for Firebase Bearer token verification.
- **`api/internal/router`**: Gin router with `/api/v1` prefix and route grouping.

**Dependency Flow:**
```
main.go → config → database (AutoMigrate) → Firebase Init → repository → service → handler → router → Run()
```

## Getting Started

### Prerequisites
- Go 1.25+
- MySQL database
- Firebase project with Authentication enabled
- Firebase Admin SDK service account JSON key

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
   FIREBASE_CREDENTIALS_PATH=firebase-credentials.json
   ```
3. Download your Firebase service account key from **Project Settings → Service Accounts** and save it as the path above.
4. Run the application:
   ```bash
   go run api/main.go
   ```

## API Endpoints

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| `POST` | `/api/v1/users/register` | ❌ | Register with name, email, password (bcrypt) |
| `GET` | `/api/v1/users` | ❌ | List all users |
| `GET` | `/api/v1/users/:id` | ❌ | Get user by ID |
| `POST` | `/api/v1/auth/firebase` | ❌ | Authenticate with Firebase ID token (auto-register) |
| `GET` | `/api/v1/auth/me` | ✅ Bearer | Get current authenticated user profile |

## Testing Firebase Auth

### 1. Create a test user in Firebase Console
**Authentication → Users → Add user** (set email & password).

### 2. Get your Web API Key
**Project Settings → General → Your apps → Web API Key**.

### 3. Login via Firebase REST API
```bash
curl -X POST "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=YOUR_WEB_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"email":"test@user.com","password":"123456","returnSecureToken":true}'
```
Save the `idToken` from the response.

### 4. Test backend endpoints
```bash
# Login / auto-register
curl -X POST http://localhost:8080/api/v1/auth/firebase \
  -H "Content-Type: application/json" \
  -d '{"id_token":"YOUR_ID_TOKEN"}'

# Get current user
curl http://localhost:8080/api/v1/auth/me \
  -H "Authorization: Bearer YOUR_ID_TOKEN"
```

## Development
- **Build**: `go build -o bin/api ./api`
- **Dependencies**: Uses Go modules.
