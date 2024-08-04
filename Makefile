# Makefile for the troubletome Go project

# Project name
PROJECT_NAME := troubletome

# Go commands
GO := go
GO_BUILD := $(GO) build
GO_TEST := $(GO) test

# Directories
SRC_DIR := ./cmd
BUILD_DIR := ./bin

# Executable name
EXECUTABLE := $(BUILD_DIR)/$(PROJECT_NAME)

# Default target
all: build

# Build the Go project
build:
	@echo "Building the project..."
	@mkdir -p $(BUILD_DIR)
	$(GO_BUILD) -o $(EXECUTABLE) $(SRC_DIR)
	@echo "Build completed: $(EXECUTABLE)"

# Run tests
test:
	@echo "Running tests..."
	$(GO_TEST) ./...

# Clean up build files
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean completed."

# Run the application
run: build
	@echo "Running the application..."
	$(EXECUTABLE)

# Install dependencies
deps:
	@echo "Installing dependencies..."
	$(GO) mod tidy

# Format the code
fmt:
	@echo "Formatting the code..."
	$(GO) fmt ./...

