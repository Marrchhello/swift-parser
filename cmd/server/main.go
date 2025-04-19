package main

import (
	"fmt"
	"log"
	"os"
	"swift-parser/internal/api"
	"swift-parser/internal/database"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("Starting API server...")

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found, using environment variables")
	}

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

	// Setup and start API server
	router := api.NewRouter(db)
	engine := router.Setup()

	fmt.Println("\n🚀 API Server Ready!")
	fmt.Println("\nAvailable endpoints:")
	fmt.Println("1. GET    http://localhost:8080/v1/swift-codes/{swift-code}")
	fmt.Println("2. GET    http://localhost:8080/v1/swift-codes/country/{countryISO2}")
	fmt.Println("3. POST   http://localhost:8080/v1/swift-codes")
	fmt.Println("4. DELETE http://localhost:8080/v1/swift-codes/{swift-code}")

	log.Printf("📡 Starting server on :8080...")
	if err := engine.Run(":8080"); err != nil {
		log.Fatalf("⚠️ Server failed to start: %v", err)
	}
}
