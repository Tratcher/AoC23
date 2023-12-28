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
	// The values on the initial seeds: line come in pairs. Within each pair, the first value is the start of the range and the second value is the length of the range.
	var seedline = lines[0]
	var seedText = strings.Split(strings.Split(seedline, ": ")[1], " ")
	var mappedValues = make([][2]int, len(seedText)/2)

	for i := 0; i < len(mappedValues); i++ {
		number, err := strconv.Atoi(seedText[i*2])
		check(err)
		mappedValues[i][0] = number
		number, err = strconv.Atoi(seedText[i*2+1])
		check(err)
		mappedValues[i][1] = number
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

		mappedValues = applyMapb(mappedValues, maps)
	}

	var lowest = mappedValues[0][0]
	for i := 1; i < len(mappedValues); i++ {
		if mappedValues[i][0] < lowest {
			lowest = mappedValues[i][0]
		}
	}

	fmt.Printf("Lowest: %d", lowest)
}

func applyMapb(mappedValues [][2]int, maps [][]int) [][2]int {
	for v := 0; v < len(mappedValues); v++ {
		value := mappedValues[v][0]
		valueRange := mappedValues[v][1]

		for m := 0; m < len(maps); m++ {
			destination := maps[m][0]
			source := maps[m][1]
			length := maps[m][2]

			if source <= value && value < source+length {
				mappedValues[v][0] = value - source + destination

				if valueRange > length {
					diff := valueRange - length
					mappedValues = append(mappedValues, [2]int{source + length, diff})
					mappedValues[v][1] = length
				}

				break
			}
		}
	}

	fmt.Print(mappedValues)
	fmt.Println()

	return mappedValues
}
