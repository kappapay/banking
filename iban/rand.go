package iban

import (
	"github.com/kappapay/banking/bban"
	"github.com/kappapay/banking/country"
)

// Rand generates a random IBAN for a given country, optionally ending in a specified string
// `isoCountryCode` must an ISO 3166-1 alpha-2 code of a country that implements IBAN
func Rand(isoCountryCode string, endsIn string) (*Iban, error) {

	countryObj, ok := country.Countries[isoCountryCode]
	if !ok {
		return nil, ErrCountryCodeNotPresent
	}

	randBban, err := bban.Rand(countryObj.Structure)
	if err != nil {
		return nil, err
	}

	if endsIn != "" {
		randBban = randBban[:len(randBban)-len(endsIn)] + endsIn
	}

	bbanParts, err := extractBbanParts(randBban, countryObj.Structure)
	if err != nil {
		return nil, err
	}

	return Construct(isoCountryCode, bbanParts...)

}
