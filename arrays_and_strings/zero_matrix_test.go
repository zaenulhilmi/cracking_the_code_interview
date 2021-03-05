package arrays_and_strings

import (
	"reflect"
	"testing"
)

func TestZeroMatrix(t *testing.T) {
	tests := []struct {
		name   string
		input  [][]int
		output [][]int
	}{
		{"Not Change", [][]int{{1, 2, 3}, {4, 5}, {6, 7, 8, 9}}, [][]int{{1, 2, 3}, {4, 5}, {6, 7, 8, 9}}},
		{"Change to zero", [][]int{{1, 2, 0}, {4, 5}, {6, 7, 8, 9}}, [][]int{{0, 0, 0}, {0, 0}, {0, 0, 0, 0}}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ZeroMatrix(test.input)
			if !reflect.DeepEqual(result, test.output) {
				t.Logf("%v: For %v, the expected answer is %v not %v", test.name, test.input, test.output, result)
				t.Fail()
			}
		})
	}
}

