package piscine

func QuadA(x, y int) {
	if x >= 1 && y >= 1 {
		for i := 1; i <= y; i++ {
			if i == 1 {
				for j := 1; j <= x; j++ {
					if j == 1 {
						print("/")
					} else if j == x {
						print("\\")
					} else {
						print("*")
					}
				}
			} else if i == y {
				print("\\")
			} else {
				for j := 1; j <= x; j++ {
					if j == 1 || j == x {
						print("*")
					} else {
						print(" ")
					}
				}
			}
			print("\n")
		}
	} else {
		print("")
	}
}
