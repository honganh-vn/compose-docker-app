# Build stage
FROM golang:1.20 AS build

WORKDIR /app

# Copy only the Go source file into the container
COPY main.go .

# Initialize the Go module and tidy up dependencies
RUN go mod init simple-web-app && go mod tidy

# Build the Go application
RUN go build -o /backend/main .

# Final stage
FROM alpine:latest

WORKDIR /root/

# Copy the compiled binary and frontend files
COPY --from=build /app/main .
COPY frontend ./frontend

# Expose the port and start the server
EXPOSE 8080
CMD ["./main"]
