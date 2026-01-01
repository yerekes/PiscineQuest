package main 

// import  "github.com/01-edu/z01"

func FirstRune(s string)rune{
// 	for _,r := range s{
// 		return r
// 	}
// 	return 0
runes:=[]rune(s)
return runes[0]
}

// func main() {
// 	z01.PrintRune(FirstRune("Hello!"))
// 	z01.PrintRune(FirstRune("Salut!"))
// 	z01.PrintRune(FirstRune("Ola!"))
// 	z01.PrintRune('\n')
// }
