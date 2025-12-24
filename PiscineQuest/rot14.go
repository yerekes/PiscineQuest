package piscine

func Rot14(s string) string {
	result := ""
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			if r+14 > 90 {
				result += string(65 + (14 - (91 - r)))
			} else {
				result += string(r + 14)
			}
		} else if r >= 'a' && r <= 'z' {
			if r+14 > 122 {
				result += string(97 + (14 - (123 - r)))
			} else {
				result += string(r + 14)
			}
		} else {
			result += string(r)
		}
	}
	return result
}
