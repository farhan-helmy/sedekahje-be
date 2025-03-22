# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application from cmd/server/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o main ./cmd/server

# Final stage
FROM alpine:latest

WORKDIR /app

# Add ca-certificates for HTTPS and timezone data
RUN apk --no-cache add ca-certificates tzdata

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose the port your application runs on
EXPOSE 8080

# Run the application
CMD ["./main"]