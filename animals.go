package main

import (
	"fmt"
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
	eat() string
	move() string
}

func step(a animal) {
	switch rand.Intn(2) {
	case 0:
		fmt.Printf("%v %v.\n", a, a.move())
	default:
		fmt.Printf("%v eats %v.\n", a, a.eat())
	}
}

func Animals() {
	animals := []animal{
		wumpkin{name: "Wumpus"},
		flumphkin{name: "Flumphus"},
		grumpkin{name: "Grumpus"},
	}
	count := 1
	for i := 0; i < 3; i++ {
		for j := 0; j < 24; j++ {
			if j < 6 || j > 20 {
				fmt.Printf("The chumplings are sleeping during day %v\n", count)
			} else {
				a := rand.Intn(len(animals))
				step(animals[a])
			}
		}
		count += 1
	}
}
