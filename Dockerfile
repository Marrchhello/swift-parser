FROM golang:1.24-alpine

WORKDIR /app

# Install PostgreSQL client for health checks
RUN apk add --no-cache postgresql-client

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Copy schema and test data
COPY internal/database/schema.sql /app/internal/database/
COPY internal/parser/testdata /app/internal/parser/testdata/

# Build the API server
RUN go build -o main cmd/server/main.go

EXPOSE 8080

# Start the API server
CMD ["./main"]