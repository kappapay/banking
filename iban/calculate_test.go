package iban

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	cases = []struct {
		iban                 string
		checkDigit           string
	}{
		{
			iban:               "AL47212110090000000235698741",
			checkDigit:         "47",
		},
		{
			iban:          "BY13NBRB3600900000002Z00AB00",
			checkDigit:    "13",
		},
		{
			iban:               "BE00539007547034",
			checkDigit:         "68",
		},
		{
			iban:               "CM0010002000300277976315008",
			checkDigit:         "21",
		},
		{
			iban:          "EG800002000156789012345180002",
			checkDigit:    "80",
		},
		{
			iban:               "FO6264600001631634",
			checkDigit:         "62",
		},
		{
			iban:          "GL8964710001000206",
			checkDigit:    "89",
		},
		{
			iban:          "GR1601101250000000012300695",
			checkDigit:    "16",
		},
		{
			iban:          "GB29NWBK60161331926819",
			checkDigit:    "29",
		},
	}
)

func TestExportedCalculateCheckDigit(t *testing.T) {
	for _, cs := range validCases {
		t.Run(cs.iban, func(t *testing.T) {
			checkDigit, err := CalculateCheckDigit(cs.iban)
			require.NoError(t, err)
			require.Equal(t, cs.checkDigit, checkDigit)
		})
	}
}
