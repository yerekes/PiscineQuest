package piscine

import "github.com/01-edu/z01"

func QuadD(x, y int) {
	if x < 1 || y < 1 {
		return
	} else {
		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				if i == 0 && j == 0 {
					z01.PrintRune('A')
				} else if i == 0 && j == x-1 {
					z01.PrintRune('C')
				} else if i == y-1 && j == 0 {
					z01.PrintRune('A')
				} else if i == y-1 && j == x-1 {
					z01.PrintRune('C')
				} else if i == 0 || j == 0 || i == y-1 || j == x-1 {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
