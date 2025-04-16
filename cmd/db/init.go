package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"swift-parser/internal/database"
	"swift-parser/internal/parser"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Use environment variables for connection
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
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create schema
	schemaPath := filepath.Join("internal", "database", "schema.sql")
	schemaSQL, err := os.ReadFile(schemaPath)
	if err != nil {
		log.Fatalf("Failed to read schema file: %v", err)
	}

	_, err = db.Exec(string(schemaSQL))
	if err != nil {
		log.Fatalf("Failed to create schema: %v", err)
	}

	// Parse Excel file
	excelPath := filepath.Join("internal", "parser", "testdata", "Interns_2025_SWIFT_CODES.xlsx")
	codes, err := parser.ParseExcelFile(excelPath)
	if err != nil {
		log.Fatalf("Failed to parse Excel file: %v", err)
	}

	log.Printf("Successfully parsed %d SWIFT codes", len(codes))
	log.Println("Successfully initialized database")
}
