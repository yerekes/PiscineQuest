package piscine

func CheckAlpha(a rune) bool {
	if a >= 'a' && a <= 'z' || a >= 'A' && a <= 'Z' || a >= '0' && a <= '9' {
		return true
	} else {
		return false
	}
}

func IsAlpha(s string) bool {
	runes := []rune(s)
	for i := range s {
		if CheckAlpha(runes[i]) == false {
			return false
		}
	}
	return true
}
