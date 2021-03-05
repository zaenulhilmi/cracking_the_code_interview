package arrays_and_strings

import (
	"reflect"
	"testing"
)

func TestRotateMatrix(t *testing.T) {
	tests := []struct {
		name   string
		input  [][]int
		output [][]int
	}{
		{"First",
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9}},
			[][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			result := RotateMatrix(test.input)
			if !reflect.DeepEqual(result, test.output) {
				t.Logf("For %v, the answer should be %v not %v", test.input, test.output, result)
				t.Fail()
			}
		})
	}
}

