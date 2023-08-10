package iban

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/kappapay/banking/bban"
	"github.com/kappapay/banking/country"
)

func TestRandBban(t *testing.T) {

	// Tests generating BBANs for all supported countries

	for _, c := range country.Countries {
		t.Run(c.Name, func(t *testing.T) {

			v, err := bban.Rand(c.Structure)
			require.NoError(t, err)

			err = validateBban(v, c.Structure)
			require.NoError(t, err)

		})
	}
}
