package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	newWord := true

	for i, r := range runes {
		if isAlnum(r) {
			if newWord {
				// первая буква слова — заглавная
				if r >= 'a' && r <= 'z' {
					runes[i] = r - ('a' - 'A')
				}
				newWord = false
			} else {
				// остальные буквы — строчные
				if r >= 'A' && r <= 'Z' {
					runes[i] = r + ('a' - 'A')
				}
			}
		} else {
			newWord = true
		}
	}

	return string(runes)
}

func isAlnum(r rune) bool {
	return (r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z') ||
		(r >= '0' && r <= '9')
}
