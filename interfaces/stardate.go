package interfaces

import (
	"fmt"
	"time"
)

type stardater interface {
	YearDay() int
	Hour() int
}

type sol int

// because we have implemented these methods for sol types,
// they can be used with the same methods on the time pkg
func (s sol) YearDay() int {
	return int(s % 668)
}
func (s sol) Hour() int {
	return 0
}

// returns a fictional measure of time for a given date
func stardate(t stardater) float64 {
	doy := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + doy + h
}

func MainStardate() {
	day := time.Date(2012, 8, 6, 5, 17, 0, 0, time.UTC)
	fmt.Printf("%.1f Curiosity has landed\n", stardate(day))

	s := sol(1422)
	fmt.Printf("%.1f Happy birthday\n", stardate(s))
}
