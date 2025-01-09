package arrays

import "fmt"

type Planets []string

func (p Planets) format() {
	for i := range p {
		p[i] = "New " + p[i]
	}
}

func Terraform() {
	planets := []string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	Planets(planets[3:4]).format()
	Planets(planets[6:]).format()

	fmt.Print(planets)
}
