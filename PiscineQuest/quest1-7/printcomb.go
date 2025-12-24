package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '7'; i++ {
		for k := i + 1; k <= '8'; k++ {
			for j := k + 1; j <= '9'; j++ {
				z01.PrintRune(i)
				z01.PrintRune(k)
				z01.PrintRune(j)
				if i != '7' || k != '8' || j != '9' {
					z01.PrintRune(',')
					z01.PrintRune(32)
				}
			}
		}
	}
	z01.PrintRune(10)
}
