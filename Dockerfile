# Use the official golang image from Docker Hub
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Generate mod file
RUN go mod init main

# Build the Go app
RUN go build -o main .

# # Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
