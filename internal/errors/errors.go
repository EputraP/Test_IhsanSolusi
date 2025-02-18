package errs

import "errors"

var (
	InvalidRequestBody = errors.New("invalid request body")

	ErrorCreatingUser    = errors.New("error creating user")
	ErrorCheckingUser    = errors.New("error checking user")
	ErrorNoHPAlreadyUsed = errors.New("phone number already in use")
	ErrorNIKAlreadyUsed  = errors.New("NIK already in use")

	ErrorCreatingUserSaldo = errors.New("error creating user saldo")
	ErrorGettingUserSaldo  = errors.New("error retrieving user saldo")

	ErrorStringIntConversion = errors.New("error converting string to int")

	ErrorNoRek12DigitsAndOnlyNumber            = errors.New("invalid account number: must be exactly 12 digits and contain only numbers")
	ErrorNominalNumberWithoutThousandSeparator = errors.New("invalid nominal input: must contain only numbers without thousand separators (e.g., 1000, 25000)")
	ErrorNominalMustValidNumericValue          = errors.New("invalid nominal input: must be a valid numeric value")
	InsufficientBalance                        = errors.New("insufficient balance")

	InvalidNoRek             = errors.New("invalid account number")
	InvalidNameLength        = errors.New("name must be between 2 and 50 characters long")
	InvalidNameFormat        = errors.New("name can only contain letters and spaces and must start and end with a letter")
	InvalidPhoneNumberFormat = errors.New("invalid phone number format: must contain only numbers and may start with '+'")
	InvalidNIKFormat         = errors.New("invalid NIK: must be exactly 16 digits and contain only numbers")
)
