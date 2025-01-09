package arrays

import (
	"fmt"
	"sort"
)

func Slices() {
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
	terrestrial := planets[:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:]
	fmt.Println(terrestrial, gasGiants, iceGiants)
	fmt.Println(gasGiants[0])
	// slice into an array and then slice into the slices!
	giants := planets[4:8]
	gas := giants[0:2]
	ice := giants[2:4]
	fmt.Println(giants, gas, ice)
	// slice into strings
	neptune := "Neptune"
	// slice off the the first 3 letters, or rather only view from index 3 onwards
	// remember indices are the number of bytes, not runes!
	tune := neptune[3:]
	fmt.Println(tune)

	// start with a slice
	dwarfArray := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Printf("dwarfs is %T, dwarfArray is %T\n", dwarfs, dwarfArray)
	// sort things using the sort package
	sort.Strings(planets)
	fmt.Println(planets)
}
