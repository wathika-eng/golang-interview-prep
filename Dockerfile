# ===========================
# 🚀 Builder Stage
# ===========================
FROM golang:1.24.1-alpine3.21 AS builder

# Environment variables for Go build
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Install air (for live-reload in development environments)
RUN go install github.com/air-verse/air@latest

# Copy the entire project
COPY . .

# Build the application
RUN go build -ldflags="-s -w" -o main .

# ===========================
# 🏃‍♂️ Runtime Stage
# ===========================
FROM alpine:3.21

# Set working directory
WORKDIR /app

# Copy the binary from the builder stage to the runtime image
COPY --from=builder /app/main /app/main

# Copy the .env file into the container (runtime only)
COPY .env .env

# Install required dependencies (optional, if needed)
RUN apk --no-cache add ca-certificates

# Create a non-root user
RUN adduser -D -s /bin/sh appuser
USER appuser

# Expose port 8080 (or your app's port)
EXPOSE 8080

# Command to run the application
CMD ["/app/main"]
