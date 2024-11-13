# Dockerfile
FROM golang:1.21

# Set working directory
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Set PATH for Air binary
ENV PATH="/go/bin:${PATH}"

# Default command to run Air
CMD ["air"]