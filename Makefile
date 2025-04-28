.PHONY: run build test clean migrate seed

# Variables
DB_NAME=movie_res
DB_USER=postgres
DB_PASSWORD=postgres
DB_HOST=localhost
DB_PORT=5432

# Build the application
build:
	go build -o bin/movie-res cmd/api/main.go

# Run the application
run:
	go run cmd/api/main.go

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Create database
db-create:
	createdb $(DB_NAME)

# Drop database
db-drop:
	dropdb $(DB_NAME)

# Migrate database
migrate:
	go run cmd/api/main.go migrate

# Seed database
seed:
	go run cmd/api/main.go seed

# Install dependencies
deps:
	go mod download

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Help
help:
	@echo "Available commands:"
	@echo "  make build      - Build the application"
	@echo "  make run        - Run the application"
	@echo "  make test       - Run tests"
	@echo "  make clean      - Clean build artifacts"
	@echo "  make db-create  - Create database"
	@echo "  make db-drop    - Drop database"
	@echo "  make migrate    - Run database migrations"
	@echo "  make seed       - Seed database with initial data"
	@echo "  make deps       - Install dependencies"
	@echo "  make fmt        - Format code"
	@echo "  make lint       - Run linter"
	@echo "  make help       - Show this help message" 