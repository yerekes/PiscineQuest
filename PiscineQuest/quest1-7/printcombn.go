package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	var comb func(start int, depth int, arr []rune)
	comb = func(start int, depth int, arr []rune) {
		if depth == n {
			for _, ch := range arr {
				z01.PrintRune(ch)
			}

			isLast := true
			for i := 0; i < n; i++ {
				if arr[i] != rune('0'+i+10-n) {
					isLast = false
					break
				}
			}
			if !isLast {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			return
		}

		for i := start; i <= 10-(n-depth); i++ {
			comb(i+1, depth+1, append(arr, rune('0'+i)))
		}
	}

	comb(0, 0, []rune{})
	z01.PrintRune('\n')
}
