package main

func ZipString(s string) string {
	c := 0
	r := ""
	for i := 0; i < len(s); i += c {
		c = 0

		for j := i; j < len(s); j++ {
			if s[j] != s[i] {
				break
			}
			c++
		}
		r += Itoa3(c) + string(s[i])
	}
	return r
}

func Itoa3(n int) (s string) {
	for n > 0 {
		s = string(rune(n%10+'0')) + s
		n /= 10
	}
	return s
}

// func main() {
// 	fmt.Println(ZipString("YouuungFellllas"))
// 	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
// 	fmt.Println(ZipString("Helloo Therre!"))
// }
