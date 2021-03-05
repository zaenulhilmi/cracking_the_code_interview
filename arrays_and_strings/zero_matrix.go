package arrays_and_strings

func ZeroMatrix(input [][]int) [][]int {
	hasZero := false
	for _, v := range input {
		for _, val := range v {
			if val == 0 {
				hasZero = true
				break
			}
		}
		if hasZero {
			break
		}
	}

	if hasZero {
		for i, v := range input {
			for j, _ := range v {
				input[i][j] = 0
			}
		}
	}

	return input
}
