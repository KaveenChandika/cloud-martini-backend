# Use the official Golang image as a build stage
FROM golang:1.23.4 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o cloud-martini-backend ./cmd

# Use a minimal base image for the final stage
FROM alpine:latest

# Install necessary dependencies
RUN apk add --no-cache ca-certificates

# Set the working directory inside the container
WORKDIR /root/

# Copy the built Go application from the builder stage
COPY --from=builder /app/cloud-martini-backend .

# Ensure the binary has execute permissions
RUN chmod +x cloud-martini-backend

# Expose the port the application runs on
EXPOSE 8080

# Command to run the application
CMD ["./cloud-martini-backend"]