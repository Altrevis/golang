package piscine

func Capitalize(s string) string {
	result := make([]rune, len(s))
	capitalizeNext := true

	for i, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			if capitalizeNext {
				if char >= 'a' && char <= 'z' {
					result[i] = char - 32
				} else {
					result[i] = char
				}
				capitalizeNext = false
			} else {
				if char >= 'A' && char <= 'Z' {
					result[i] = char + 32
				} else {
					result[i] = char
				}
			}
		} else {
			result[i] = char
			capitalizeNext = true
		}
	}
	return string(result)
}
