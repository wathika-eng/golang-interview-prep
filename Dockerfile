# Builder stage
FROM golang:1.23-alpine3.21 AS builder

# Environment variables for Go build
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 

# Set the current working directory inside the container, build and copy the application
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Install air (for live-reload in development environments)
RUN go install github.com/air-verse/air@latest

# Copy all the files into the image
COPY . .

# Build the application
RUN go build -ldflags="-s -w" -o main .

# Final runtime stage
FROM alpine:3.21

# Copy the binary from the builder stage to the runtime image
COPY --from=builder /app/main /app/main

# Copy the .env file into the container (do this only for runtime, not in the builder)
COPY .env .env

# Optionally create a non-root user
RUN adduser -D -s /bin/sh appuser
USER appuser

# Expose port 8080 (or whatever port your app uses)
EXPOSE 8080

# Command to run the application
CMD ["/app/main"]
