package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"swift-parser/internal/database"
	"swift-parser/internal/models"
	"testing"

	"github.com/gin-gonic/gin"
)

// TestResponse matches the API response structure
type TestResponse struct {
	Address       string             `json:"address"`
	BankName      string             `json:"bankName"`
	CountryISO2   string             `json:"countryISO2"`
	CountryName   string             `json:"countryName"`
	IsHeadquarter bool               `json:"isHeadquarter"`
	SwiftCode     string             `json:"swiftCode"`
	Branches      []models.SwiftCode `json:"branches,omitempty"`
}

func TestGetSWIFTCode(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, err := database.NewDB("host=localhost port=5432 user=postgres password=3107 dbname=postgres sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	router := NewRouter(db)
	engine := router.Setup()

	tests := []struct {
		name       string
		swiftCode  string
		wantStatus int
		wantHQ     bool
	}{
		{
			name:       "Get Headquarter SWIFT",
			swiftCode:  "ANIBAWA1XXX",
			wantStatus: http.StatusOK,
			wantHQ:     true,
		},
		{
			name:       "Get Branch SWIFT",
			swiftCode:  "BCHICLR10R2",
			wantStatus: http.StatusOK,
			wantHQ:     false,
		},
		{
			name:       "Invalid SWIFT",
			swiftCode:  "INVALID123",
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/v1/swift-codes/"+tt.swiftCode, nil)
			engine.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("want status %d, got %d", tt.wantStatus, w.Code)
			}

			// Only check response body for successful requests
			if w.Code == http.StatusOK {
				var response TestResponse
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}

				// Validate response fields
				if response.SwiftCode != tt.swiftCode {
					t.Errorf("want SWIFT code %s, got %s", tt.swiftCode, response.SwiftCode)
				}

				if response.IsHeadquarter != tt.wantHQ {
					t.Errorf("want IsHeadquarter %v, got %v", tt.wantHQ, response.IsHeadquarter)
				}

				// Check branches array for headquarters
				if tt.wantHQ && response.Branches == nil {
					t.Error("want branches array for headquarter, got nil")
				}
			}
		})
	}
}
