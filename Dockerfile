# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install git for private repositories
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod ./

# Initialize module and download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary and .env from builder
COPY --from=builder /app/main .
COPY .env .

# Expose port
EXPOSE 8080

# Command to run the application
CMD ["./main"] 