package structs

import "fmt"

type location struct {
	name string
	lat  float64
	long float64
}

func Struct() {
	curiosity := location{"jon", -4.123, 135.1358}
	locations := []location{
		{name: "Bradbury Landing", lat: -4.5895, long: 137.4417},
		{name: "Columbia Memorial Station", lat: -14.5684, long: 175.472636},
		{name: "Challenger Memorial Station", lat: -1.9462, long: 354.4734},
	}

	fmt.Printf("%v\n", curiosity)
	// using the + sign adds field names to the output
	fmt.Printf("%+v\n", curiosity)
	fmt.Println(locations[1])
}
