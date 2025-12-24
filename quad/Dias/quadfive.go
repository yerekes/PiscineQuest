package piscine

func QuadE(x, y int) {
	if x < 1 || y < 1 {
		return
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			switch {
			case i == 0 && j == 0:
				print("A")
			case i == y-1 && j == 0:
				print("C")
			case i == 0 && j == x-1:
				print("C")
			case i == y-1 && j == x-1:
				print("A")
			case i == 0 || j == 0 || i == y-1 || j == x-1:
				print("B")
			default:
				print(" ")
			}
		}
		print("\n")
	}
}
