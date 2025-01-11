package pointers

import "fmt"

func Memory() {
	answer := 42
	// points to the memory address, prints a hexcode
	fmt.Println(&answer)

	address := &answer
	// dereferences the pointer and will print 42
	fmt.Println(*address)
	fmt.Printf("address is a %T\n", address) // *int

	america := "America!"
	var home *string
	fmt.Printf("home is a %T\n", home) // *string
	home = &america
	fmt.Println(*home) // America!

	var administrator *string
	scolese := "Christopher J. Scolese"
	administrator = &scolese
	fmt.Println(*administrator)

	bolden := "Charles F. Bolden"
	administrator = &bolden
	fmt.Println(*administrator)
}
