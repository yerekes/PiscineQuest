package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	x, _ := strconv.Atoi(os.Args[1])
	y, _ := strconv.Atoi(os.Args[2])
	QuadC(x, y)
}

func QuadC(x, y int) {
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && (j == 1 || j == x) {
				fmt.Print("A")
			} else if i == y && (j == 1 || j == x) {
				fmt.Print("C")
			} else if i == 1 || i == y || j == 1 || j == x {
				fmt.Print("B")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
