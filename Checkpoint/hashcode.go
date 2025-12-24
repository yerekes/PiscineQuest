package main

func HashCode(dec string) string {
	result := ""
	for _, r := range dec {
		formula := (int(r) + len(dec)) % 127
		if formula <= 31 {
			formula += 33
		}

		result += string(rune(formula))
	}
	return result
}

// func main() {
// 	fmt.Println(HashCode("A"))
// 	fmt.Println(HashCode("AB"))
// 	fmt.Println(HashCode("BAC"))
// 	fmt.Println(HashCode("Hello World"))
// }
