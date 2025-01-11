package pointers

import "fmt"

func reclassify(planets *[]string) {
	*planets = (*planets)[0:8]
}

func MemSlice() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
		"Pluto",
	}
	reclassify(&planets)
	fmt.Println(planets)
}
