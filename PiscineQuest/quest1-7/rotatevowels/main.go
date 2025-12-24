package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		z01.PrintRune('\n')
		return
	}

	// 1. Собираем все индексы гласных по всем аргументам
	type pos struct {
		argIndex  int
		charIndex int
	}
	var vowels []pos
	runesArgs := make([][]rune, len(args))

	for i, arg := range args {
		runes := []rune(arg)
		runesArgs[i] = runes
		for j, r := range runes {
			if isVowel(r) {
				vowels = append(vowels, pos{i, j})
			}
		}
	}

	// 2. "Зеркалим" гласные по глобальной последовательности
	n := len(vowels)
	for i := 0; i < n/2; i++ {
		a := vowels[i]
		b := vowels[n-1-i]
		runesArgs[a.argIndex][a.charIndex], runesArgs[b.argIndex][b.charIndex] = runesArgs[b.argIndex][b.charIndex], runesArgs[a.argIndex][a.charIndex]
	}

	// 3. Печатаем результат
	for i, runes := range runesArgs {
		for _, r := range runes {
			z01.PrintRune(r)
		}
		if i != len(runesArgs)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

// Проверка гласной
func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
