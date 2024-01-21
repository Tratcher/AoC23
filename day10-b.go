package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	Inside   = 1
	Outside  = 2
	Boundary = 3
)

func main_10b() {
	fmt.Println("hello world")

	bytes, err := os.ReadFile("day10-input-test-b3.txt")
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

	insideOrOut := make([][]int, len(lines))
	for i := 0; i < len(lines); i++ {
		insideOrOut[i] = make([]int, len(lines[i]))
	}
	insideOrOut[startR][startC] = Boundary

	// Find the loop
	for {
		current := lines[r][c]
		// fmt.Printf("%c %v\n", current, enteredFrom)
		if current == Start {
			break
		}

		insideOrOut[r][c] = Boundary

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

	// Figure out how many spaces are inside the loop and how many are outside
	changed := false
	for {
		for r = 0; r < len(insideOrOut); r++ {
			for c = 0; c < len(insideOrOut[r]); c++ {
				// Already known
				if insideOrOut[r][c] != 0 {
					continue
				}

				// Edges must be outside or boundary
				if r == 0 || c == 0 || r == len(insideOrOut)-1 || c == len(insideOrOut[r])-1 {
					insideOrOut[r][c] = Outside
					changed = true
					continue
				}

				// If we're adjacent to any Inside or Outside space then we must be the same
				if r > 0 && (insideOrOut[r-1][c] == Inside || insideOrOut[r-1][c] == Outside) {
					insideOrOut[r][c] = insideOrOut[r-1][c]
					changed = true
					continue
				}
				if r < len(insideOrOut)-1 && (insideOrOut[r+1][c] == Inside || insideOrOut[r+1][c] == Outside) {
					insideOrOut[r][c] = insideOrOut[r+1][c]
					changed = true
					continue
				}
				if c > 0 && (insideOrOut[r][c-1] == Inside || insideOrOut[r][c-1] == Outside) {
					insideOrOut[r][c] = insideOrOut[r][c-1]
					changed = true
					continue
				}
				if c < len(insideOrOut[r])-1 && (insideOrOut[r][c+1] == Inside || insideOrOut[r][c+1] == Outside) {
					insideOrOut[r][c] = insideOrOut[r][c+1]
					changed = true
					continue
				}

				// Count the number of boundary spaces between here and the end of the row. Odd means we're inside, even means we're outside
				boundaries := 0
				southOnly := 0
				northOnly := 0
				for b := c + 1; b < len(insideOrOut[r]); b++ {
					// Don't count east-west lines, they don't affect parity
					if insideOrOut[r][b] == Boundary {
						switch lines[r][b] {
						case NorthToSouth:
							boundaries++
						case NorthToEast:
						case NorthToWest:
							northOnly++
						case SouthToEast:
						case SouthToWest:
							southOnly++
						}
					}
				}

				// Cross over
				boundaries += (southOnly + northOnly) / 2

				if boundaries%2 == 0 {
					insideOrOut[r][c] = Outside
				} else {
					insideOrOut[r][c] = Inside
				}
				changed = true
			}
		}

		if !changed {
			break
		}
		changed = false
	}

	for i := 0; i < len(insideOrOut); i++ {
		fmt.Println(lines[i])
	}
	for i := 0; i < len(insideOrOut); i++ {
		fmt.Println(insideOrOut[i])
	}

	total := 0
	for r = 0; r < len(insideOrOut); r++ {
		for c = 0; c < len(insideOrOut[r]); c++ {
			if insideOrOut[r][c] == Inside {
				total++
			}
		}
	}

	fmt.Println(total)
}
