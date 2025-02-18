package validator

import (
	"regexp"
	"strconv"

	errs "github.com/EputraP/Test_IhsanSolusi/internal/errors"
)

func Validate12DigitNumber(input string) error {
	// Regular expression to match exactly 12 digits
	re := regexp.MustCompile(`^\d{12}$`)

	if !re.MatchString(input) {
		return errs.Error12DigitsAndOnlyNumber
	}

	return nil
}

func ValidateRupiahNominal(input string) error {
	// Regex pattern to allow only digits (no dots, no commas, no letters)
	re := regexp.MustCompile(`^\d+$`)

	if !re.MatchString(input) {
		return errs.NumberWithoutThousandSeparator
	}

	// Convert to integer to ensure it's a valid number
	if _, err := strconv.Atoi(input); err != nil {
		return errs.MustValidNumericValue
	}

	return nil
}
