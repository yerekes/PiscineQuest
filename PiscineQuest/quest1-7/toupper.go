package piscine

func ToUpper(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		letter := s[i]
		if s[i] >= 'a' && s[i] <= 'z' {
			letter -= 32
		}
		result += string(letter)
	}
	return result
}
