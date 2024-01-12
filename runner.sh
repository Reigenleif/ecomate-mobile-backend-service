export APP_NAME=SimpleApp
export SERVER_PORT=9000
export SERVER_READ_TIMEOUT_IN_MINUTE=2
export SERVER_WRITE_TIMEOUT_IN_MINUTE=2
export JWT_SECRET="mZu5Z63yXOFeD8nNoLJJGBIf3/Y3X8PUj3xzjxev//A="
export DATABASE_URL=postgres://postgres:arexia@34.134.124.160:5432/affable-doodad-353108:us-central1:ecomate-db?sslmode=disable
go run cmd/web/*.go