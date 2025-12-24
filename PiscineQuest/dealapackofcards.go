package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	card := 0

	for i := 1; i <= 4; i++ {
		fmt.Printf("Player %d: ", i)

		for j := 0; j < 3; j++ {
			if j > 0 {
				fmt.Printf(", ")
			}
			fmt.Printf("%d", deck[card])
			card++
		}
		fmt.Printf("\n")
	}
}
