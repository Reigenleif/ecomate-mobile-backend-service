export APP_NAME=SimpleApp
export SERVER_PORT=9000
export SERVER_READ_TIMEOUT_IN_MINUTE=2
export SERVER_WRITE_TIMEOUT_IN_MINUTE=2
export DATABASE_URL=postgres://postgres:postgres@localhost:5432/emoji?sslmode=disable
go run cmd/web/*.go