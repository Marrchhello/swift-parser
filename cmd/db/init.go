package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"swift-parser/internal/database"
	"swift-parser/internal/parser"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("Starting database initialization...")

	if err := godotenv.Load(); err != nil {
		log.Fatal("⚠️ Error loading .env file")
	}
	log.Println("✅ Environment variables loaded")

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

	// Insert data into database
	ctx := context.Background()
	if err := db.InsertSwiftCodes(ctx, codes); err != nil {
		log.Fatalf("⚠️ Failed to insert SWIFT codes: %v", err)
	}
	log.Printf("✅ Inserted %d SWIFT codes in %v", len(codes), time.Since(start))

	log.Println("🎉 Database initialization complete!")
}
