# Makefile for mcp-docker project

# Variables
VERSION ?= 0.1.0
IMAGE := mcp-docker
BINARY := mcp-docker
REGISTRY ?= 
TAG := $(VERSION)
LDFLAGS := "-s -w"
CGO_ENABLED := 0
GOOS ?= linux
GOARCH ?= amd64

# If registry is specified, prepend it to the image name
ifneq ($(REGISTRY),)
  FULL_IMAGE := $(REGISTRY)/$(IMAGE)
else
  FULL_IMAGE := $(IMAGE)
endif

# Default target
.PHONY: all
all: build docker

# Build the binary
.PHONY: build
build:
	@echo "Building $(BINARY)..."
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags=$(LDFLAGS) -o $(BINARY) .

# Build multi-platform binaries
.PHONY: build-all
build-all:
	@echo "Building for multiple platforms..."
	CGO_ENABLED=$(CGO_ENABLED) GOOS=linux GOARCH=amd64 go build -ldflags=$(LDFLAGS) -o $(BINARY)-linux-amd64 .
	CGO_ENABLED=$(CGO_ENABLED) GOOS=linux GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o $(BINARY)-linux-arm64 .
	CGO_ENABLED=$(CGO_ENABLED) GOOS=darwin GOARCH=amd64 go build -ldflags=$(LDFLAGS) -o $(BINARY)-darwin-amd64 .
	CGO_ENABLED=$(CGO_ENABLED) GOOS=darwin GOARCH=arm64 go build -ldflags=$(LDFLAGS) -o $(BINARY)-darwin-arm64 .

# Build Docker image
.PHONY: docker
docker:
	@echo "Building Docker image $(FULL_IMAGE):$(TAG)..."
	docker build -t $(FULL_IMAGE):$(TAG) .
	docker tag $(FULL_IMAGE):$(TAG) $(FULL_IMAGE):latest

# Push Docker image
.PHONY: docker-push
docker-push:
	@echo "Pushing Docker image $(FULL_IMAGE):$(TAG)..."
	docker push $(FULL_IMAGE):$(TAG)
	docker push $(FULL_IMAGE):latest

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	go test -v ./...

# Run test with coverage
.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -cover ./...

# Run linting
.PHONY: lint
lint:
	@echo "Linting..."
	@if ! command -v golangci-lint &> /dev/null; then \
		echo "golangci-lint not found, installing..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	golangci-lint run

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning..."
	rm -f $(BINARY) $(BINARY)-*

# Clean Docker images
.PHONY: clean-docker
clean-docker:
	@echo "Removing Docker images for $(FULL_IMAGE)..."
	-docker rmi $(FULL_IMAGE):$(TAG) 2>/dev/null || true
	-docker rmi $(FULL_IMAGE):latest 2>/dev/null || true

# Clean all (binary and Docker)
.PHONY: clean-all
clean-all: clean clean-docker

# Help target
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  all             : Build binary and Docker image (default)"
	@echo "  build           : Build binary for current platform"
	@echo "  build-all       : Build binary for multiple platforms"
	@echo "  docker          : Build Docker image"
	@echo "  docker-push     : Push Docker image to registry"
	@echo "  test            : Run tests"
	@echo "  test-coverage   : Run tests with coverage"
	@echo "  lint            : Run linter"
	@echo "  clean           : Remove build artifacts"
	@echo "  clean-docker    : Remove Docker images"
	@echo "  clean-all       : Remove all artifacts and Docker images"
	@echo "  help            : Show this help message"
	@echo ""
	@echo "Variables:"
	@echo "  VERSION         : $(VERSION)"
	@echo "  IMAGE           : $(IMAGE)"
	@echo "  REGISTRY        : $(REGISTRY)"
	@echo "  GOOS            : $(GOOS)"
	@echo "  GOARCH          : $(GOARCH)"

