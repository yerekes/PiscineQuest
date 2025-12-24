package main

func FirstWord(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i-1] == ' ' {
			return string(s[i:]) + "\n"
		}
	}
	return s + "\n"
}

// func main() {
// 	fmt.Print(FirstWord("hello there"))
// 	fmt.Print(FirstWord(""))
// 	fmt.Print(FirstWord("hello   .........  bye"))
// }
