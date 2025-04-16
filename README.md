# swift-parser

This project is a Go application for parsing SWIFT codes and exposing a REST API.

## Project Structure

```
swift-parser
├── cmd
│   └── server
│       └── main.go
├── internal
│   ├── api
│   │   ├── handlers.go
│   │   └── routes.go
│   ├── parser
│   │   └── swift.go
│   └── models
│       └── swift.go
├── pkg
│   └── validator
│       └── swift.go
├── go.mod
├── go.sum
└── README.md
```

## Setup Instructions

1. Clone the repository:
   ```
   git clone <repository-url>
   cd swift-parser
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Run the application:
   ```
   go run cmd/server/main.go
   ```

## Usage

- **GET /swift**: Retrieve information about a SWIFT code.
- **POST /swift**: Submit a SWIFT code for parsing and validation.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.