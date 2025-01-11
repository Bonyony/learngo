package pointers

import "fmt"

type turtle struct {
	x, y int
}

// if you were to call the turtle type without dereferencing
// you will be warned with 'ineffective assignment'

func (t *turtle) up() {
	t.y++
}
func (t *turtle) down() {
	t.y--
}
func (t *turtle) left() {
	t.x--
}
func (t *turtle) right() {
	t.x++
}

func Turtle() {
	t := turtle{0, 0}
	t.up()
	t.right()
	t.right()
	fmt.Println(t)
	t.down()
	t.down()
	t.left()
	t.left()
	fmt.Println(t)
}
