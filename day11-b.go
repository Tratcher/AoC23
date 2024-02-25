package main

import (
	"fmt"
	"os"
	"strings"
)

func main_11b() {
	fmt.Println("hello world")

	bytes, err := os.ReadFile("day11-input.txt")
	check(err)

	text := string(bytes)

	lines := strings.Split(text, "\r\n")

	fmt.Println(strings.Join(lines, "\n"))

	// Find the galaxies
	galaxies := make([][]int, 0)
	for r := 0; r < len(lines); r++ {
		line := lines[r]
		for c := 0; c < len(line); c++ {
			if line[c] == '#' {
				galaxy := []int{r, c}
				galaxies = append(galaxies, galaxy)
			}
		}
	}

	// Make a copy for moves
	galaxies2 := make([][]int, len(galaxies))
	for g := 0; g < len(galaxies); g++ {
		galaxies2[g] = make([]int, 2)
		galaxies2[g][0] = galaxies[g][0]
		galaxies2[g][1] = galaxies[g][1]
	}

	expansionFactor := 999999

	// Find empty rows, expand them
	for r := 0; r < len(lines); r++ {
		line := lines[r]

		if !strings.Contains(line, "#") {
			// For any galaxy who's row is after this, move it down more
			for g := 0; g < len(galaxies); g++ {
				if galaxies[g][0] > r {
					galaxies2[g][0] += expansionFactor
				}
			}
		}
	}

	// Find empty columns, expand them
	for c := 0; c < len(lines[0]); c++ {
		empty := true
		for r := 0; r < len(lines); r++ {
			if lines[r][c] != '.' {
				empty = false
				break
			}
		}
		if empty {
			// For any galaxy who's colum is after this, move it over more
			for g := 0; g < len(galaxies); g++ {
				if galaxies[g][1] > c {
					galaxies2[g][1] += expansionFactor
				}
			}
		}
	}

	fmt.Println(strings.Join(lines, "\n"))

	total := 0

	// Find the shortest path between all galaxies
	for g := 0; g < len(galaxies2); g++ {
		for o := g + 1; o < len(galaxies2); o++ {
			total += getDistance(galaxies2[g], galaxies2[o])
		}
	}

	fmt.Println(total)
}
