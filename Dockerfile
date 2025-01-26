# Use the official Golang image as the base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod .

# Download dependencies
RUN go mod download

# Install Vegeta for load testing
RUN go install github.com/tsenart/vegeta@latest

# Copy the source code
COPY . .

# Build the application
RUN go build -o prime-backoff .

# Run the application
CMD ["./prime-backoff"]