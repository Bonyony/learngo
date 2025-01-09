package interfaces

import (
	"encoding/json"
	"fmt"
	"os"
)

// the coordinate type is in stringer.go
// the String() method for coordinate types is too

func (c coordinate) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		DMS string  `json:"dms"`
		D   float64 `json:"degrees"`
		M   float64 `json:"minutes"`
		S   float64 `json:"seconds"`
		H   string  `json:"hemisphere"`
	}{
		DMS: c.String(),
		D:   c.d,
		M:   c.m,
		S:   c.s,
		H:   string(c.h),
	})
}

func Marshaler() {
	j := coordinate{135, 54, 0.0, 'E'}

	b, err := json.MarshalIndent(j, "", "	")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Println(string(b))
}
