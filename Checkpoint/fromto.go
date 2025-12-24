package main

func FromTo(a, b int) string {
	if a > 99 || b > 99 || a < 0 || b < 0 {
		return "Invalid\n"
	}
	s := 1
	if a > b {
		s = -1
	}
	n := ""
	for i := a; ; i += s {
		if i < 10 {
			n += "0"
		}

		n += Itoa2(i)

		if i == b {
			break
		}
		n += ", "
	}
	return n + "\n"
}

func Itoa2(d int) string {
	if d == 0 {
		return "0"
	}
	res := ""
	for d > 0 {
		res = string(d%10+'0') + res
		d /= 10
	}
	return res
}

// func main() {
// 	fmt.Print(FromTo(1, 10))
// 	fmt.Print(FromTo(10, 1))
// 	fmt.Print(FromTo(10, 10))
// 	fmt.Print(FromTo(100, 10))
// }
