package interfaces

import (
	"fmt"
	"strings"
)

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

type rover struct{}

func (r rover) talk() string {
	return "whir whir"
}

// the interface neccessitates the talk() method
type talker interface {
	talk() string
}

type starship struct {
	laser
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func MainTalk() {
	shout(martian{})
	shout(laser(2))
	shout(rover{})

	// this block takes advantage of struct embedding
	// when the starship 'talks', the laser does the talking
	s := starship{laser(3)}
	fmt.Println(s.talk())
	shout(s)
}
