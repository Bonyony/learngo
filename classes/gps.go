package classes

import (
	"fmt"
	"math"
)

type gps struct {
	current     location
	destination location
	world
}
type location struct {
	name      string
	lat, long float64
}

func (l location) description() string {
	return fmt.Sprintf("%v (%.1f, %.1f)", l.name, l.lat, l.long)
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func GPS() {
	type rover struct {
		gps
	}

}
