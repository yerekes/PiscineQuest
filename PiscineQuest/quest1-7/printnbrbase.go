package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if len(base) < 2 {
		printNV()
		return
	}
	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			printNV()
			return
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				printNV()
				return
			}
		}
	}
	if nbr < 0 {
		z01.PrintRune('-')
	}
	printBase(nbr, base)
}

func printBase(nbr int, base string) {
	b := len(base)
	var digit int
	if nbr < 0 {
		digit = -(nbr % b)
	} else {
		digit = nbr % b
	}
	next := nbr / b
	if next != 0 {
		printBase(next, base)
	}
	z01.PrintRune(rune(base[digit]))
}

func printNV() {
	z01.PrintRune('N')
	z01.PrintRune('V')
}
