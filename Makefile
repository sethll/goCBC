# goCBC Makefile

# Variables
BINARY_NAME=goCBC
MAIN_PATH=./cmd/goCBC
BUILD_DIR=./build
PKG_PATH=github.com/sethll/goCBC/pkg/progmeta
GIT_COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
LDFLAGS=-ldflags "-X $(PKG_PATH).build=$(GIT_COMMIT)"
EXAMPLE_ARGS=75 '1030:150' '1230:200' '3215:100' '1788:100'

# Default target
.PHONY: all
all: build

# Build the binary with git commit hash
.PHONY: build
build:
	@echo "Building $(BINARY_NAME) with git commit: $(GIT_COMMIT)"
	test -d $(BUILD_DIR) || mkdir -p $(BUILD_DIR)
	go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)

# Build without git commit hash (development)
.PHONY: build-dev
build-dev:
	@echo "Building $(BINARY_NAME) (development build)"
	test -d $(BUILD_DIR) || mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME)-dev $(MAIN_PATH)

# Run built binary
.PHONY: exec-example
exec-example:
	$(BUILD_DIR)/$(BINARY_NAME) $(EXAMPLE_ARGS)

# Run the application
.PHONY: run
run:
	go run $(LDFLAGS) $(MAIN_PATH)

# Run with example arguments
.PHONY: run-example
run-example:
	go run $(LDFLAGS) $(MAIN_PATH) $(EXAMPLE_ARGS)

# Test all packages
.PHONY: test
test:
	go test ./...

# Test with coverage
.PHONY: test-coverage
test-coverage:
	go test -cover ./...

# Test a single package
.PHONY: test-hlcalc
test-hlcalc:
	go test ./pkg/hlcalc

# Format code
.PHONY: fmt
fmt:
	go fmt ./...

# Vet code
.PHONY: vet
vet:
	go vet ./...

# Tidy dependencies
.PHONY: tidy
tidy:
	go mod tidy

# Clean build artifacts
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

# Lint and check (runs fmt, vet, and test)
.PHONY: check
check: fmt vet test

# Install the binary to GOPATH/bin
.PHONY: install
install:
	go install $(LDFLAGS) $(MAIN_PATH)

# Show help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build        - Build binary with git commit hash"
	@echo "  build-dev    - Build binary without git commit hash"
	@echo "  run          - Run the application"
	@echo "  run-example  - Run with example arguments"
	@echo "  test         - Run all tests"
	@echo "  test-coverage- Run tests with coverage"
	@echo "  test-hlcalc  - Run tests for hlcalc package"
	@echo "  fmt          - Format code"
	@echo "  vet          - Vet code"
	@echo "  tidy         - Tidy dependencies"
	@echo "  clean        - Clean build artifacts"
	@echo "  check        - Run fmt, vet, and test"
	@echo "  install      - Install binary to GOPATH/bin"
	@echo "  help         - Show this help"