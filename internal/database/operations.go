package database

import (
	"context"
	"database/sql"
	"errors"
	"swift-parser/internal/models"
)

func (db *DB) InsertSwiftCodes(ctx context.Context, codes []models.SwiftCode) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `
        INSERT INTO swift_codes (
            swift_code, country_iso2, country_name, 
            bank_name, address, is_headquarter
        ) VALUES ($1, $2, $3, $4, $5, $6)
        ON CONFLICT (swift_code) DO UPDATE SET
            country_iso2 = $2,
            country_name = $3,
            bank_name = $4,
            address = $5,
            is_headquarter = $6
    `)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, code := range codes {
		_, err = stmt.ExecContext(ctx,
			code.SwiftCode,
			code.CountryISO2,
			code.CountryName,
			code.BankName,
			code.Address,
			code.IsHeadquarter,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (db *DB) GetSWIFTCode(code string) (*models.SwiftCode, error) {
	query := `
        SELECT swift_code, country_iso2, country_name, 
               bank_name, address, is_headquarter
        FROM swift_codes 
        WHERE swift_code = $1`

	var swiftCode models.SwiftCode
	err := db.QueryRow(query, code).Scan(
		&swiftCode.SwiftCode,
		&swiftCode.CountryISO2,
		&swiftCode.CountryName,
		&swiftCode.BankName,
		&swiftCode.Address,
		&swiftCode.IsHeadquarter,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("swift code not found")
	}
	if err != nil {
		return nil, err
	}
	return &swiftCode, nil
}

// GetBranches retrieves all branches for a headquarter SWIFT code
func (db *DB) GetBranches(headquarterCode string) ([]models.SwiftCode, error) {
	// Get base code (first 6 characters) and add wildcard
	baseCode := headquarterCode[:6] + "%"

	query := `
        SELECT swift_code, country_iso2, bank_name, 
               address, is_headquarter
        FROM swift_codes 
        WHERE swift_code LIKE $1 
        AND swift_code != $2 
        AND NOT is_headquarter`

	rows, err := db.Query(query, baseCode, headquarterCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var branches []models.SwiftCode
	for rows.Next() {
		var branch models.SwiftCode
		err := rows.Scan(
			&branch.SwiftCode,
			&branch.CountryISO2,
			&branch.BankName,
			&branch.Address,
			&branch.IsHeadquarter,
		)
		if err != nil {
			return nil, err
		}
		branches = append(branches, branch)
	}
	return branches, nil
}

// GetSWIFTCodesByCountry retrieves all SWIFT codes for a specific country
func (db *DB) GetSWIFTCodesByCountry(countryISO2 string) ([]models.SwiftCode, error) {
	query := `
        SELECT swift_code, country_iso2, country_name,
               bank_name, address, is_headquarter
        FROM swift_codes 
        WHERE country_iso2 = $1
        ORDER BY is_headquarter DESC, swift_code`

	rows, err := db.Query(query, countryISO2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var codes []models.SwiftCode
	for rows.Next() {
		var code models.SwiftCode
		err := rows.Scan(
			&code.SwiftCode,
			&code.CountryISO2,
			&code.CountryName,
			&code.BankName,
			&code.Address,
			&code.IsHeadquarter,
		)
		if err != nil {
			return nil, err
		}
		codes = append(codes, code)
	}

	if len(codes) == 0 {
		return nil, errors.New("no swift codes found for this country")
	}
	return codes, nil
}

// AddSWIFTCode adds a new SWIFT code to the database
func (db *DB) AddSWIFTCode(code *models.SwiftCode) error {
	query := `
        INSERT INTO swift_codes (
            swift_code, country_iso2, country_name,
            bank_name, address, is_headquarter
        ) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := db.Exec(query,
		code.SwiftCode,
		code.CountryISO2,
		code.CountryName,
		code.BankName,
		code.Address,
		code.IsHeadquarter,
	)
	return err
}

// DeleteSWIFTCode deletes a SWIFT code from the database
func (db *DB) DeleteSWIFTCode(code string) error {
	query := `DELETE FROM swift_codes WHERE swift_code = $1`

	result, err := db.Exec(query, code)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("swift code not found")
	}
	return nil
}
