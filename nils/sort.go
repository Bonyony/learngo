package nils

import (
	"fmt"
	"sort"
)

func sortStrings(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i, j int) bool {
			return s[i] < s[j]
		}
	}
	sort.Slice(s, less)
}

func MainSort() {
	food := []string{"onion", "carrot", "celery"}
	sortStrings(food, nil)
	fmt.Println(food)

	// slices
	var soup []string
	fmt.Println(soup == nil)
	for _, ingredient := range soup {
		fmt.Println(ingredient)
	}
	fmt.Println(len(soup))
	soup = append(soup, "onion", "celery", "carrot")
	fmt.Println(soup)
}
