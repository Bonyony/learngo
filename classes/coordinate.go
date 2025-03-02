package classes

import "fmt"

// coordinate in degrees, minutes, seconds in a N/S/E/W hemisphere.
type coordinate struct {
	d, m, s float64
	h       rune
}

// type location struct {
// 	lat, long float64
// }

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

// constructor function, conventions is as such: newType or newLocation
func newLocation(n string, lat, long coordinate) location {
	return location{"uou", lat.decimal(), long.decimal()}
}

func Coords() {
	// Bradbury Landing: 4º35'22.2" S, 137º26'30.1" E
	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}
	// one way to do it
	// curiosity := location{lat.decimal(), long.decimal()}

	// another way
	curiosity := newLocation("as", lat, long)
	fmt.Println(lat.decimal(), long.decimal())
	fmt.Println(curiosity)
}
