package main

func Itoa(n int) string {
	// return strconv.Itoa(n)
	sign := ""
	if n < 0 {
		n = -n
		sign += "-"
	}
	if n == 0 {
		return string(n%10 + '0')
	}
	var s string
	for n > 0 {
		s = string(n%10+'0') + s
		n /= 10
	}
	return sign + s
}

// func main() {
// 	fmt.Println(Itoa1("12345"))
// 	fmt.Println(Itoa(0))
// 	fmt.Println(Itoa(-1234))
// 	fmt.Println(Itoa(987654321))
// }

func Itoa1(n string) int {
	result := 0
	for i := 0; i < len(n); i++ {
		result = result*10 + int(n[i]-'0')
	}
	return result
}
