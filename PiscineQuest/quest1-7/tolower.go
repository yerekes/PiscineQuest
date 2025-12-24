package piscine

func ToLower(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		letter := s[i]
		if s[i] >= 'A' && s[i] <= 'Z' {
			letter += 32
		}
		result += string(letter)
	}
	return result
}
