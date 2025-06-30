# Use Go 1.24 Alpine image
FROM golang:1.24-alpine

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Go app (adjust path if main.go is elsewhere)
RUN go build -o jobqueue ./cmd/api

# Expose the app port
EXPOSE 8080

# Run the app
CMD ["./jobqueue"]
