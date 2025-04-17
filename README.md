# MCP Docker Integration Tool

## Overview

MCP Docker is a powerful integration tool that provides Docker operations through the Model Control Protocol (MCP) platform. It serves as a bridge between Docker commands and API services, allowing seamless interaction with Docker containers and images through a standardized API interface.

### Available Operations

- `ps` - List and manage Docker containers
- `exec` - Execute commands in containers
- `diff` - Inspect changes to container filesystems
- `history` - Show image history
- `image` - Manage Docker images
- `inspect` - Get detailed information about Docker objects
- `pull` - Pull images from registries
- `run` - Create and start containers
- `sbom` - Generate Software Bill of Materials
- `search` - Search Docker images

## Prerequisites

- Go 1.24.2 or higher
- Docker installed and running
- golangci-lint (for development)

## Installation

### Build from Source

```bash
# Clone the repository
git clone https://github.com/mark3labs/mcp-docker
cd mcp-docker

# Build for current platform
make build

# Build for all supported platforms
make build-all
```

### Docker Image

```bash
# Build Docker image
make docker

# Build and push to registry
REGISTRY=your-registry make docker-push
```

## Makefile Commands

| Command | Description | Example |
|---------|-------------|---------|
| `make build` | Build binary for current platform | `make build` |
| `make build-all` | Build for multiple platforms (linux/amd64, linux/arm64, darwin/amd64, darwin/arm64) | `make build-all` |
| `make docker` | Build Docker image | `make docker` |
| `make docker-push` | Push Docker image to registry | `REGISTRY=your-registry make docker-push` |
| `make test` | Run tests | `make test` |
| `make test-coverage` | Run tests with coverage | `make test-coverage` |
| `make lint` | Run linter | `make lint` |
| `make clean` | Remove build artifacts | `make clean` |
| `make clean-docker` | Remove Docker images | `make clean-docker` |
| `make clean-all` | Remove all artifacts and Docker images | `make clean-all` |

### Customizable Variables

- `VERSION` (default: 0.1.0)
- `REGISTRY` (default: empty)
- `GOOS` (default: linux)
- `GOARCH` (default: amd64)

## Usage

### Starting the Service

```bash
# Run directly
./mcp-docker serve

# Run via Docker
docker run -d \
  --name mcp-docker \
  --privileged \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -e DOCKER_TLS_CERTDIR= \
  -p 8080:8080 \
  mcp-docker:latest
```

## Development

### Project Structure

```
.
├── api/          # API definitions
├── cmd/          # Command line interface
├── internal/     # Internal packages
│   └── docker/   # Docker operations implementation
├── pkg/          # Public packages
├── Dockerfile    # Container image definition
├── Makefile     # Build and development commands
└── main.go      # Application entry point
```

### Testing

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage
```

### Linting

```bash
# Run linter
make lint
```

## Docker Configuration

### Environment Variables

- `DOCKER_TLS_CERTDIR`: TLS certificate directory
- `DOCKER_API_VERSION`: Docker API version (default: 1.48)
- `TZ`: Timezone (default: UTC)

### Volume Mounts

- `/var/run/docker.sock:/var/run/docker.sock` - Docker daemon socket

### Network Configuration

The container requires access to:
- Host network (when using host networking mode)
- Docker daemon socket
- Exposed port 8080 for API access

## Claude.AI Integration

```json
{
  "docker": {
    "command": "docker",
    "args": [
      "run",
      "--rm",
      "-i",
      "--name",
      "mcp-docker",
      "--privileged",
      "-v",
      "/var/run/docker.sock:/var/run/docker.sock",
      "-e",
      "DOCKER_TLS_CERTDIR=",
      "-e",
      "DOCKER_API_VERSION=1.48",
      "--network",
      "host",
      "mcp-docker:latest"
    ]
  }
}
```

## Dependencies

### Main Dependencies

- github.com/mark3labs/mcp-go v0.20.1
- github.com/spf13/cobra v1.9.1

### Indirect Dependencies

- github.com/google/uuid v1.6.0
- github.com/inconshreveable/mousetrap v1.1.0
- github.com/spf13/pflag v1.0.6
- github.com/yosida95/uritemplate/v3 v3.0.2

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

The MIT License is a permissive license that is short and to the point. It lets people do anything they want with your code as long as they provide attribution back to you and don't hold you liable.
## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Security

This tool requires privileged access to the Docker daemon. Ensure proper security measures are in place when deploying in production environments.

