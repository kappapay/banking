package iban

import (
	"github.com/kappapay/banking/country"
)

// Construct constructs an IBAN from BBAN parts for a given country
// bbanParts must be provided in the order specified by the country's BBAN structure
func Construct(alpha2Code string, bbanParts ...string) (*Iban, error) {

	countryObject, ok := country.Countries[alpha2Code]
	if !ok {
		return nil, ErrCountryCodeNotPresent
	}

	if len(bbanParts) != len(countryObject.Structure.Parts()) {
		return nil, ErrInvalidNumberOfBbanParts
	}

	// construct a temporary IBAN with the check digit set to 00

	tempIban := alpha2Code + "00"

	for i, part := range countryObject.Structure.Parts() {
		if !part.Validate(bbanParts[i]) {
			return nil, ErrInvalidBbanPart
		}
		tempIban += bbanParts[i]
	}

	// calculate the check digit for the temporary IBAN

	checkDigit, err := CalculateCheckDigit(tempIban)
	if err != nil {
		return nil, err
	}

	ibanValue := tempIban[:2] + checkDigit + tempIban[4:]

	// check that the IBAN we've constructed is valid

	result, err := Parse(ibanValue)
	if err != nil {
		return nil, ErrConstructedIbanInvalid
	}

	return result, nil

}
