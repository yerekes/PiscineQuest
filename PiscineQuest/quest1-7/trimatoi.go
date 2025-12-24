package piscine

func TrimAtoi(s string) int {
	sign := 1
	started := false
	num := 0

	for _, r := range s {
		// Проверка на знак ДО первой цифры
		if !started {
			if r == '-' {
				sign = -1
				continue
			}
			if r == '+' {
				continue
			}
		}

		// Если это цифра — начинаем собирать число
		if r >= '0' && r <= '9' {
			started = true
			num = num*10 + int(r-'0')
		}
	}

	return sign * num
}
