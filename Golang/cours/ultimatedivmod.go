package piscine

func UltimateDivMod(a *int, b *int) {
	aa := *a
	*a = *a / *b
	*b = aa % *b
}
