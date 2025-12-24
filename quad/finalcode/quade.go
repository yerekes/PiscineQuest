package piscine

import "github.com/01-edu/z01"

func QuadE(x, y int) {
	if x < 1 || y < 1 {
		return
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			switch {
			case i == 0 && j == 0:
				z01.PrintRune('A')
			case i == 0 && j == x-1:
				z01.PrintRune('C')
			case i == y-1 && j == 0:
				z01.PrintRune('C')
			case i == y-1 && j == x-1:
				z01.PrintRune('A')
			case i == 0 || j == 0 || i == y-1 || j == x-1:
				z01.PrintRune('B')
			default:
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
