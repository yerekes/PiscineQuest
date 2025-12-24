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
	QuadD(x, y)
}

func QuadD(x, y int) {
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if j == 1 && (i == 1 || i == y) {
				fmt.Print("A")
			} else if j == x && (i == 1 || i == y) {
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
