package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	last := 0
	for i, r := range name {
		if r == '.' {
			last = i + 1
		}
	}
	for _, r := range name[last:] {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
