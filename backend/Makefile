all: build test

build:
	@echo "Building..."
	@go build -o main cmd/api/main.go

run:
	@go run cmd/api/main.go

test:
	@echo "Testing..."
	@go test ./... -v

itest:
	@echo "Running integration tests..."
	@go test ./internal/database -v

clean:
	@echo "Cleaning..."
	@rm -f main

watch:
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

gen-graphs:
	@go run github.com/99designs/gqlgen generate

.PHONY: all build run test clean watch itest
