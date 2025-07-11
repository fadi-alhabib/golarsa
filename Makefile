# GoLarsa Makefile
.PHONY: build install clean test help

# Variables
BINARY_NAME=golarsa
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS=-ldflags "-X main.version=$(VERSION)"

# Default target
all: build

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	go build $(LDFLAGS) -o $(BINARY_NAME) .

# Install globally
install:
	@echo "Installing $(BINARY_NAME) globally..."
	go install $(LDFLAGS) .

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -f $(BINARY_NAME)
	rm -rf pkg/services/

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Lint code (optional - requires golangci-lint to be installed)
lint:
	@echo "Linting code..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Install it with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

# Check for vulnerabilities (optional - requires govulncheck)
vulncheck:
	@echo "Checking for vulnerabilities..."
	@if command -v govulncheck >/dev/null 2>&1; then \
		govulncheck ./...; \
	else \
		echo "govulncheck not installed. Install it with: go install golang.org/x/vuln/cmd/govulncheck@latest"; \
	fi

# Tidy dependencies
tidy:
	@echo "Tidying dependencies..."
	go mod tidy

# Build for multiple platforms
build-all:
	@echo "Building for multiple platforms..."
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-linux-amd64 .
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-arm64 .
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-windows-amd64.exe .

# Create release directory
dist:
	mkdir -p dist

# Release build (requires dist directory)
release: dist build-all
	@echo "Creating release builds..."
	@echo "Binaries created in dist/ directory"

# Test the CLI by creating a sample service
test-cli:
	@echo "Testing CLI functionality..."
	@rm -rf test-output
	@mkdir -p test-output
	@cd test-output && go mod init test-module && ../$(BINARY_NAME) service sample
	@echo "✓ CLI test completed successfully"
	@rm -rf test-output

# Help target
help:
	@echo "Available targets:"
	@echo "  build         - Build the binary"
	@echo "  install       - Install globally using go install"
	@echo "  clean         - Clean build artifacts"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage"
	@echo "  fmt           - Format code"
	@echo "  lint          - Lint code (optional, requires golangci-lint)"
	@echo "  vulncheck     - Check for vulnerabilities (optional, requires govulncheck)"
	@echo "  tidy          - Tidy dependencies"
	@echo "  build-all     - Build for multiple platforms"
	@echo "  release       - Create release builds"
	@echo "  test-cli      - Test CLI functionality"
	@echo "  help          - Show this help message" 