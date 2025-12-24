package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	// Если 0 — выводим сразу
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Делаем число положительным
	if n < 0 {
		n = -n
	}

	// Разбираем число на цифры
	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	// Сортировка пузырьком
	for i := 0; i < len(digits); i++ {
		for j := 0; j < len(digits)-1; j++ {
			if digits[j] > digits[j+1] {
				digits[j], digits[j+1] = digits[j+1], digits[j]
			}
		}
	}

	// Печать цифр
	for _, d := range digits {
		z01.PrintRune(rune(d + '0'))
	}
}
