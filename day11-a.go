package main

import (
	"fmt"
	"os"
	"strings"
)

func main_11a() {
	fmt.Println("hello world")

	bytes, err := os.ReadFile("day11-input.txt")
	check(err)

	text := string(bytes)

	lines := strings.Split(text, "\r\n")

	fmt.Println(strings.Join(lines, "\n"))

	// Find empty rows, expand them
	for r := 0; r < len(lines); r++ {
		line := lines[r]

		if !strings.Contains(line, "#") {
			if r == len(lines)-1 {
				lines = append(lines, line)
			} else {
				lines = append(lines[:r+1], lines[r:]...)
				lines[r] = line
			}
			// Skip the inserted line
			r++
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
			for r := 0; r < len(lines); r++ {
				lines[r] = lines[r][:c] + "." + lines[r][c:]
			}
			// Skip the inserted column
			c++
		}
	}

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

	total := 0

	// Find the shortest path between all galaxies
	for g := 0; g < len(galaxies); g++ {
		for o := g + 1; o < len(galaxies); o++ {
			total += getDistance(galaxies[g], galaxies[o])
		}
	}

	fmt.Println(total)
}

// Manhatten distance
func getDistance(a []int, b []int) int {
	x := a[0] - b[0]
	y := a[1] - b[1]
	return absolute(x) + absolute(y)
}

func absolute(a int) int {
	if a > 0 {
		return a
	}
	return a * -1
}
