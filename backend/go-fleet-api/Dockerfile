# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install git and certificates
RUN apk add --no-cache git ca-certificates

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
# Since we replaced confluent-kafka with segmentio/kafka-go, we don't need CGO enabled
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o fleet-api .

# Final stage
FROM alpine:latest

WORKDIR /app

# Install ca-certificates for HTTPS (e.g. Gemini API)
RUN apk --no-cache add ca-certificates

# Copy the binary from builder
COPY --from=builder /app/fleet-api .

# Expose port
EXPOSE 8080

# Run the executable
CMD ["./fleet-api"]
