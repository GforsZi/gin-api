# Gin API Project — User Management Service

## Architecture (Clean Architecture Layers)
- **Framework:** Gin (Web), GORM (Database/ORM), bcrypt (Password Hashing)
- **Flow:** `main.go` → `config.Load()` → `database.Connect(cfg)` → `AutoMigrate` → `repository` → `service` → `handler` → `router` → `r.Run()`
- **Structure:**
  - `api/main.go`: Application entrypoint; wires all layers together.
  - `api/internal/config`: Configuration loader using `godotenv` (`.env` file).
  - `api/internal/database`: MySQL connection setup via GORM.
  - `api/internal/model`: Domain models with GORM tags (e.g., `User`).
  - `api/internal/repository`: Data access layer (interface + GORM implementation).
  - `api/internal/service`: Business logic layer (interface + implementation); includes bcrypt hashing & duplicate email check.
  - `api/internal/handler`: HTTP handlers with Gin; input validation via binding tags.
  - `api/internal/router`: Gin router setup with `/api/v1` prefix.
- **API Endpoints:**
  - `POST /api/v1/users/register` — Register a new user
  - `GET  /api/v1/users/:id` — Get user by ID
  - `GET  /api/v1/users` — List all users

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
  ```

## Development Notes
- **Env:** Use a `.env` file for `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, and `APP_PORT`.
- **Dependencies:** Gin, GORM (MySQL driver), godotenv, bcrypt, validator.
- **Paths:** All source code under `api/` directory.
