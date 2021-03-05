package coding_interview

import (
	"coding_interview/arrays_and_strings"
	"reflect"
	"testing"
)

func TestIsSubstring(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		output bool
	}{
		{"Substring", []string{"hello", "ell"}, true},
		{"Substring 2", []string{"ell", "Hello"}, true},
		{"Not substring", []string{"elxl", "Hello"}, false},
		{"Substring", []string{"", ""}, true},
		{"Substring", []string{"Hello", "Hello"}, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := arrays_and_strings.IsSubstring(test.input[0], test.input[1])
			if !reflect.DeepEqual(result, test.output) {
				t.Logf("%v: For %v, the expected answer is %v not %v", test.name, test.input, test.output, result)
				t.Fail()
			}

		})
	}
}

func TestIsRotation(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		output bool
	}{
		{"Rotation", []string{"waterbottle", "erbottlewat"}, true},
		{"Not rotation", []string{"waterbottl", "erbottlewat"}, false},
		{"Not rotation", []string{"", ""}, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := arrays_and_strings.IsRotation(test.input[0], test.input[1])
			if !reflect.DeepEqual(result, test.output) {
				t.Logf("%v: For %v, the expected answer is %v not %v", test.name, test.input, test.output, result)
				t.Fail()
			}

		})
	}
}

