package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"swift-parser/internal/api"
	"swift-parser/internal/database"
	"swift-parser/internal/parser"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("Starting application initialization...")

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("⚠️ Error loading .env file")
	}
	log.Println("✅ Environment variables loaded")

	// Database connection
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := database.NewDB(connStr)
	if err != nil {
		log.Fatalf("⚠️ Database connection failed: %v", err)
	}
	defer db.Close()
	log.Println("✅ Connected to database")

	// Create schema
	schemaPath := filepath.Join("internal", "database", "schema.sql")
	schemaSQL, err := os.ReadFile(schemaPath)
	if err != nil {
		log.Fatalf("⚠️ Failed to read schema file: %v", err)
	}

	_, err = db.Exec(string(schemaSQL))
	if err != nil {
		log.Fatalf("⚠️ Failed to create schema: %v", err)
	}
	log.Println("✅ Database schema created")

	// Parse and insert Excel data
	start := time.Now()
	excelPath := filepath.Join("internal", "parser", "testdata", "Interns_2025_SWIFT_CODES.xlsx")
	codes, err := parser.ParseExcelFile(excelPath)
	if err != nil {
		log.Fatalf("⚠️ Failed to parse Excel file: %v", err)
	}
	log.Printf("✅ Parsed %d SWIFT codes from Excel", len(codes))

	// Debug info for specific SWIFT code
	for _, code := range codes {
		if code.SwiftCode == "AAISALTRXXX" {
			log.Printf("📍 Found SWIFT code: %s", code.SwiftCode)
			log.Printf("   Bank: %s", code.BankName)
			log.Printf("   Country: %s (%s)", code.CountryName, code.CountryISO2)
			log.Printf("   Is Headquarter: %v", code.IsHeadquarter)
		}
	}

	// Insert data into database
	ctx := context.Background()
	if err := db.InsertSwiftCodes(ctx, codes); err != nil {
		log.Fatalf("⚠️ Failed to insert SWIFT codes: %v", err)
	}
	log.Printf("✅ Inserted %d SWIFT codes in %v", len(codes), time.Since(start))

	// Setup and start API server
	router := api.NewRouter(db)
	engine := router.Setup()

	fmt.Println("\n🚀 API Server Ready!")
	fmt.Println("\nAvailable endpoints:")
	fmt.Println("1. GET    http://localhost:8080/v1/swift-codes/{swift-code}")
	fmt.Println("2. GET    http://localhost:8080/v1/swift-codes/country/{countryISO2}")
	fmt.Println("3. POST   http://localhost:8080/v1/swift-codes")
	fmt.Println("4. DELETE http://localhost:8080/v1/swift-codes/{swift-code}")
	fmt.Println("\nExample usage:")
	fmt.Println("- GET http://localhost:8080/v1/swift-codes/BCHICLRMXXX")
	fmt.Println("- GET http://localhost:8080/v1/swift-codes/country/TR")
	fmt.Println("\n📡 Starting server on :8080...")

	if err := engine.Run(":8080"); err != nil {
		log.Fatalf("⚠️ Server failed to start: %v", err)
	}
}
