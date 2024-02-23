package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for i := 0; i < len(s); i++ {
		char := rune(s[i])
		z01.PrintRune(char)
	}
}
