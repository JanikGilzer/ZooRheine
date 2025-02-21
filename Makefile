APP_NAME=go_build_ZooDaBa
DOCKER_IMAGE=zoorheine:latest

# Default target that runs the full pipeline
all: docker-down clean build docker-build docker-up

# Stop and remove containers
docker-down:
	docker compose down

# Clean up binary and Docker images
clean:
	rm -rf build/$(APP_NAME)
	docker rmi $(DOCKER_IMAGE) || true

# Build the Go binary
build:
	mkdir -p build
	go build -o build/$(APP_NAME) main.go

# Build the Docker image
docker-build: build
	docker build -t $(DOCKER_IMAGE) .

# Run Docker Compose
docker-up: docker-build
	docker compose up -d

# Restart the server
restart: docker-down docker-up

.PHONY: all build docker-build docker-up docker-down clean restart
