package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hello world")

	bytes, err := os.ReadFile("day6-input.txt")
	check(err)

	var text = string(bytes)

	var lines = strings.Split(text, "\r\n")

	var possibilities int

	var targetTime = readNumber(lines[0])
	var targetDistance = readNumber(lines[1])

	for t := 1; t < targetTime; t++ {
		speed := t
		distance := (targetTime - t) * speed
		if distance > targetDistance {
			possibilities++
		}
	}

	fmt.Printf("Total: %d", possibilities)
}

func readNumber(input string) int {

	// foo: 1   2  3   4   5 -> 12345
	var parts = strings.Split(strings.Split(input, ":")[1], " ")
	parts = slices.DeleteFunc(parts, func(n string) bool {
		return n == ""
	})
	var numberText = strings.Join(parts, "")
	number, err := strconv.Atoi(numberText)
	check(err)
	return number
}
