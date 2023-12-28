package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hello world")

	bytes, err := os.ReadFile("day5-input.txt")
	check(err)

	var text = string(bytes)

	var lines = strings.Split(text, "\r\n")

	// seeds: 79 14 55 13
	var seedline = lines[0]
	var seedText = strings.Split(strings.Split(seedline, ": ")[1], " ")
	var mappedValues = make([]int, len(seedText))

	for i := 0; i < len(mappedValues); i++ {
		number, err := strconv.Atoi(seedText[i])
		check(err)
		mappedValues[i] = number
	}

	// Skip seeds, blank, and next title
	lines = lines[3:]

	// There are 7 maps
	for m := 0; m < 7; m++ {
		var maps = readMap(lines)

		if m < 6 {
			// Skip the maps, a blank line, and next title
			lines = lines[len(maps)+2:]
		}

		applyMap(mappedValues, maps)
	}

	var lowest = mappedValues[0]
	for i := 1; i < len(mappedValues); i++ {
		if mappedValues[i] < lowest {
			lowest = mappedValues[i]
		}
	}

	fmt.Printf("Lowest: %d", lowest)
}

func readMap(lines []string) [][]int {
	var lineCount = 0

	// Scan until blank line
	for i := 0; i < len(lines) && lines[i] != ""; i++ {
		lineCount++
	}

	var maps = make([][]int, lineCount)

	// Read Map
	for i := 0; i < len(lines) && lines[i] != ""; i++ {
		var line = lines[i]
		var numberTexts = strings.Split(line, " ")
		var numbers = make([]int, len(numberTexts))
		for n := 0; n < len(numbers); n++ {
			value, err := strconv.Atoi(numberTexts[n])
			check(err)
			numbers[n] = value
		}

		maps[i] = numbers
	}

	return maps
}

func applyMap(mappedValues []int, maps [][]int) {
	for v := 0; v < len(mappedValues); v++ {
		value := mappedValues[v]

		for m := 0; m < len(maps); m++ {
			destination := maps[m][0]
			source := maps[m][1]
			length := maps[m][2]

			if source <= value && value < source+length {
				mappedValues[v] = value - source + destination
				break
			}
		}
	}

	fmt.Print(mappedValues)
	fmt.Println()
}
