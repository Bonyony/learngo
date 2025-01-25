package concurrency

import (
	"fmt"
	"strings"
)

func sourceGopher(downstream chan string) {
	// DuplicateGopher
	// for _, v := range []string{"a", "b", "b", "c", "d", "d", "d", "e"} {
	// 	downstream <- v
	// }

	// FilterGopher and SplitGopher
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downstream <- v
	}
	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	for item := range upstream {
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
	close(downstream)
}

// quick check 1
func duplicateGopher(upstream, downstream chan string) {
	holder := ""
	for item := range upstream {
		if item != holder {
			downstream <- item
			holder = item
		}

	}
	close(downstream)
}

// quick check 2
func splitGopher(upstream, downstream chan string) {
	for item := range upstream {
		for _, word := range strings.Fields(item) {
			downstream <- word
		}
	}
	close(downstream)
}

func printGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}

func PipeMain() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	// go filterGopher(c0, c1)
	// go duplicateGopher(c0, c1)
	go splitGopher(c0, c1)
	printGopher(c1)
}
