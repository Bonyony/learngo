package learnerrors

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const rows, columns = 9, 9

var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit  = errors.New("invalid digit")
)

type Grid [rows][columns]int8
type SodukoError []error

// Error returns one or more errors separated by commas.
func (se SodukoError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ", ")
}

func (g *Grid) Set(row, column int, digit int8) error {
	var errs SodukoError
	if !inBounds(row, column) {
		errs = append(errs, ErrBounds)
	}
	if !validDigit(digit) {
		errs = append(errs, ErrDigit)
	}
	if len(errs) > 0 {
		return errs
	}
	g[row][column] = digit
	return nil
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if column < 0 || column >= columns {
		return false
	}
	return true
}

func validDigit(digit int8) bool {
	if digit < 1 || digit > 9 {
		return false
	}
	return true
}

func MakeGrid() {
	var g Grid

	err := g.Set(15, 5, 10)
	if err != nil {
		if errs, ok := err.(SodukoError); ok {
			fmt.Printf("%d error(s) occurred:\n", len(errs))
			for _, e := range errs {
				fmt.Printf("- %v\n", e)
			}
		}
		os.Exit(1)
	}
}
