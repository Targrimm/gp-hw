package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"

	"github.com/rivo/uniseg"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(original string) (string, error) {
	if len(original) == 0 {
		return "", nil
	}

	result := strings.Builder{}

	// Some hope for performance improvement. Still it's less than needed in most cases.
	result.Grow(len(original))

	result.Len()

	gs := uniseg.NewGraphemes(original)

	var previous *string
	for gs.Next() {
		current := gs.Str()

		number, err := strconv.Atoi(current)

		// Two numbers in a row or number as first grapheme
		if err == nil && previous == nil {
			return "", ErrInvalidString
		}

		if err == nil {
			// Got number: add multiple previous graphemes
			for range number {
				result.WriteString(*previous)
			}
			previous = nil
		} else {
			// Gor common grapheme: add previous if exist
			if previous != nil {
				result.WriteString(*previous)
			}
			previous = &current
		}
	}

	if previous != nil {
		result.WriteString(*previous)
	}

	return result.String(), nil
}
