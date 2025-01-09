package maps

import (
	"fmt"
	"math"
)

func Freq() {
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	frequency := make(map[float64]int)
	groups := make(map[float64][]float64)
	// iterates over a slice (index, value)
	for _, t := range temperatures {
		frequency[t]++

		g := math.Trunc(t/10) * 10
		groups[g] = append(groups[g], t)
	}
	// iterates over a map (key, value)
	for t, num := range frequency {
		fmt.Printf("%+.2f occurs %d times\n", t, num)
	}
	for g, temperatures := range groups {
		fmt.Printf("%v: %v\n", g, temperatures)
	}
}
