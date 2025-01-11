package pointers

import "fmt"

type person struct {
	name, superpower string
	age              int
}

// functional
func birthday(p *person) {
	p.age++
}

// reciever or method
func (p *person) birthday() {
	p.age++
}

func MemStruct() {
	timmy := &person{
		name: "Timothy",
		age:  10,
	}
	timmy.superpower = "flying"
	birthday(timmy)
	fmt.Printf("%+v\n", timmy) // prints a pointer to these values &{name: ...}
	superpowers := &[3]string{"flight", "invisibility", "super strength"}
	// when an array is indexed or slided it will automatically be dereferenced
	fmt.Println(superpowers[0])
	fmt.Println(superpowers[1:2])

	rebecca := person{
		name:       "Rebecca",
		superpower: "imagination",
		age:        14,
	}
	birthday(&rebecca)
	fmt.Printf("%+v\n", rebecca)

	terry := &person{
		name: "Terry",
		age:  15,
	}
	terry.birthday()
	fmt.Printf("%+v\n", terry)
}
