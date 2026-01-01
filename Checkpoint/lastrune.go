package main 

import  "github.com/01-edu/z01"

func LastRune(s string) rune{
	for _,r :=range s{
		if r==len(s)-1{
			return r
		}
	}
	return 0
}
func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune('\n')
}