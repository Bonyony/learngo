package main

import (
	"math/rand"
)

// Grumps
type grumpkin struct {
	name string
}

func (g grumpkin) String() string {
	return g.name
}
func (g grumpkin) eat() string {
	switch rand.Intn(3) {
	case 0:
		return "pumpkins and mash"
	case 1:
		return "cheezburgers"
	default:
		return "human flesh"
	}
}
func (g grumpkin) move() string {
	return "The grumpkin scurries around"
}

// Flumphs
type flumphkin struct {
	name string
}

func (f flumphkin) String() string {
	return f.name
}
func (f flumphkin) eat() string {
	switch rand.Intn(3) {
	case 0:
		return "feathers"
	case 1:
		return "big bubbles"
	default:
		return "human flesh"
	}
}
func (f flumphkin) move() string {
	return "The flumphkin flumphs @ U"
}

// Wumps
type wumpkin struct {
	name string
}

func (w wumpkin) String() string {
	return w.name
}
func (w wumpkin) eat() string {
	switch rand.Intn(3) {
	case 0:
		return "elephant sugar"
	case 1:
		return "deer hooves"
	default:
		return "human flesh"
	}
}
func (w wumpkin) move() string {
	return "The wumpkin wumps around a bit"
}

type animal interface {
	eat()
	move()
}

func Animals() {

}
