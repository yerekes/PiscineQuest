package main

import "fmt"

func Chunk(slice []int, size int) {
	if size == 0 {
		fmt.Println()
		return
	}
	if len(slice) == 0 {
		fmt.Println(slice)
		return
	}
	start := 0
	var r [][]int
	for i := 1; i < len(slice); i++ {
		if i%size == 0 {
			r = append(r, slice[start:i])
			start = i
		}
	}
	r = append(r, slice[start:])
	fmt.Println(r)
}

// func main() {
// 	Chunk([]int{}, 3)
// 	Chunk([]int{1, 6, 6}, 10)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
// }
