FROM golang:1.24-alpine
LABEL authors="gustavo"

WORKDIR /app
# Copy go.mod and go.sum files
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
# Copy the source from the current directory to the Working Directory inside the container
COPY . .
# Build the Go app
RUN go build -o /app/cmd/main /app/cmd/main.go
# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["/app/cmd/main"]