package integration

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"swift-parser/internal/models"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	// Wait for API to be ready
	client := &http.Client{Timeout: 5 * time.Second}
	maxRetries := 10

	for i := 0; i < maxRetries; i++ {
		_, err := client.Get("http://localhost:8080/v1/swift-codes/TESTTR00XXX")
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}

	os.Exit(m.Run())
}

func TestAPIEndpoints(t *testing.T) {
	baseURL := "http://localhost:8080"
	client := &http.Client{Timeout: 30 * time.Second}

	// Clean up any existing test data
	cleanup := func() {
		req, _ := http.NewRequest("DELETE", baseURL+"/v1/swift-codes/TESTTR00XXX", nil)
		client.Do(req)
	}
	cleanup()
	defer cleanup()

	t.Run("SWIFT Code Operations", func(t *testing.T) {
		// Test POST new SWIFT code
		newCode := models.SwiftCode{
			SwiftCode:     "TESTTR00XXX",
			CountryISO2:   "TR",
			CountryName:   "Turkey",
			BankName:      "Test Bank",
			Address:       "Test Address",
			IsHeadquarter: true,
		}

		body, _ := json.Marshal(newCode)
		resp, err := client.Post(
			baseURL+"/v1/swift-codes",
			"application/json",
			bytes.NewBuffer(body),
		)
		if err != nil {
			t.Fatalf("Failed to create SWIFT code: %v", err)
		}
		defer resp.Body.Close()

		// Read and log response body for debugging
		respBody, _ := io.ReadAll(resp.Body)
		t.Logf("POST Response: %s", string(respBody))

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("Expected status 201, got %d. Response: %s", resp.StatusCode, respBody)
			return
		}

		// ...rest of the test cases...
	})
}
