package arrays_and_strings

import (
	"testing"
	"unicode"
)

func TestPalindromePermutation(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output bool
	}{
		{"Should be true", "Tact Coa", true},
		{"Should be true", "tact cao", true},
		{"Should be false", "tact  xcao", false},
		{"Should be true", "", true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := PalindromePermutation(test.input)
			if result != test.output {
				t.Logf("For %v the answer should be %v not %v", test.input, test.output, result)
				t.Fail()
			}
		})
	}
}

func PalindromePermutation(input string) bool {
	box := make(map[rune]bool)
	for _, v := range input {
		if v != ' ' {
			v = unicode.ToLower(v)
			if box[v] {
				delete(box, v)
			} else {
				box[v] = true
			}
		}
	}
	if len(box) <= 1 {
		return true
	}
	return false
}
