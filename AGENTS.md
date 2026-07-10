# Gin API Project — User Management Service

## Architecture (Clean Architecture Layers)
- **Framework:** Gin (Web), GORM (Database/ORM), bcrypt (Password Hashing), **Firebase Admin SDK (Authentication)**
- **Flow:** `main.go` → `config.Load()` → `database.Connect(cfg)` → `AutoMigrate` → `Firebase Init` → `repository` → `service` → `handler` → `router` → `r.Run()`
- **Structure:**
  - `api/main.go`: Application entrypoint; wires all layers together.
  - `api/internal/config`: Configuration loader using `godotenv` (`.env` file).
  - `api/internal/database`: MySQL connection setup via GORM.
  - `api/internal/firebase`: Firebase Admin SDK initialization (Auth client).
  - `api/internal/model`: Domain models with GORM tags (e.g., `User` includes `FirebaseUID`).
  - `api/internal/repository`: Data access layer (interface + GORM implementation); includes `FindByFirebaseUID`.
  - `api/internal/service`: Business logic layer (interface + implementation); includes bcrypt hashing, duplicate email check, **Firebase token verification & auto-register**.
  - `api/internal/handler`: HTTP handlers with Gin; input validation via binding tags.
  - `api/internal/middleware`: Gin middleware (e.g., `AuthMiddleware` for Firebase token verification).
  - `api/internal/router`: Gin router setup with `/api/v1` prefix.
- **API Endpoints:**
  - `POST /api/v1/users/register` — Register a new user (bcrypt)
  - `GET  /api/v1/users` — List all users (public)
  - `GET  /api/v1/users/:id` — Get user by ID (public)
  - `POST /api/v1/auth/firebase` — Authenticate via Firebase ID token (auto-register if new)
  - `GET  /api/v1/auth/me` — Get current authenticated user (requires Bearer token)

## Build & Run
- **Run:** `go run api/main.go`
- **Build:** `go build -o bin/api ./api`
- **Database:** MySQL, configured via `.env`:
  ```env
  DB_HOST=127.0.0.1
  DB_PORT=3306
  DB_USER=root
  DB_PASSWORD=yourpassword
  DB_NAME=your_db
  APP_PORT=8080
  FIREBASE_CREDENTIALS_PATH=firebase-credentials.json
  ```

## Development Notes
- **Env:** Use a `.env` file for `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, `APP_PORT`, and `FIREBASE_CREDENTIALS_PATH`.
- **Dependencies:** Gin, GORM (MySQL driver), godotenv, bcrypt, validator, Firebase Admin SDK.
- **Firebase:** Service account JSON file must be in project root (path set via `FIREBASE_CREDENTIALS_PATH`). Already in `.gitignore`.
- **Paths:** All source code under `api/` directory.
