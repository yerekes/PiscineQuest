package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	upperFlag := false

	// Проверяем флаг
	if len(args) > 0 && args[0] == "--upper" {
		upperFlag = true
		args = args[1:]
	}

	printedSomething := false

	for _, arg := range args {
		n := atoi(arg)

		if n < 1 || n > 26 {
			z01.PrintRune(' ')
			printedSomething = true
			continue
		}

		var r rune
		if upperFlag {
			r = rune('A' + n - 1)
		} else {
			r = rune('a' + n - 1)
		}

		z01.PrintRune(r)
		printedSomething = true
	}

	// Печатаем \n ТОЛЬКО если  что-то напечатано
	if printedSomething {
		z01.PrintRune('\n')
	}
}

// atoi — своя версия преобразования строки в число
func atoi(s string) int {
	result := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return -1
		}
		result = result*10 + int(r-'0')
	}
	return result
}
