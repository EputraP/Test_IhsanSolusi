package errs

import "errors"

var (
	InvalidRequestBody = errors.New("invalid request body")

	ErrorCreatingUser    = errors.New("Error on creating user")
	ErrorCheckingUser    = errors.New("Error on checking user")
	ErrorNoHPAlreadyUsed = errors.New("No Hp already used")
	ErrorNIKAlreadyUsed  = errors.New("NIK already used")

	ErrorCreatingUserSaldo = errors.New("Error on creating user saldo")
	ErrorGettingUserSaldo  = errors.New("Error on getting user saldo")

	ErrorStringIntConvertion = errors.New("Error on converting string to int")

	ErrorNoRek12DigitsAndOnlyNumber            = errors.New("invalid no_rekening input: must be exactly 12 digits and contain only numbers")
	ErrorNominalNumberWithoutThousandSeparator = errors.New("invalid nominal input: must contain only numbers without thousand separators (e.g., 1000, 25000)")
	ErrorNominalMustValidNumericValue          = errors.New("invalid nominal input: must be a valid numeric value")
	InsufficientBalance                        = errors.New("insufficient balance")

	InvalidNoRek             = errors.New("invalid no rekening")
	InvalidNameLength        = errors.New("name must be between 2 and 50 characters long")
	InvalidNameFormat        = errors.New("name can only contain letters and spaces, and must start and end with a letter")
	InvalidPhoneNumberFormat = errors.New("invalid phone number format: must contain only numbers and may start with '+'")
	InvalidNIKFormat         = errors.New("invalid NIK: must be exactly 16 digits and contain only numbers")
)
