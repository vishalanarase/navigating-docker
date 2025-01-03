# Start with the official Golang image as a base
FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o webserver .

# Expose the port the app runs on
EXPOSE 8081

# Command to run the application
CMD ["./webserver"]
