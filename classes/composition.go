package classes

import "fmt"

type sol int

// this struct is composed of multiple structs
type report struct {
	sol
	temperature
	// the location type is stored in coordinate.go
	// it has: lat, long float64
	location
}

type temperature struct {
	high, low celsius
}
type celsius float64

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}
func (r report) average() celsius {
	return r.temperature.average()
}

func Comp() {
	// bradbury := location{-4.5895, 137.4417}
	// t := temperature{high: -1.0, low: -78.0}
	report := report{
		sol:         15,
		location:    location{"a", -4.5895, 137.4417},
		temperature: temperature{high: -1.0, low: -78.0},
	}
	fmt.Println(report.sol.days(1446))
	fmt.Println(report.days(1446))
	fmt.Printf("average %v celsius\n", report.average())

	fmt.Printf("%+v\n", report)
	fmt.Printf("a balmy %v celsius\n", report.temperature.high)
}
