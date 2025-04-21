# SWIFT Code Parser and API

A Go-based RESTful API for managing and querying SWIFT codes, backed by PostgreSQL and containerized using Docker.

---

## 📦 Prerequisites

### 1️⃣ Install Go

- Download and install Go from the official website:

```powershell
Start-Process "https://golang.org/dl/"
```

- Verify installation:

```bash
go version
```

---

### 2️⃣ Install Required Go Libraries

- Initialize the module:

```bash
go mod init swift-parser
```

- Install dependencies:

```bash
go get github.com/gin-gonic/gin
go get github.com/lib/pq
go get github.com/joho/godotenv
go get github.com/xuri/excelize/v2
```

---

### 3️⃣ Install Docker Desktop

- Download Docker Desktop for Windows:

```powershell
Start-Process "https://www.docker.com/products/docker-desktop"
```

- Verify installation:

```bash
docker --version
docker-compose --version
```

---

### 4️⃣ Install PostgreSQL Client

- Install using `winget`:

```bash
winget install PostgreSQL.PostgreSQL
```

---

## 📂 Project Setup

- Clone the repository and navigate into the project directory:

```bash
git clone <repository-url>
cd swift-parser
```

---

## 🚀 Running the Application

- Build and start Docker services:

```bash
docker compose up -d --build
```

- Check services status:

```bash
docker compose ps
```

---

## 🧪 Testing

- Run all tests with verbose output:

```bash
go test ./... -v
```

---

## 📡 API Usage

### 📖 1. Get SWIFT Code Details

```powershell
$response = Invoke-RestMethod -Uri "http://localhost:8080/v1/swift-codes/BCECCLRFXXX" -Method GET
$response | ConvertTo-Json -Depth 10
---

```powershell

$response = Invoke-RestMethod -Uri "http://localhost:8080/v1/swift-codes/BCECCLRFXXX" -Method GET
$response | ConvertTo-Json -Depth 10```

---

### 🌍 2. Get Country SWIFT Codes

```powershell
$response = Invoke-RestMethod -Uri "http://localhost:8080/v1/swift-codes/country/PL" -Method GET
$response | ConvertTo-Json -Depth 10
```

---

### 📝 3. Create New SWIFT Code

```powershell
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
```

---

### ❌ 4. Delete SWIFT Code

```powershell
$response = Invoke-RestMethod -Uri "http://localhost:8080/v1/swift-codes/TESTTR05XXX" -Method DELETE
$response | ConvertTo-Json
```



---

## 📧 Contact

For questions, reach out to markijanvoloshyn@gmail.com.
