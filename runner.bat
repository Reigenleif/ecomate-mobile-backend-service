set APP_NAME=SimpleApp
set SERVER_PORT=9000
set SERVER_READ_TIMEOUT_IN_MINUTE=2
set SERVER_WRITE_TIMEOUT_IN_MINUTE=2
set DATABASE_URL=postgres://postgres:postgres@localhost:5432/ecomate-test?sslmode=disable
set JWT_SECRET="mZu5Z63yXOFeD8nNoLJJGBIf3/Y3X8PUj3xzjxev//A="
go run main.go