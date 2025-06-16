# Use the official Golang image as the base image
FROM golang:1.24

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application
RUN go build -o myapp main.go

# Expose port 8080 (the port our Go web server listens on)
EXPOSE 8080

# Command to run when the container starts
CMD ["./myapp"]