package main

//	func ConcatSlice(slice1, slice2 []int) []int {
//		snake := []int{}
//		if len(slice1) < len(slice2) {
//			slice1, slice2 = slice2, slice1
//		}
//		for _, n := range slice1 {
//			snake = append(snake, n)
//		}
//		for _, r := range slice2 {
//			snake = append(snake, r)
//		}
//		return snake
//	}
func ConcatSlice(slice1, slice2 []int) (result []int) {
	result = append(result, slice1...)
	result = append(result, slice2...)
	return result
}

// func main() {
// 	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
// 	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
// 	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))
// }
