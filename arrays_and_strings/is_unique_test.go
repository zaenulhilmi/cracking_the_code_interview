package arrays_and_strings

import "testing"

func TestIsUnique(t *testing.T) {
	tests := []struct {
		name string
		input string
		output bool
	}{
		{"double", "hello", false},
		{"unique", "abc", true},
		{"no input", "", true},
		{"double space", "a  b", false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsUnique(test.input)
			if result != test.output {
				t.Logf("For %v of input %v, the answer should be %v not %v", test.name, test.input, test.output, result)
				t.Fail()
			}
		})
	}
}


