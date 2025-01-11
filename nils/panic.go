package nils

import "fmt"

type person struct {
	age int
}

func (p *person) birthday() {
	// guard against a nil pointer
	if p == nil {
		return
	}
	p.age++
}

func PanicMain() {
	var nobody *person
	fmt.Println(nobody)
	nobody.birthday()

	var fn func(a, b int) int
	fmt.Println(fn == nil) // prints true
}
