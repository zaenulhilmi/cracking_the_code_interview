package arrays_and_strings

import (
	"strings"
	"testing"
)

func TestUrlify(t *testing.T) {
	tests := []struct {
		name string
		input string
		output string
	}{
		{"trim whitespace", "   abc  ", "abc"},
		{"urlify", "   a bc  ", "a%20bc"},
		{"example", "Mr John Smith   ", "Mr%20John%20Smith"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Urlify(test.input)
			if result != test.output {
				t.Logf("For %v, the answer should be %v not %v", test.input, test.output, result)
				t.Fail()
			}
		})
	}
}

func Urlify(input string) string {
	newString := strings.TrimSpace(input)
	result := ""
	for _, v := range newString {
		if v == ' '	 {
			result += "%20"
		} else {
			result += string(v)
		}
	}
	return result
}

