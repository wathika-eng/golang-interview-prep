# Description: Makefile for the API
@PHONY: run build clean

# Run the API
run:
	@echo "Running..."
	@go mod tidy
	@go run main.go

# Build the binary
build:
	@echo "Building..."
	@go mod tidy
	@go build -o main main.go
	@echo "Build complete"

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -rf main
	@echo "Clean complete"


# Live Reload
watch:
	@go mod tidy
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

.PHONY: all build run test clean watch