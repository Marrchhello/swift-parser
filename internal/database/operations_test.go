package database

import (
	"context"
	"swift-parser/internal/models"
	"testing"
)

func setupTestDB(t *testing.T) *DB {
	db, err := NewDB("host=localhost port=5432 user=postgres password=3107 dbname=postgres sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func TestInsertAndGetSWIFTCode(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	testCode := models.SwiftCode{
		SwiftCode:     "TESTTR00XXX",
		CountryISO2:   "TR",
		CountryName:   "Turkey",
		BankName:      "Test Bank",
		Address:       "Test Address",
		IsHeadquarter: true,
	}

	// Test Insert
	err := db.InsertSwiftCodes(context.Background(), []models.SwiftCode{testCode})
	if err != nil {
		t.Fatalf("Failed to insert SWIFT code: %v", err)
	}

	// Test Get
	got, err := db.GetSWIFTCode(testCode.SwiftCode)
	if err != nil {
		t.Fatalf("Failed to get SWIFT code: %v", err)
	}

	if got.SwiftCode != testCode.SwiftCode {
		t.Errorf("want SwiftCode %s, got %s", testCode.SwiftCode, got.SwiftCode)
	}
}

func TestGetBranches(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	// Insert test data
	hq := models.SwiftCode{
		SwiftCode:     "TESTTR00XXX",
		CountryISO2:   "TR",
		CountryName:   "Turkey",
		BankName:      "Test Bank HQ",
		Address:       "HQ Address",
		IsHeadquarter: true,
	}

	branch := models.SwiftCode{
		SwiftCode:     "TESTTR00001",
		CountryISO2:   "TR",
		CountryName:   "Turkey",
		BankName:      "Test Bank Branch",
		Address:       "Branch Address",
		IsHeadquarter: false,
	}

	ctx := context.Background()
	err := db.InsertSwiftCodes(ctx, []models.SwiftCode{hq, branch})
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	// Test GetBranches
	branches, err := db.GetBranches(hq.SwiftCode)
	if err != nil {
		t.Fatalf("Failed to get branches: %v", err)
	}

	if len(branches) != 1 {
		t.Errorf("want 1 branch, got %d", len(branches))
	}

	if branches[0].SwiftCode != branch.SwiftCode {
		t.Errorf("want branch code %s, got %s", branch.SwiftCode, branches[0].SwiftCode)
	}
}
