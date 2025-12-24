package piscine

func QuadD(x, y int) {
	if x >= 1 && y >= 1 {
		for i := 1; i <= y; i++ {
			for j := 1; j <= x; j++ {
				if i == 1 && j == 1 {
					print("A")
				} else if i == y && j == 1 {
					print("A")
				} else if i == 1 && j == x {
					print("C")
				} else if i == y && j == x {
					print("C")
				} else if j == 1 || i == 1 || i == y || j == x {
					print("B")
				} else {
					print(" ")
				}
			}
			print("\n")
		}
	}
}
