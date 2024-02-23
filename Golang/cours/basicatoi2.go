package piscine

func BasicAtoi2(s string) int {
	result := 0
	for _, char := range s {
		if char >= 48 && char <= 57 {
			result = result*10 + int(char-'0')
		} else {
			return 0
		}
	}
	return result
}
