// package piscine

// import "github.com/01-edu/z01"

// func QuadA1(x, y int) {
// 	if x >= 1 && y >= 1 {
// 		for i := 1; i <= y; i++ {
// 			if i == 1 || i == y {
// 				for j := 1; j <= x; j++ {
// 					if j == 1 || j == x {
// 						z01.PrintRune('o')
// 					} else {
// 						z01.PrintRune('-')
// 					}
// 				}
// 			} else {
// 				for j := 1; j <= x; j++ {
// 					if j == 1 || j == x {
// 						z01.PrintRune('|')
// 					} else {
// 						z01.PrintRune(' ')
// 					}
// 				}
// 			}
// 			print("\n")
// 		}
// 	}
// }

// func QuadB1(x, y int) {
// 	if x >= 1 && y >= 1 {
// 		for i := 1; i <= y; i++ {
// 			if i == 1 {
// 				for j := 1; j <= x; j++ {
// 					if j == 1 {
// 						z01.PrintRune('/')
// 					} else if j == x {
// 						z01.PrintRune('\\')
// 					} else {
// 						z01.PrintRune('*')
// 					}
// 				}
// 			} else if i == y {
// 				for j := 1; j <= x; j++ {
// 					if j == 1 {
// 						z01.PrintRune('\\')
// 					} else if j == x {
// 						z01.PrintRune('/')
// 					} else {
// 						z01.PrintRune('*')
// 					}
// 				}
// 			} else {
// 				for j := 1; j <= x; j++ {
// 					if j == 1 || j == x {
// 						z01.PrintRune('*')
// 					} else {
// 						z01.PrintRune(' ')
// 					}
// 				}
// 			}
// 			z01.PrintRune('\n')
// 		}
// 	}
// }

// func QuadC1(x, y int) {
// 	if x >= 1 && y >= 1 {
// 		for i := 1; i <= y; i++ {
// 			if i == 1 {
// 				for j := 1; j <= x; j++ {
// 					if j == 1 {
// 						z01.PrintRune('A')
// 					} else if j == x {
// 						z01.PrintRune('A')
// 					} else {
// 						z01.PrintRune('B')
// 					}
// 				}
// 			} else if i == y {
// 				for j := 1; j <= x; j++ {
// 					if j == 1 {
// 						z01.PrintRune('C')
// 					} else if j == x {
// 						z01.PrintRune('C')
// 					} else {
// 						z01.PrintRune('B')
// 					}
// 				}
// 			} else {
// 				for j := 1; j <= x; j++ {
// 					if j == 1 || j == x {
// 						z01.PrintRune('B')
// 					} else {
// 						z01.PrintRune(' ')
// 					}
// 				}
// 			}
// 			z01.PrintRune('\n')
// 		}
// 	}
// }

// func QuadD1(x, y int) {
// 	if x >= 1 && y >= 1 {
// 		for i := 1; i <= y; i++ {
// 			if i == 1 {
// 				for j := 1; j <= x; j++ {
// 					if j == 1 {
// 						z01.PrintRune('A')
// 					} else if j == x {
// 						z01.PrintRune('C')
// 					} else {
// 						z01.PrintRune('B')
// 					}
// 				}
// 			} else if i == y {
// 				for j := 1; j <= x; j++ {
// 					if j == 1 {
// 						z01.PrintRune('A')
// 					} else if j == x {
// 						z01.PrintRune('C')
// 					} else {
// 						z01.PrintRune('B')
// 					}
// 				}
// 			} else {
// 				for j := 1; j <= x; j++ {
// 					if j == 1 || j == x {
// 						z01.PrintRune('B')
// 					} else {
// 						z01.PrintRune(' ')
// 					}
// 				}
// 			}
// 			z01.PrintRune('\n')
// 		}
// 	}
// }

// func QuadE1(x, y int) {
// 	if x >= 1 && y >= 1 {
// 		for i := 1; i <= y; i++ {
// 			if i == 1 {
// 				for j := 1; j <= x; j++ {
// 					if j == 1 {
// 						z01.PrintRune('A')
// 					} else if j == x {
// 						z01.PrintRune('C')
// 					} else {
// 						z01.PrintRune('B')
// 					}
// 				}
// 			} else if i == y {
// 				for j := 1; j <= x; j++ {
// 					if j == 1 {
// 						z01.PrintRune('C')
// 					} else if j == x {
// 						z01.PrintRune('A')
// 					} else {
// 						z01.PrintRune('B')
// 					}
// 				}
// 			} else {
// 				for j := 1; j <= x; j++ {
// 					if j == 1 || j == x {
// 						z01.PrintRune('B')
// 					} else {
// 						z01.PrintRune(' ')
// 					}
// 				}
// 			}
// 			z01.PrintRune('\n')
// 		}
// 	}
// }
