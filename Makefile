# Variables for Go Commands
GO := go
BIN_DIR := bin

# Build the Go Project
build:
	@echo "Building the project..."
	@if not exist $(BIN_DIR) mkdir $(BIN_DIR)
	@$(GO) build -o $(BIN_DIR)/npmdc cmd/main.go
	@echo "Building Successful"

# Run the Go Project
run:
	@echo "Running the project..."
	@cls
	@$(GO) run cmd/main.go
	@echo "Run completed."

# Build and Run the Go Project
buildandrun: build
	@$(BIN_DIR)/npmdc

# Clean the Project (Remove bin directory)
clean:
	@echo "Cleaning the project..."
	@if exist $(BIN_DIR) rmdir /s /q $(BIN_DIR)
	@echo "Cleaning Successful"

# Install Go dependencies (like nimble install)
install:
	@echo "Installing dependencies..."
	@$(GO) get github.com/cligen/cligen
	@$(GO) get github.com/colorize/colorize
	@$(GO) get github.com/spinny/spinny
	@$(GO) get github.com/illwill/illwill
	@$(GO) get github.com/dashing/dashing
	@echo "Dependencies installed successfully."
