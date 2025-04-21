SWIFT Code Parser and API
Prerequisites Installation Guide
1. Install Go

# Download Go for Windows from official website
Start-Process "https://golang.org/dl/"

# Verify installation
go version

2. Install Required Go Libraries

# Initialize module
go mod init swift-parser

# Install dependencies
go get github.com/gin-gonic/gin
go get github.com/lib/pq
go get github.com/joho/godotenv
go get github.com/xuri/excelize/v2

3. Install Docker Desktop

# Download Docker Desktop for Windows
Start-Process "https://www.docker.com/products/docker-desktop"

# Verify installation
docker --version
docker-compose --version

4. Install PostgreSQL Client 
# Using winget
winget install PostgreSQL.PostgreSQL

Project Setup

1. Clone Repository

git clone <repository-url>
cd swift-parser

Running the Application

1. Start Docker Services

# Build and start services
docker compose up -d --build

# Check services status
docker compose ps

Testing

1. Run All Tests

# Run all tests with verbose output
go test ./... -v

API Usage

1. Get SWIFT Code Details

$response = Invoke-RestMethod -Uri "http://localhost:8080/v1/swift-codes/AAISALTRXXX" -Method GET
$response | ConvertTo-Json -Depth 10

2. Get Country SWIFT Codes

$response = Invoke-RestMethod -Uri "http://localhost:8080/v1/swift-codes/country/PL" -Method GET
$response | ConvertTo-Json -Depth 10

3. Create New SWIFT Code

$body = @{
    swiftCode = "TESTTR05XXX"
    countryISO2 = "TR"
    countryName = "TURKEY"
    bankName = "TEST BANK"
    address = "TEST ADDRESS"
    isHeadquarter = $true
} | ConvertTo-Json

$response = Invoke-RestMethod -Uri "http://localhost:8080/v1/swift-codes" `
    -Method POST `
    -Body $body `
    -ContentType "application/json"
$response | ConvertTo-Json

4. Delete SWIFT Code

$response = Invoke-RestMethod -Uri "http://localhost:8080/v1/swift-codes/TESTTR05XXX" -Method DELETE
$response | ConvertTo-Json