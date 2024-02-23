package piscine

func Swap(a *int, b *int) {
	aa := *a
	*a = *b
	*b = aa
}
