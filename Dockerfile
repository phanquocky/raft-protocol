# Use an official Golang image as the base image
FROM golang:1.23.3 as builder

# Set environment variables for Go
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum for dependency installation
COPY go.mod go.sum ./

# Download and cache dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the application
RUN go run main.go

# Expose the port your gRPC server listens on
EXPOSE 8080
