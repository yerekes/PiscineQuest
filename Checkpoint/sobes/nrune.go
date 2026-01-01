package main

import "github.com/01-edu/z01"

// import "github.com/01-edu/z01"

func NRune(s string, n int) rune {
	// var last rune
	// for i,r := range s{
	// 	if i+1 == n{
	// 		last = r
	// 		break
	// 	}
	// }
	// return last
	runes := []rune(s)
	if n > len(s) {
		return 0
	}
	if n <= 0 {
		return 0
	}
	return runes[n-1]
}

func main() {
	z01.PrintRune(NRune("Hello!", 3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Bye!", -1))
	z01.PrintRune(NRune("Bye!", 5))
	z01.PrintRune(NRune("Ola!", 4))
	z01.PrintRune('\n')
}
