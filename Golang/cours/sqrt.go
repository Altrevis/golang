package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}

	x := nb
	y := (x + 1) / 2

	for x > y {
		x = (x + y) / 2
		y = nb / x
	}

	if x*x == nb {
		return x
	} else {
		return 0
	}
}
