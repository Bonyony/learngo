package arrays

import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func makeWorld(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))

	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}

func Append() {
	// original array
	dwarfs1 := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	// these are treated as copies
	dwarfs2 := append(dwarfs1, "Orcus")
	dwarfs3 := append(dwarfs2, "Salacia", "Quaoar", "Sedna")
	dump("dwarfs1", dwarfs1)
	dump("dwarfs1[1:2]", dwarfs1[1:2])
	dwarfs3[1] = "Joo-joo"
	// only 2 and 3 are mutated by the above line.
	// dwarfs1 points to a different array in memory!
	fmt.Println(dwarfs1, dwarfs2, dwarfs3)

	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	// triple slicing declears the capacity of the slice
	terrestrial := planets[0:4:4]
	worlds := append(terrestrial, "Ceres")
	fmt.Println(planets)

	terrestrial = planets[0:4]
	worlds = append(terrestrial, "Ceres")
	fmt.Println(planets, worlds)

	dwarfs := make([]string, 0, 10)
	dwarfs = append(dwarfs, "Cerres", "Pluto", "Haumea", "Makemake", "Eris")
	fmt.Println(dwarfs)

	twoWorlds := makeWorld("New", "Venus", "Mars")
	fmt.Println(twoWorlds)
}
