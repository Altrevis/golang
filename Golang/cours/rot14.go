package piscine

func Rot14(s string) string {
	rotated := make([]byte, len(s))
	for i, char := range s {
		switch {
		case 'a' <= char && char <= 'z':
			rotated[i] = byte(((int(char-'a') + 14) % 26) + 'a')
		case 'A' <= char && char <= 'Z':
			rotated[i] = byte(((int(char-'A') + 14) % 26) + 'A')
		default:
			rotated[i] = byte(char)
		}
	}
	return string(rotated)
}
