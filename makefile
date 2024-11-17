# Makefile for Go project dependencies setup

# Default target: set up project dependencies
setup: tidy deps-upgrade
	@echo "Project dependencies are set up."

build:
	@echo "Building the project..."
	@go build -o bin/ ./cmd/main.go
	@echo "done build"

chmod-runner:
	@chmod +x ./runner
	
# Tidy up the Go modules
tidy:
	@echo "Tidying Go modules..."
	@go mod tidy
	@echo "done"

# Upgrade all dependencies
deps-upgrade:
	@echo "Upgrading dependencies..."
	@go get -u -t -d -v ./...
	@go mod tidy
	@echo "done"

# # Reset dependencies to the state in go.mod
# deps-reset:
# 	@echo "Resetting dependencies..."
# 	git checkout -- go.mod
# 	go mod tidy


# # Clean Go module cache
# deps-cleancache:
# 	@echo "Cleaning module cache..."
# 	@go clean -modcache
# 	@echo "done"

# Declare targets as phony
.PHONY: setup tidy deps-reset deps-upgrade deps-cleancache
