package piscine

func checkup(a rune) bool {
	if a >= 'A' && a <= 'Z' {
		return true
	} else {
		return false
	}
}

func IsUpper(s string) bool {
	runes := []rune(s)
	for i := range s {
		if checkup(runes[i]) == false {
			return false
		}
	}
	return true
}
