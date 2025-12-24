package piscine

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}

	result := ""
	for i, str := range args {
		result += str
		if i < len(args)-1 {
			result += "\n"
		}
	}
	return result
}
