package main

import (
	"fmt"
	"os"
	"strings"
)

type Direction int

const (
	North Direction = 1
	East  Direction = 2
	South Direction = 3
	West  Direction = 4
)

const (
	NorthToSouth byte = '|'
	EastToWest   byte = '-'
	NorthToEast  byte = 'L'
	NorthToWest  byte = 'J'
	SouthToWest  byte = '7'
	SouthToEast  byte = 'F'
	Ground       byte = '.'
	Start        byte = 'S'
)

func main() {
	fmt.Println("hello world")

	bytes, err := os.ReadFile("day10-input.txt")
	check(err)

	text := string(bytes)

	lines := strings.Split(text, "\r\n")

	startR := 0
	startC := 0

	// Find the start S
	for r := 0; r < len(lines); r++ {
		for c := 0; c < len(lines[r]); c++ {
			if lines[r][c] == Start {
				startR = r
				startC = c
				break
			}
		}
	}

	var enteredFrom Direction
	r := 0
	c := 0
	// Find a connecting pipe, there should be two.
	if startR > 0 {
		if lines[startR-1][startC] == NorthToSouth ||
			lines[startR-1][startC] == SouthToWest ||
			lines[startR-1][startC] == SouthToEast {
			enteredFrom = South
			r = startR - 1
			c = startC
		}
	}
	if startR < len(lines)-1 {
		if lines[startR+1][startC] == NorthToSouth ||
			lines[startR+1][startC] == NorthToWest ||
			lines[startR+1][startC] == NorthToEast {
			enteredFrom = North
			r = startR + 1
			c = startC
		}
	}
	if startC > 0 {
		if lines[startR][startC-1] == EastToWest ||
			lines[startR][startC-1] == NorthToEast ||
			lines[startR][startC-1] == SouthToEast {
			enteredFrom = East
			r = startR
			c = startC - 1
		}
	}
	if startC < len(lines[startR])-1 {
		if lines[startR][startC+1] == EastToWest ||
			lines[startR][startC+1] == NorthToWest ||
			lines[startR][startC+1] == SouthToWest {
			enteredFrom = West
			r = startR
			c = startC + 1
		}
	}

	// Find the length of the loop
	total := 0
	for {
		total++
		current := lines[r][c]
		fmt.Printf("%c %v\n", current, enteredFrom)
		if current == Start {
			break
		}

		switch enteredFrom {
		case North:
			switch current {
			case NorthToSouth:
				r++
			case NorthToEast:
				c++
				enteredFrom = West
			case NorthToWest:
				c--
				enteredFrom = East
			default:
				panic(current)
			}
		case South:
			switch current {
			case NorthToSouth:
				r--
			case SouthToEast:
				c++
				enteredFrom = West
			case SouthToWest:
				c--
				enteredFrom = East
			default:
				panic(current)
			}
		case West:
			switch current {
			case EastToWest:
				c++
			case NorthToWest:
				r--
				enteredFrom = South
			case SouthToWest:
				r++
				enteredFrom = North
			default:
				panic(current)
			}
		case East:
			switch current {
			case EastToWest:
				c--
			case NorthToEast:
				r--
				enteredFrom = South
			case SouthToEast:
				r++
				enteredFrom = North
			default:
				panic(current)
			}
		default:
			panic(enteredFrom)
		}
	}

	// Find the half-way distance
	fmt.Println(total / 2)
}
