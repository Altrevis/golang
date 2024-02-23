package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	}
	fact := int(nb) * RecursiveFactorial(nb-1)
	if fact < 0 {
		return 0
	}
	return fact
}
