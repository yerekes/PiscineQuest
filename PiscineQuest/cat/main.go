package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func efafe() {
	args := os.Args[1:]

	if len(args) == 0 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			printStr("ERROR: ")
			printStr(err.Error())
			z01.PrintRune('\n')
			os.Exit(1)
		}
	} else {
		for _, fileName := range args {
			file, err := os.Open(fileName)
			if err != nil {
				printStr("ERROR: ")
				printStr(err.Error())
				z01.PrintRune('\n')
				os.Exit(1)
			}
			_, err = io.Copy(os.Stdout, file)
			file.Close()
			if err != nil {
				printStr("ERROR: ")
				printStr(err.Error())
				z01.PrintRune('\n')
				os.Exit(1)
			}
		}
	}
}
