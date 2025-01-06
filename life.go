package main

import (
	"fmt"
	"math/rand"
	"time"
)

// this was a fun learning exercise.
// Life is the exporting function in this file!

const (
	width  = 50
	height = 15
)

// declares a 2-d field of cells
type Universe [][]bool

// returns an empty universe
func NewUniverse() Universe {
	u := make(Universe, height)
	for i := range u {
		u[i] = make([]bool, width)
	}
	return u
}

// populates the board randomly
func (u Universe) Seed() {
	for i := 0; i < (width * height / 10); i++ {
		u.Set(rand.Intn(width), rand.Intn(height), true)
	}
}

// sets the state of a specific cell
func (u Universe) Set(x, y int, b bool) {
	u[y][x] = b
}

// turns the Universe into a string
func (u Universe) String() string {
	var b byte
	/* buf is an array (slice so it is dynamic) of bytes that starts with no length,
	but has the capacity of (width+1)*height */
	buf := make([]byte, 0, (width+1)*height)
	//  nested for loop scans the entire 2d grid. if u[y][x] returns trues, b is populated
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b = ' '
			if u[y][x] {
				b = '*'
			}
			buf = append(buf, b)
		}
		buf = append(buf, '\n')
	}
	return string(buf)
}

func (u Universe) Show() {
	fmt.Print("\x0c", u.String())
}

// returns true if the specified cell is alive
func (u Universe) Alive(x, y int) bool {
	// using the mod operator wraps around the universe so there are not out of bounds checks
	x = (x + width) % width
	y = (y + height) % height
	return u[y][x]
}

// counts how many adjacent cells are alive
func (u Universe) Neighbors(x, y int) int {
	// count variable
	n := 0
	// double nested loops go through each possible cell and increment n if it is alive
	for v := -1; v <= 1; v++ {
		for h := -1; h <= 1; h++ {
			if !(v == 0 && h == 0) && u.Alive(x+h, y+h) {
				n++
			}
		}
	}
	return n
}

func (u Universe) Next(x, y int) bool {
	n := u.Neighbors(x, y)
	return n == 3 || n == 2 && u.Alive(x, y)
}

func Step(a, b Universe) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b.Set(x, y, a.Next(x, y))
		}
	}
}

func Life() {
	a, b := NewUniverse(), NewUniverse()
	a.Seed()
	for i := 0; i < 300; i++ {
		Step(a, b)
		a.Show()
		time.Sleep(time.Second)
		// this swaps the universes
		a, b = b, a
	}
}
