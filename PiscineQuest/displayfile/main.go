package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	filename := os.Args[1]
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("", err)
		return
	}
	fmt.Print(string(data))
}
