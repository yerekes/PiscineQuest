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
	QuadB(x, y)
}

func QuadB(x, y int) {
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 || i == y && j == x {
				fmt.Print("/")
			} else if i == 1 && j == x || i == y && j == 1 {
				fmt.Print("\\")
			} else if i == 1 || i == y || j == 1 || j == x {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
