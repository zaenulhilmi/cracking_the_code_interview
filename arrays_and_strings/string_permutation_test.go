package arrays_and_strings

import "testing"

func TestIsPermutateEachOther(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		output bool
	}{
		{"permutate each other", []string{"abc", "bca"}, true},
		{"not permutate each other", []string{"abbc", "bcaa"}, false},
		{"not permutate each other", []string{"abc", "bcx"}, false},
		{"different length", []string{"abc", "bcaa"}, false},
		{"empty char", []string{"", ""}, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := PermutateEachOther(test.input[0], test.input[1])
			if result != test.output {
				t.Logf("For %v the answer should be %v not %v", test.input, test.output, result)
				t.Fail()
			}
		})
	}
}

