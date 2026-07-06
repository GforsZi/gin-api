# Gin API Project

## Architecture
- **Framework:** Gin (Web), GORM (Database/ORM)
- **Structure:**
  - `api/main.go`: Application entrypoint.
  - `api/internal/config`: Configuration loader using `godotenv`.
  - `api/internal/database`: Database connection setup (MySQL).
  - `api/internal/model`: Domain models (e.g., `User`).
  - `api/internal/repository`: Data access layer using interfaces.

## Build & Run
- **Run:** `go run api/main.go`
- **Build:** `go build -o bin/api ./api`
- **Database:** Project expects a MySQL database configured via environment variables (e.g., `.env` file).

## Development Notes
- **Env:** Use a `.env` file for `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, and `APP_PORT`.
- **Dependencies:** Uses GORM for database interactions.
- **Paths:** Source code is located in the `api/` directory.
