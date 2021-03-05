package arrays_and_strings

func RotateMatrix(input [][]int) [][]int {
	length := len(input)
	result := make([][]int, length)
	for i, _ := range result {
		result[i] = make([]int, length)
	}

	for i, v := range input {
		for j, _ := range v {
			result[i][j] = input[length-1-j][i]
		}
	}
	return result
}
