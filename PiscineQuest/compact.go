package piscine

func Compact(ptr *[]string) int {
	if ptr == nil || len(*ptr) == 0 {
		return 0
	}

	slice := *ptr
	j := 0

	for _, v := range slice {
		if v != "" {
			slice[j] = v
			j++
		}
	}

	*ptr = slice[:j]

	return j
}
