package arrays_and_strings

func IsUnique(input string) bool {
	dict := make(map[rune]int)
	for _, v := range input {
		if dict[v] == 1 {
			return false
		}
		dict[v] = 1
	}
	return true
}

