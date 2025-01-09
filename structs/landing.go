package structs

import (
	"fmt"

	"encoding/json"
)

func Landing() {
	type location struct {
		Name string
		Lat  float64
		Long float64
	}
	locations := []location{
		{Name: "Bradbury Landing", Lat: -4.5895, Long: 137.4417},
		{Name: "Columbia Memorial Station", Lat: -14.5684, Long: 175.472636},
		{Name: "Challenger Memorial Station", Lat: -1.9462, Long: 354.4734},
	}
	for i := range locations {
		bytes, err := json.MarshalIndent(locations[i], "", "	")
		exitOnError(err)
		fmt.Println(string(bytes))
	}
}
