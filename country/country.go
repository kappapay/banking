package country

import (
	"github.com/kappapay/banking/bban"
)

// Country holds country related banking info.
type Country struct {
	Name       string
	Alpha2Code string
	Alpha3Code string
	Structure  bban.Structure
}

// String returns text representation of country.
func (c Country) String() string {
	return c.Name
}

// Exists returns true if country code exists.
func Exists(code string) bool {
	_, ok := Countries[code]
	return ok
}

// Get returns country by given country code.
func Get(code string) (Country, bool) {
	country, ok := Countries[code]
	return country, ok
}

// GetBbanStructure returns bban.Structure by given country code.
func GetBbanStructure(code string) (bban.Structure, bool) {
	if country, ok := Get(code); ok {
		return country.Structure, true
	}
	return bban.Structure{}, false
}
