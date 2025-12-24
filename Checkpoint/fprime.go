package main

import (
	"fmt"
	"os"
	"strconv"
)

// func isPrime(n int) bool {
// 	if n <= 1 {
// 		return false
// 	}
// 	for i := 2; i*i <= n; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

func Fprime() {
	if len(os.Args) != 2 {
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 1 {
		return
	}

	first := true
	num := n

	for i := 2; i*i <= num; i++ {
		for num%i == 0 {
			if !first {
				fmt.Print("*")
			}
			fmt.Print(i)
			first = false
			num /= i
		}
	}
	if num > 1 {
		if !first {
			fmt.Print("*")
		}
		fmt.Print(num)
	}

	fmt.Println()
}
