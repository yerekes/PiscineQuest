package piscine

func CheckNumber(a rune) bool {
	if a >= '0' && a <= '9' {
		return true
	} else {
		return false
	}
}

func IsNumeric(s string) bool {
	runes := []rune(s)
	for i := range s {
		if CheckNumber(runes[i]) == false {
			return false
		}
	}
	return true
}
