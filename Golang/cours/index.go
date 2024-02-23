package piscine

func Index(s string, toFind string) int {
	sLen := len(s)
	toFindLen := len(toFind)

	for i := 0; i <= sLen-toFindLen; i++ {
		if s[i:i+toFindLen] == toFind {
			return i
		}
	}
	return -1
}
