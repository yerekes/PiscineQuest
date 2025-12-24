package main

func CanJump(arr []uint) bool {
	if len(arr) == 0 {
		return false
	}
	i := 0
	lastindex := len(arr) - 1
	for i < lastindex {
		steps := int(arr[i])
		if arr[i] == 0 {
			return false
		}
		i += steps
	}
	return i == lastindex
}

// func main() {
// 	input1 := []uint{4, 3, 1, 1, 4}
// 	fmt.Println(CanJump(input1))

// 	input2 := []uint{3, 2, 1, 0, 4}
// 	fmt.Println(CanJump(input2))

// 	input3 := []uint{0}
// 	fmt.Println(CanJump(input3))
// }
