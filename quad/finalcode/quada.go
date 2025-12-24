package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	if x < 1 || y < 1 {
		return
	}
	for i := 1; i <= y; i++ {
		if i == 1 || i == y {
			for j := 1; j <= x; j++ {
				if j == 1 || j == x {
					z01.PrintRune('o')
				} else {
					z01.PrintRune('-')
				}
			}
		} else {
			for j := 1; j <= x; j++ {
				if j == 1 || j == x {
					z01.PrintRune('|')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}
