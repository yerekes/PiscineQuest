package piscine

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	var r int = 0
	i := 0
	checkminus := false
	if s[i] == '-' {
		checkminus = true
		i++
	} else if s[i] == '+' {
		checkminus = false
		i++
	}
	for i < len(s) {
		if s[i] < 48 || s[i] > 57 {
			return 0
		}

		r = r*10 + int(s[i]-48)
		i++
	}
	if checkminus == true {
		r *= -1
	}

	return r
}
