package maps

import "fmt"

func Map() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
		"Moon":  0,
	}
	temp := temperature["Earth"]
	fmt.Printf("On average the earth is %vº C.\n", temp)

	temperature["Earth"] = 16
	temperature["Venus"] = 464
	fmt.Println(temperature)
	moon := temperature["Moon"]
	fmt.Println(moon)
	if moon, ok := temperature["Moon"]; ok {
		fmt.Printf("On average the moon is %vº C.\n", moon)
	} else {
		fmt.Println("Where is the moon?")
	}

	// values in maps are not copied, so the original is mutated by other assignments
	// other variables that point to the map, use the same data
	planets := map[string]string{
		"Earth": "Sector ZZ9",
		"Mars":  "Sector ZZ9",
	}
	planetsMarkII := planets
	planets["Earth"] = "whoops"
	fmt.Println(planets)
	fmt.Println(planetsMarkII)
	delete(planets, "Earth")
	fmt.Println(planetsMarkII)
}
