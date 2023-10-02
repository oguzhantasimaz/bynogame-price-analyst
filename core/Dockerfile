# Stage 1: Build the Go application
FROM golang:1.21.1-alpine AS builder

RUN mkdir /app
WORKDIR /app

# Copy only the necessary Go source code files.
# This helps to avoid unnecessary dependencies in the final image.
COPY cmd cmd/
COPY go.mod go.mod
COPY go.sum go.sum

# Run go mod tidy to download and cache the Go modules
RUN go mod tidy

# Build the Go application
RUN go build -o main cmd/main.go

# Stage 2
FROM golang:1.21.1-alpine

RUN mkdir /app
WORKDIR /app

# Copy the built Go binary from the builder stage
COPY --from=builder /app/main /app/main

# Copy the .env file to /app
COPY .env /app/.env

# Set the entry point for the final image
CMD ["/app/main"]
