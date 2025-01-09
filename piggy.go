package main

import (
	"fmt"
	"math/rand"
)

func Piggy() {
	fmt.Printf("\nLet's add some shit together.\n")

	nickel := 0.05
	dime := 0.10
	quarter := 0.25

	for total := 0.0; total < 20.00; {
		randCoin := rand.Intn(3) + 1
		switch randCoin {
		case 1:
			total += nickel
			fmt.Printf("Piggy has $%5.2f\n", total)
			fmt.Println("Added: ", nickel, " cents")
		case 2:
			total += dime
			fmt.Printf("Piggy has $%5.2f\n", total)
			fmt.Println("Added: ", dime, " cents")
		case 3:
			total += quarter
			fmt.Printf("Piggy has $%5.2f\n", total)
			fmt.Println("Added: ", quarter, " cents")
		}
	}
	fmt.Println("Yay! My piggy is full!")
}
