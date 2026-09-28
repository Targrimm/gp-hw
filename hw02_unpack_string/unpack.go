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
	for range uniseg.GraphemeClusterCount(original) + 1 {
		gs.Next()
		current := gs.Str()

		number, err := strconv.Atoi(current)

		if err == nil {
			if previous == nil {
				return "", ErrInvalidString
			}
			for range number {
				result.WriteString(*previous)
			}
			previous = nil
		} else {
			if previous != nil {
				result.WriteString(*previous)
			}
			previous = &current
		}
	}

	return result.String(), nil
}
