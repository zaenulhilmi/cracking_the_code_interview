package arrays_and_strings

func IsRotation(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	if !IsSubstring(a, b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a == b {
			return true
		} else {
			b = string(b[len(b)-1]) + b[:len(b)-1]
			if a == b {
				return true
			}
		}
	}

	return false
}

func IsSubstring(a string, b string) bool {
	var longer, shorter string
	if len(a) > len(b) {
		longer = a
		shorter = b
	} else {
		longer = b
		shorter = a
	}

	boxLonger := make(map[rune]int)
	for _, v := range longer {
		boxLonger[v] += 1
	}

	for _, v := range shorter {
		if boxLonger[v] == 0 {
			return false
		} else {
			boxLonger[v] -= 1
		}
	}
	return true
}
