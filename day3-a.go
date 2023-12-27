package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main_3a() {
	var total = 0
	fmt.Println("hello world")

	var digits = "0123456789"

	bytes, err := os.ReadFile("day3-input.txt")
	check(err)

	var text = string(bytes)

	var lines = strings.Split(text, "\r\n")

	for r := 0; r < len(lines); r++ {
		var numberStart = -1
		var numberEnd = -1

		for c := 0; c < len(lines[r]); c++ {
			var ch = lines[r][c : c+1]
			// Scan for a digit
			if strings.Contains(digits, ch) {
				// Start of a number
				if numberStart == -1 {
					numberStart = c
				}

				numberEnd = c

				// End of a number
			} else if numberEnd != -1 {
				total += handleNumber(lines, r, numberStart, numberEnd)
				// Reset
				numberStart = -1
				numberEnd = -1
			}
		}

		// Number went to the end of the line
		if numberEnd != -1 {
			total += handleNumber(lines, r, numberStart, numberEnd)
		}
	}

	fmt.Print("\n")
	fmt.Print(total)
}

func handleNumber(lines []string, row int, start int, end int) int {

	// Search around this number, see if it's adjacent to a symbol
	var value = lines[row][start : end+1]
	fmt.Printf("r: %d, %s\n", row, value)

	if checkForSymbols(lines, row, start, end) {
		number, err := strconv.Atoi(value)
		check(err)

		fmt.Printf("+%s\n", value)

		return number
	}

	fmt.Printf("-%s\n", value)

	return 0
}

// 467.
// ...*
func checkForSymbols(lines []string, row int, start int, end int) bool {

	if start > 0 {
		start--
	}

	if end < len(lines[row])-1 {
		end++
	}

	if row > 0 {
		if containsSymbols(lines[row-1][start : end+1]) {
			return true
		}
	}

	if containsSymbols(lines[row][start : end+1]) {
		return true
	}

	if row < len(lines)-1 {
		if containsSymbols(lines[row+1][start : end+1]) {
			return true
		}
	}

	return false
}

func containsSymbols(segment string) bool {

	var digitsAndDot = "0123456789."

	for i := 0; i < len(segment); i++ {
		var ch = segment[i : i+1]

		if !strings.Contains(digitsAndDot, ch) {
			return true
		}
	}

	return false
}
