# Description: Dockerfile for building the Go application
# Builder stage
FROM golang:1.23-alpine3.21 AS builder

#env vars for GOBUILD
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 

# Set the current working directory inside the container, build and copy the application
WORKDIR /app
# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./

# install air
RUN go install github.com/air-verse/air@latest

# Download dependencies
RUN go mod download
# source to copy (all files .) and build directory .
COPY . .

# COPY .env .env
# Build the application
RUN go build -ldflags="-s -w" -o main .

# add a user (non-root)
# RUN adduser -D -s /bin/sh appuser

# stage 2
# distro image
FROM alpine:3.21
#copy the binary from the builder stage to the distro image
COPY --from=builder /app/main main
COPY .env .env
#set user
# USER appuser

EXPOSE 8080

# command to run the application with live reload
CMD ["./main"]