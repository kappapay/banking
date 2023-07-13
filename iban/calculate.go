package iban

// CalculateCheckDigit returns check digits with temporary iban (iban with 00 check digits)
func CalculateCheckDigit(tempIBAN string) (string, error) {
	code := extractCountryCode(tempIBAN)
	if err := validateCountryCode(code); err != nil {
		return "", err
	}

	calc, err := calculateCheckDigit(tempIBAN, code)
	if err != nil {
		return "", err
	}
	return calc, nil
}
