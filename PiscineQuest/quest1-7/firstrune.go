package piscine

func FirstRune(s string) rune {
	runes := []rune(s)
	return runes[len(runes)-len(runes)]
}
