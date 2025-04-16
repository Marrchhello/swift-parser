package validator

import (
	"regexp"
)

// ValidateSWIFT checks if the given SWIFT code is valid according to specific rules.
func ValidateSWIFT(swiftCode string) bool {
	// SWIFT codes must be 8 or 11 characters long
	if len(swiftCode) != 8 && len(swiftCode) != 11 {
		return false
	}

	// Regular expression to match the SWIFT code format
	swiftRegex := regexp.MustCompile(`^[A-Z]{6}[A-Z0-9]{2}([A-Z0-9]{3})?$`)
	return swiftRegex.MatchString(swiftCode)
}