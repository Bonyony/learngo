package learnerrors

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Files() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
