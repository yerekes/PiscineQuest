package piscine

import "github.com/01-edu/z01"

func isSafe(nb, value int, pos [8]int) bool {
	for i := 0; i < nb; i++ {
		if pos[i] == value || pos[i] == value-(nb-i) || pos[i] == value+(nb-i) {
			return false
		}
	}
	return true
}

func resolve(nb int, pos [8]int) {
	if nb == 8 {
		for _, value := range pos {
			z01.PrintRune(rune(value + '0'))
		}
		z01.PrintRune('\n')
	} else {
		for i := 1; i <= 8; i++ {
			if isSafe(nb, i, pos) {
				pos[nb] = i
				resolve(nb+1, pos) // a chaque fois que le fonction est appeler le boucle for est reinitialiser
			}
		}
	}
}

func EightQueens() {
	pos := [8]int{0, 0, 0, 0, 0, 0, 0, 0}
	resolve(0, pos)
}
