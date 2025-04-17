# SWIFT Code Parser and API

## Quick Start

1. **Clone the Repository**
```powershell
git clone <repository-url>
cd swift-parser
```

2. **Start the Services**
```powershell
docker compose up -d
```
This will:
- Create PostgreSQL database
- Initialize schema
- Start the API server

3. **Import SWIFT Codes Data**
```powershell
go run cmd/db/init.go
```

4. **Test API Endpoints**

Using PowerShell:
```powershell
# Get SWIFT code details
Invoke-RestMethod -Uri "http://localhost:8080/v1/swift-codes/ANIBAWA1XXX" -Method GET

# Get country codes
Invoke-RestMethod -Uri "http://localhost:8080/v1/swift-codes/country/TR" -Method GET

# Add new SWIFT code
$body = @{
    "swiftCode" = "TESTTR00XXX"
    "countryISO2" = "TR"
    "countryName" = "TURKEY"
    "bankName" = "TEST BANK"
    "address" = "TEST ADDRESS"
    "isHeadquarter" = $true
} | ConvertTo-Json

Invoke-RestMethod -Uri "http://localhost:8080/v1/swift-codes" -Method POST -Body $body -ContentType "application/json"
```

## Requirements
- Docker Desktop
- Go 1.21+ (for local development)
- PostgreSQL client (optional, for direct DB access)

## Stopping the Services
```powershell
docker compose down
```

To remove all data and start fresh:
```powershell
docker compose down -v
```