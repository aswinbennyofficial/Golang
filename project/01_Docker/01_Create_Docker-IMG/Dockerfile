# Use the Alpine Linux version 3.9 as the base image with Go version 1.12.0
FROM golang:1.12.0-alpine3.9

# Create a directory named "app" in the container
RUN mkdir /app

# Copy the content of the current directory (the context) into the "/app" directory in the container
ADD . /app

# Set the working directory to "/app"
WORKDIR /app

# Build the Go application and create an executable named "main" in the "/app" directory
RUN go build -o main .

# Specify the default command to run when the container starts, which runs the "main" executable
CMD ["/app/main"]
