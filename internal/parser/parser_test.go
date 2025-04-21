package parser

import (
	"path/filepath"
	"strings"
	"swift-parser/internal/models"
	"testing"
)

func TestParseExcelFile(t *testing.T) {

	tests := []struct {
		name     string
		filePath string
		wantErr  bool
		validate func(t *testing.T, codes []models.SwiftCode)
	}{
		{
			name:     "Valid Excel file",
			filePath: filepath.Join("testdata", "Interns_2025_SWIFT_CODES.xlsx"),
			wantErr:  false,
			validate: func(t *testing.T, codes []models.SwiftCode) {
				if len(codes) == 0 {
					t.Error("Expected non-empty result")
				}

				firstCode := codes[0]
				if firstCode.CountryISO2 == "" {
					t.Error("CountryISO2 should not be empty")
				}
				if firstCode.SwiftCode == "" {
					t.Error("SwiftCode should not be empty")
				}

				if firstCode.CountryISO2 != strings.ToUpper(firstCode.CountryISO2) {
					t.Error("CountryISO2 should be uppercase")
				}

				if strings.HasSuffix(firstCode.SwiftCode, "XXX") != firstCode.IsHeadquarter {
					t.Error("IsHeadquarter flag doesn't match XXX suffix")
				}
			},
		},
		{
			name:     "Non-existent file",
			filePath: "nonexistent.xlsx",
			wantErr:  true,
			validate: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			codes, err := ParseExcelFile(tt.filePath)

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseExcelFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && tt.validate != nil {
				tt.validate(t, codes)
			}
		})
	}
}
