package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main_2b() {
	var total = 0
	fmt.Println("hello world")

	bytes, err := os.ReadFile("day2-input-a.txt")
	check(err)

	var text = string(bytes)

	var lines = strings.Split(text, "\r\n")
	/*
		var lines = [5]string{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"}
	*/
	for l := 0; l < len(lines); l++ {
		var line = lines[l]
		fmt.Print(l + 1)
		fmt.Print(" " + line + "\n")

		var value = processLinePower(line)
		fmt.Print(value)
		fmt.Print("\n")

		total += value
	}

	fmt.Print("\n")
	fmt.Print(total)
}

func processLinePower(line string) int {

	// "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
	var titleAndPlays = strings.Split(line, ": ")
	// "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
	var plays = titleAndPlays[1]

	var pullsText = strings.Split(plays, "; ")

	var mostRed = 0
	var mostBlue = 0
	var mostGreen = 0

	// 3 blue, 4 red
	// 1 red, 2 green, 6 blue
	// 2 green
	for p := 0; p < len(pullsText); p++ {

		// 1 red, 2 green, 6 blue
		var pullText = pullsText[p]
		var colors = strings.Split(pullText, ", ")

		// 1 red
		// 2 green
		// 6 blue
		for c := 0; c < len(colors); c++ {
			// 1 red
			var colorPair = colors[c]
			// "1"
			// red
			var parts = strings.Split(colorPair, " ")
			var numberText = parts[0]
			var colorText = parts[1]

			number, err := strconv.Atoi(numberText)
			check(err)

			if colorText == "blue" {
				if number > mostBlue {
					mostBlue = number
				}
			} else if colorText == "green" {
				if number > mostGreen {
					mostGreen = number
				}
			} else if colorText == "red" {
				if number > mostRed {
					mostRed = number
				}
			}
		}
	}

	return mostBlue * mostGreen * mostRed
}
