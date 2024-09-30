# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Stage 2: Run the Go binary in a lightweight container
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the built Go binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the application runs on
EXPOSE 8080

# Run the Go binary
CMD ["./main"]