package validator

import (
	"regexp"
)

func ValidateSWIFT(swiftCode string) bool {
	if len(swiftCode) != 8 && len(swiftCode) != 11 {
		return false
	}

	swiftRegex := regexp.MustCompile(`^[A-Z]{6}[A-Z0-9]{2}([A-Z0-9]{3})?$`)
	return swiftRegex.MatchString(swiftCode)
}
