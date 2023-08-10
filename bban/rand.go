package bban

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

const (
	alphaUpperCharSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphaNumCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	numCharset        = "0123456789"
)

var (
	randSeeded                  = rand.New(rand.NewSource(time.Now().UnixNano()))
	ErrGeneratedBbanPartInvalid = errors.New("bban: generated bban part is invalid")
)

// stringFromCharset generates a random string of given length from character set
func stringFromCharset(length int, charSet string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charSet[randSeeded.Intn(len(charSet))]
	}
	return string(b)
}

// Rand generates a random valid BBAN of the given structure for testing purposes
func Rand(structure Structure) (string, error) {

	rv := ""

	for _, part := range structure.Parts() {

		c := part.CharType()
		var value string

		switch c {
		case Num:
			value = stringFromCharset(part.Length, numCharset)
		case AlphaUpper:
			value = stringFromCharset(part.Length, alphaUpperCharSet)
		case AlphaNum:
			value = stringFromCharset(part.Length, alphaNumCharSet)
		case Zero:
			value = strings.Repeat("0", part.Length)
		}

		if ok := c.Validate(value); !ok {
			return "", ErrGeneratedBbanPartInvalid
		}

		rv = rv + value
	}
	return rv, nil
}
