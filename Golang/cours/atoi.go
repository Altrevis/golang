package piscine

func Atoi(str string) int {
	var result int
	sign := 1

	for i, char := range str {
		if i == 0 && char == '-' {
			sign = -1
			continue
		}
		if i == 0 && char == '+' {
			continue
		}
		if char >= '0' && char <= '9' {
			digit := int(char - '0')
			result = result*10 + digit
		} else {
			return 0
		}
	}
	return result * sign
}
