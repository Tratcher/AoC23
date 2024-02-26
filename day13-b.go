package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("hello world")

	bytes, err := os.ReadFile("day13-input.txt")
	check(err)

	text := string(bytes)

	patterns := strings.Split(text, "\r\n\r\n")

	total := 0

	for p := 0; p < len(patterns); p++ {
		pattern := patterns[p]

		lines := strings.Split(pattern, "\r\n")

		reflected := false
		// Check columns
		for c := 0; c < len(lines[0])-1; c++ {
			// Find mirrored columns
			smudges := 0
			match := true
			for offset := 0; offset <= c && offset+c+1 < len(lines[0]) && match; offset++ {
				smudges += CheckColumnsMatchB(lines, c-offset, c+offset+1)
				if smudges > 1 {
					match = false
				}
			}
			if match && smudges == 1 {
				// add up the number of columns to the left of each vertical line of reflection;
				total += c + 1
				reflected = true
				break
			}
		}

		if reflected {
			continue
		}

		// Check rows
		for r := 0; r < len(lines)-1; r++ {
			// Find mirrored columns
			smudges := 0
			match := true
			for offset := 0; offset <= r && offset+r+1 < len(lines) && match; offset++ {
				smudges += CheckRowsMatchB(lines, r-offset, r+offset+1)
				if smudges > 1 {
					match = false
				}
			}
			if match && smudges == 1 {
				// add 100 multiplied by the number of rows above each horizontal line of reflection
				total += (r + 1) * 100
				reflected = true
				break
			}
		}

		if !reflected {
			panic("Reflection not found in pattern: \n" + pattern)
		}

	}

	fmt.Println(total)
}

func CheckColumnsMatchB(lines []string, columnA int, columnB int) int {
	mismatches := 0
	for r := 0; r < len(lines); r++ {
		line := lines[r]
		if line[columnA] != line[columnB] {
			mismatches++
		}
	}
	return mismatches
}

func CheckRowsMatchB(lines []string, rowA int, rowB int) int {
	mismatches := 0
	lineA := lines[rowA]
	lineB := lines[rowB]
	for c := 0; c < len(lineA); c++ {
		if lineA[c] != lineB[c] {
			mismatches++
		}
	}
	return mismatches
}
