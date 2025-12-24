package main

import (
	"os"
)

const (
	MaxInt64 = 9223372036854775807
	MinInt64 = -9223372036854775808
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}

	val1, ok1 := atoi(args[0])
	val2, ok2 := atoi(args[2])
	op := args[1]

	if !ok1 || !ok2 {
		return
	}

	switch op {
	case "+":
		res, ok := add(val1, val2)
		if ok {
			printNbr(res)
		}
	case "-":
		res, ok := sub(val1, val2)
		if ok {
			printNbr(res)
		}
	case "*":
		res, ok := mul(val1, val2)
		if ok {
			printNbr(res)
		}
	case "/":
		if val2 == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		res, ok := div(val1, val2)
		if ok {
			printNbr(res)
		}
	case "%":
		if val2 == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
		res, ok := mod(val1, val2)
		if ok {
			printNbr(res)
		}
	default:
		return
	}
}

func atoi(s string) (int64, bool) {
	if s == "" {
		return 0, false
	}
	neg := false
	i := 0
	if s[0] == '-' {
		neg = true
		i++
	} else if s[0] == '+' {
		i++
	}

	if i == len(s) {
		return 0, false
	}

	var n int64 = 0
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, false
		}
		digit := int64(c - '0')

		if neg {
			if n < MinInt64/10 {
				return 0, false
			}
			n *= 10
			if n < MinInt64+digit {
				return 0, false
			}
			n -= digit
		} else {
			if n > MaxInt64/10 {
				return 0, false
			}
			n *= 10
			if n > MaxInt64-digit {
				return 0, false
			}
			n += digit
		}
	}
	return n, true
}

func add(a, b int64) (int64, bool) {
	sum := a + b
	if (sum^a) < 0 && (sum^b) < 0 {
		return 0, false
	}
	return sum, true
}

func sub(a, b int64) (int64, bool) {
	diff := a - b
	if (a^b) < 0 && (a^diff) < 0 {
		return 0, false
	}
	return diff, true
}

func mul(a, b int64) (int64, bool) {
	if a == 0 || b == 0 {
		return 0, true
	}
	res := a * b
	if (a == -1 && b == MinInt64) || (b == -1 && a == MinInt64) {
		return 0, false
	}
	if res/b != a {
		return 0, false
	}
	return res, true
}

func div(a, b int64) (int64, bool) {
	if a == MinInt64 && b == -1 {
		return 0, false
	}
	return a / b, true
}

func mod(a, b int64) (int64, bool) {
	return a % b, true
}

func printNbr(n int64) {
	if n == 0 {
		os.Stdout.WriteString("0\n")
		return
	}

	if n == MinInt64 {
		os.Stdout.WriteString("-9223372036854775808\n")
		return
	}

	if n < 0 {
		os.Stdout.WriteString("-")
		n = -n
	}

	var buf [24]byte
	idx := 23
	buf[idx] = '\n'
	idx--

	for n > 0 {
		buf[idx] = byte(n%10 + '0')
		idx--
		n /= 10
	}
	os.Stdout.Write(buf[idx+1:])
}
