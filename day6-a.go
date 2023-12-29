package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main_6a() {
	fmt.Println("hello world")

	bytes, err := os.ReadFile("day6-input.txt")
	check(err)

	var text = string(bytes)

	var lines = strings.Split(text, "\r\n")

	times := makeList(lines[0])
	distances := makeList(lines[1])

	total := 1

	for i := 0; i < len(times); i++ {
		var possibilities int

		var targetTime = times[i]
		var targetDistance = distances[i]

		for t := 1; t < targetTime; t++ {
			speed := t
			distance := (targetTime - t) * speed
			if distance > targetDistance {
				possibilities++
			}
		}

		total *= possibilities
	}

	fmt.Printf("Total: %d", total)
}

func makeList(input string) []int {

	// foo: 1   2  3   4   5
	var parts = strings.Split(strings.Split(input, ":")[1], " ")
	parts = slices.DeleteFunc(parts, func(n string) bool {
		return n == ""
	})
	var numbers = make([]int, len(parts))
	for i := 0; i < len(parts); i++ {
		number, err := strconv.Atoi(parts[i])
		check(err)

		numbers[i] = number
	}
	return numbers
}
