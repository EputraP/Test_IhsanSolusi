package validator

import (
	"regexp"
	"strings"

	errs "github.com/EputraP/Test_IhsanSolusi/internal/errors"
)

func ValidateName(name string) error {
	// Trim spaces
	name = strings.TrimSpace(name)

	// Check length
	if len(name) < 2 || len(name) > 50 {
		return errs.InvalidNameLength
	}

	// Define regex pattern: only letters and spaces, no numbers or special characters
	pattern := `^[A-Za-z][A-Za-z\s]*[A-Za-z]$`

	// Compile and match regex
	matched, _ := regexp.MatchString(pattern, name)
	if !matched {
		return errs.InvalidNameFormat
	}

	return nil
}

// ValidatePhoneNumber checks if a phone number is valid
func ValidatePhoneNumber(phone string) error {
	// Trim spaces
	phone = strings.TrimSpace(phone)

	// Define regex pattern: Allows numbers, optional leading `+`, and length 10-15
	pattern := `^\+?[0-9]{10,15}$`

	matched, _ := regexp.MatchString(pattern, phone)
	if !matched {
		return errs.InvalidPhoneNumberFormat
	}

	return nil
}

func ValidateNIK(nik string) error {
	// Trim spaces
	nik = strings.TrimSpace(nik)

	// Define regex pattern: exactly 16 digits
	pattern := `^[0-9]{16}$`

	// Check if NIK matches the pattern
	matched, _ := regexp.MatchString(pattern, nik)
	if !matched {
		return errs.InvalidNIKFormat
	}

	return nil
}
