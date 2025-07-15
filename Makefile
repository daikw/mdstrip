VERSION = $(shell git describe --tags --always --dirty)
CURRENT_REVISION = $(shell git rev-parse --short HEAD)
BUILD_LDFLAGS = "-s -w -X main.version=$(VERSION) -X main.revision=$(CURRENT_REVISION)"
VERBOSE_FLAG = $(if $(VERBOSE),-v)
EXECUTABLE = mdstrip
ifeq ($(OS),Windows_NT)
    EXECUTABLE := $(EXECUTABLE).exe
endif

.PHONY: help
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: test
test: ## Run tests
	go test $(VERBOSE_FLAG) ./...

.PHONY: lint
lint: ## Run golangci-lint
	golangci-lint run $(VERBOSE_FLAG) --timeout=5m

.PHONY: fmt
fmt: ## Format code
	go fmt ./...

.PHONY: build
build: ## Build the binary
	go build $(VERBOSE_FLAG) -ldflags=$(BUILD_LDFLAGS) -o $(EXECUTABLE) cmd/*

.PHONY: install
install: ## Install the binary
	go install $(VERBOSE_FLAG) -ldflags=$(BUILD_LDFLAGS) ./cmd

.PHONY: clean
clean: ## Clean build artifacts
	rm -f $(EXECUTABLE)
	rm -rf dist/

.PHONY: deps
deps: ## Download dependencies
	go mod download
	go mod tidy

.PHONY: goreleaser
goreleaser: ## Run goreleaser in snapshot mode
	goreleaser release --snapshot --clean

.PHONY: run
run: build ## Build and run
	./$(EXECUTABLE)