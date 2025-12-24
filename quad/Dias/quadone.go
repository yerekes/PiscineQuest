package piscine

func QuadA(x, y int) {
	if x < 1 || y < 1 {
		return
	}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			switch {
			case i == 0 && j == 0:
				print("o")
			case i == 0 && j == x-1:
				print("o")
			case i == y-1 && j == 0:
				print("o")
			case i == y-1 && j == x-1:
				print("o")
			case i == 0 || i == y-1:
				print("-")
			case j == 0 || j == x-1:
				print("|")
			default:
				print(" ")
			}
		}
		print("\n")
	}
}
