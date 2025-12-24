package main

func Atoi(s string) int {
	var n int
	for _, char := range s {
		if char > '9' || char < '0' {
			return 0
		}
		n = n*10 + int(char-'0')
	}
	return n
}

func IsPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// func main() {
// 	if
// 	var result int
// 	n := Atoi(os.Args[1])
// 	for i := 2; i <= n; i++ {
// 		if IsPrime(i) {
// 			result += i
// 		}
// 	}
// 	println(result)
// }
