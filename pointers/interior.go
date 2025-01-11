package pointers

import "fmt"

type stats struct {
	level             int
	endurance, health int
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

type character struct {
	name  string
	stats stats
}

func Interior() {
	player := character{name: "Franklin"}
	levelUp(&player.stats)
	fmt.Printf("%+v\n", player)
}
