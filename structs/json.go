package structs

import (
	"encoding/json"
	"fmt"
	"os"
)

func Json() {
	// normal struct
	type location struct {
		// These values MUST be exported, hence the uppercase letters
		// the struct tags in `` alter the output of the Marshall function, for custom fields
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}

	curiosity := location{-4.48484, 123.234}
	// marshall the data into json
	bytes, err := json.Marshal(curiosity)
	// error handling
	exitOnError(err)
	// change the bytes to a string to make it readable
	fmt.Println(string(bytes))
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
