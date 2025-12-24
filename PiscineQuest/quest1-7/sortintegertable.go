package piscine

func SortIntegerTable(table []int) {
	for num := len(table) - 1; num > 0; num-- {
		for i := 0; i < num; i++ {
			if table[i] > table[i+1] {
				temp := table[i]
				table[i] = table[i+1]
				table[i+1] = temp
			}
		}
	}
}
