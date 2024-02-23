package main

import "github.com/01-edu/z01"

func main() {
	for j := 122; j >= 97; j-- {
		z01.PrintRune(rune(j))
	}
	z01.PrintRune('\n')
}
