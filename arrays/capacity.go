package arrays

import "fmt"

func Capacity() {
	s := []string{}
	lastCap := cap(s)

	for i := 0; i < 1000; i++ {
		s = append(s, "Freaky")
		if cap(s) != lastCap {
			fmt.Println(cap(s))
			lastCap = cap(s)
		}
	}
}
