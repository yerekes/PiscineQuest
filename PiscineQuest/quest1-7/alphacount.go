package piscine

func AlphaCount(s string) int {
	nb := 0
	runes := []rune(s)
	for _, letter := range runes {
		if letter >= 'a' && letter <= 'z' || letter >= 'A' && letter <= 'Z' {
			nb++
		}
	}
	return nb
}
