# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod files first for better caching
# Use go.sum if exists, otherwise it will be generated
COPY go.mod go.sum* ./

# Download dependencies (this layer will be cached if go.mod/go.sum don't change)
RUN go mod download

# Copy source code
COPY . .

# Create go.sum by getting all dependencies
# This creates go.sum using cached dependencies (offline after go mod download)
RUN go get ./... || true

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o qa-api ./cmd/server

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates iputils

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/qa-api .
COPY --from=builder /app/migrations ./migrations

# Copy entrypoint script
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Expose port
EXPOSE 8080

# Use entrypoint script
ENTRYPOINT ["/entrypoint.sh"]
CMD ["./qa-api"]

