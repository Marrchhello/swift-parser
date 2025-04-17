package integration

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"swift-parser/internal/models"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	// Setup test environment
	log.Println("Setting up integration test environment...")

	// Wait for API to be ready
	client := &http.Client{Timeout: 5 * time.Second}
	maxRetries := 15
	baseURL := "http://localhost:8080"

	for i := 0; i < maxRetries; i++ {
		_, err := client.Get(baseURL + "/v1/swift-codes/TESTTR00XXX")
		if err == nil {
			log.Println("API is ready")
			break
		}
		log.Printf("Waiting for API (attempt %d/%d)...", i+1, maxRetries)
		time.Sleep(time.Second)
	}

	// Run tests
	code := m.Run()
	os.Exit(code)
}

func TestSWIFTCodeOperations(t *testing.T) {
	baseURL := "http://localhost:8080"
	client := &http.Client{Timeout: 30 * time.Second}

	// Test data cleanup
	cleanup := func() {
		req, _ := http.NewRequest("DELETE", baseURL+"/v1/swift-codes/TESTTR00XXX", nil)
		client.Do(req)
	}
	cleanup()
	defer cleanup()

	tests := []struct {
		name       string
		operation  string
		swiftCode  models.SwiftCode
		wantStatus int
	}{
		{
			name:      "Create SWIFT Code",
			operation: "POST",
			swiftCode: models.SwiftCode{
				SwiftCode:     "TESTTR00XXX",
				CountryISO2:   "TR",
				CountryName:   "Turkey",
				BankName:      "Test Bank",
				Address:       "Test Address",
				IsHeadquarter: true,
			},
			wantStatus: http.StatusCreated,
		},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.swiftCode)
			resp, err := client.Post(
				baseURL+"/v1/swift-codes",
				"application/json",
				bytes.NewBuffer(body),
			)
			if err != nil {
				t.Fatalf("Failed to create SWIFT code: %v", err)
			}
			defer resp.Body.Close()

			respBody, _ := io.ReadAll(resp.Body)
			t.Logf("Response: %s", string(respBody))

			if resp.StatusCode != tt.wantStatus {
				t.Errorf("want status %d, got %d. Response: %s",
					tt.wantStatus, resp.StatusCode, respBody)
			}
		})
	}
}
