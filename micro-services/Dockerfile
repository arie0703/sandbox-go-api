# Build stage
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Install git (required for some Go modules)
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o microservice-cli .

# Runtime stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Create non-root user for security
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

WORKDIR /root/

# Copy the binary from builder stage with executable permissions
COPY --from=builder /app/microservice-cli .

# Set the binary as executable and change ownership to non-root user
RUN chmod +x microservice-cli && \
    chown appuser:appgroup microservice-cli

# Switch to non-root user
USER appuser

# Set entrypoint to handle command line arguments properly
ENTRYPOINT ["./microservice-cli"]

# Default command (can be overridden)
CMD ["--help"]