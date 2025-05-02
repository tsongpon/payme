# Stage 1: Build the Go application
FROM golang:1.22-alpine AS builder

# Set the working directory inside the builder container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the builder container
COPY . .

# Build the Go application
RUN go build -o payme .

# Stage 2: Create a minimal runtime image
FROM alpine:latest

# Set the working directory inside the runtime container
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/payme .

# Expose the application port
EXPOSE 8080

# Command to run the executable
CMD ["./payme"]