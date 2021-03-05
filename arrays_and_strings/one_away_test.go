package arrays_and_strings

import "testing"

func TestOneAway(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		output bool
	}{
		{"First", []string{"pale", "ple"}, true},
		{"Second", []string{"pales", "pale"}, true},
		{"Third", []string{"pale", "bale"}, true},
		{"Fourth", []string{"pale", "bake"}, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := OneAway(test.input[0], test.input[1])
			if result != test.output {
				t.Logf("For %v, answer should be %v not %v", test.input, test.output, result)
				t.Fail()
			}
		})
	}
}

func OneAway(a string, b string) bool {
	diff := len(a) - len(b)
	if diff < 0 {
		diff *= -1
	}

	box := make(map[rune]int)
	for _, v := range a {
		box[v] += 1
	}

	for _, v := range b {
		if box[v] >= 1 {
			box[v] -= 1
			if box[v] == 0 {
				delete(box, v)
			}
		} else {
			box[v] += 1
		}
	}
	if len(box) > 2 {
		return false
	}
	return true
}
