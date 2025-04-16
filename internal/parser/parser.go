package parser

import (
	"fmt"
	"strings"

	"github.com/xuri/excelize/v2"

	"swift-parser/internal/models"
)

func ParseExcelFile(filePath string) ([]models.SwiftCode, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open excel file: %w", err)
	}
	defer f.Close()

	// Get the first sheet
	sheetName := f.GetSheetList()[0]
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to get rows: %w", err)
	}

	var swiftCodes []models.SwiftCode
	// Skip header row
	for i := 1; i < len(rows); i++ {
		row := rows[i]
		if len(row) < 7 {
			continue // Skip invalid rows
		}

		swiftCode := models.SwiftCode{
			Address:       row[4],
			BankName:      row[3],
			CountryISO2:   strings.ToUpper(row[0]),
			CountryName:   strings.ToUpper(row[6]),
			IsHeadquarter: strings.HasSuffix(row[1], "XXX"),
			SwiftCode:     row[1],
		}
		swiftCodes = append(swiftCodes, swiftCode)
	}

	return swiftCodes, nil
}
