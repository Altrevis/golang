package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	result := int(1)
	for i := 1; i <= nb; i++ {
		// Vérifier les dépassements possibles
		if result > (1<<63-1)/int(i) {
			return 0
		}
		result *= int(i)
	}
	return result
}
