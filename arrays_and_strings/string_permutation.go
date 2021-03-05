package arrays_and_strings

func PermutateEachOther(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	dict := make(map[rune]int)
	for _, v := range a {
		dict[v] += 1
	}
	for _, v := range b {
		dict[v] -= 1
		if dict[v] == -1 {
			return false
		}
	}
	return true
}
