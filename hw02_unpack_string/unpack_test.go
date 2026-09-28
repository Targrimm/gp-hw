package hw02unpackstring

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input       string
		output      string
		expectError bool
	}{
		{input: "a4bc2d5e", output: "aaaabccddddde", expectError: false},
		{input: "abcd", output: "abcd", expectError: false},
		{input: "3abc", output: "", expectError: true},
		{input: "45", output: "", expectError: true},
		{input: "aaa10b", output: "", expectError: true},
		{input: "aaa0b", output: "aab", expectError: false},
		{input: "", output: "", expectError: false},
		{input: "d\n5abc", output: "d\n\n\n\n\nabc", expectError: false},
		{input: "☺4☻3ツ2🏖3🇧🇷2🇳🇱5", output: "☺☺☺☺☻☻☻ツツ🏖🏖🏖🇧🇷🇧🇷🇳🇱🇳🇱🇳🇱🇳🇱🇳🇱", expectError: false},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			output, err := Unpack(test.input)

			if test.output != output {
				t.Errorf("unpack(%q): expected %q, got %q", test.input, test.output, output)
			}
			if err == nil && test.expectError {
				t.Errorf("unpack(%q): expected no error, got %v", test.input, err)
			}
		})
	}
}
