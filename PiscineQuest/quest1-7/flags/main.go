package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	insertStr := ""
	orderFlag := false
	mainStr := ""
	printHelp := false

	if len(args) == 0 {
		printHelp = true
	}

	for _, arg := range args {

		if arg == "--help" || arg == "-h" {
			printHelp = true
			break
		}

		if len(arg) >= 9 && arg[:9] == "--insert=" {
			insertStr = arg[9:]
			continue
		}
		if len(arg) >= 3 && arg[:3] == "-i=" {
			insertStr = arg[3:]
			continue
		}

		if arg == "--order" || arg == "-o" {
			orderFlag = true
			continue
		}

		mainStr = arg
	}

	if printHelp {
		printHelpText()
		return
	}

	if insertStr != "" {
		mainStr += insertStr
	}

	if orderFlag {
		mainStr = sortString(mainStr)
	}

	for _, r := range mainStr {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func printHelpText() {
	printLine("--insert")
	printLine("  -i")
	printLine("\t This flag inserts the string into the string passed as argument.")
	printLine("--order")
	printLine("  -o")
	printLine("\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func printLine(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func sortString(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)-1; i++ {
		for j := 0; j < len(runes)-1-i; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}
	return string(runes)
}
