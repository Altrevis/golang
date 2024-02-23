package piscine

func SortIntegerTable(table []int) {
	n := len(table)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if table[i-1] > table[i] {
				table[i-1], table[i] = table[i], table[i-1]
				swapped = true
			}
		}
	}
}
