package interfaces

import "fmt"

// a location with latitude and longitude in decimal degrees
type location struct {
	lat, long coordinate
}

// providing a String() method allows the fmt package to use the type
// String formats a location with latitude and longitude
func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) String() string {
	return fmt.Sprintf("%vº%v'%.1f\" %c", c.d, c.m, c.s, c.h)
}

func StrMain() {
	elysium := location{coordinate{4, 30, 0.0, 'N'}, coordinate{135, 54, 0.0, 'E'}}
	fmt.Println("Elysium Planitia is at", elysium)
}
