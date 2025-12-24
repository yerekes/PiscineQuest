package main

import "os"

func ReverseStrCap() {
	args := os.Args[1:]
	for _, str := range args {
		for i := 0; i < len(str); i++ {
			if str[i] >= 'a' && str[i] <= 'z' && ((i+1 < len(str) && str[i+1] == ' ') || (i == len(str)-1)) {
				os.Stdout.WriteString(string(str[i] - 32))
			} else if str[i] >= 'A' && str[i] <= 'Z' && i+1 < len(str) && str[i+1] != ' ' && i != len(str)-1 {
				os.Stdout.WriteString(string(str[i] + 32))
			} else {
				os.Stdout.WriteString(string(str[i]))
			}
		}
		os.Stdout.WriteString("\n")
	}
}
