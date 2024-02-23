package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		z01.PrintRune('N')
		z01.PrintRune('e')
		z01.PrintRune('g')
		z01.PrintRune('a')
		z01.PrintRune('t')
		z01.PrintRune('i')
		z01.PrintRune('v')
		z01.PrintRune('e')
		z01.PrintRune(' ')
		z01.PrintRune('n')
		z01.PrintRune('u')
		z01.PrintRune('m')
		z01.PrintRune('b')
		z01.PrintRune('e')
		z01.PrintRune('r')
		z01.PrintRune('s')
		z01.PrintRune(' ')
		z01.PrintRune('a')
		z01.PrintRune('r')
		z01.PrintRune('e')
		z01.PrintRune(' ')
		z01.PrintRune('n')
		z01.PrintRune('o')
		z01.PrintRune('t')
		z01.PrintRune(' ')
		z01.PrintRune('a')
		z01.PrintRune('l')
		z01.PrintRune('l')
		z01.PrintRune('o')
		z01.PrintRune('w')
		z01.PrintRune('e')
		z01.PrintRune('d')
		z01.PrintRune('.')
		z01.PrintRune('\n')
		return
	}

	if n == 0 {
		z01.PrintRune('0')
		return
	}

	digits := [10]int{}

	for n > 0 {
		digit := n % 10
		digits[digit]++
		n /= 10
	}

	for i, count := range digits {
		for j := 0; j < count; j++ {
			z01.PrintRune(rune('0' + byte(i)))
		}
	}
}
