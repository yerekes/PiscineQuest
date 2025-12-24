package piscine

func checup(a rune) bool {
	if a >= 'a' && a <= 'z' {
		return true
	} else {
		return false
	}
}

func IsLower(s string) bool {
	runes := []rune(s)
	for i := range s {
		if checup(runes[i]) == false {
			return false
		}
	}
	return true
}
