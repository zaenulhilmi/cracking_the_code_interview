package arrays_and_strings

import (
	"strconv"
	"testing"
)

func TestStringCompression(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{"first", "aaabbcdddaa", "a3b2c1d3a2"},
		{"second", "aaabbcda", "aaabbcda"},
		{"first", "aaabbcddda", "aaabbcddda"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Compress(test.input)
			if result != test.output {
				t.Logf("for %v the answer should be %v not %v", test.input, test.output, result)
				t.Fail()
			}
		})
	}
}

func Compress(input string) string {
	temp := ""
	curr := '_'
	count := 1

	for _, v := range input {
		if curr != v {
			if curr != '_' {
				temp += string(curr) + strconv.Itoa(count)
			}
			curr = v
			count = 1
		} else {
			count += 1
		}
	}

	temp += string(curr) + strconv.Itoa(count)
	if len(temp) >= len(input) {
		return input
	}
	return temp
}
