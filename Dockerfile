# Build stage
FROM golang:1.24.2-alpine AS builder

# Install necessary build dependencies
RUN apk add --no-cache git curl ca-certificates tzdata gcc musl-dev

RUN curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sh -s -- -b /usr/local/bin v0.61.1

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the binary with settings for dind compatibility
# CGO_ENABLED=0 ensures a static binary
# -ldflags="-s -w" reduces binary size
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o /app/mcp-docker .

# Final stage using Docker-in-Docker
FROM docker:dind

# Set metadata
LABEL maintainer="c0rwin"

# Set timezone
ENV TZ=UTC

# Copy the binary from the builder stage
COPY --from=builder /app/mcp-docker /usr/local/bin/mcp-docker
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/local/bin/trivy /usr/local/bin/trivy

# Create a non-root user for the application
RUN addgroup -S mcp && adduser -S mcp -G mcp && \
    addgroup mcp docker

# Set necessary permissions
RUN chmod +x /usr/local/bin/mcp-docker

# Environment variables
ENV PATH="/usr/local/bin:${PATH}"
ENV DOCKER_TLS_CERTDIR=""

# Set working directory
WORKDIR /home/mcp

# Expose the MCP server port (default is 8080, adjust as needed)
EXPOSE 8080

# Add healthcheck to verify Docker daemon is running
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
  CMD docker info >/dev/null || exit 1

# Command to run (Docker daemon must run as root)
ENTRYPOINT ["/usr/local/bin/mcp-docker", "serve"]
