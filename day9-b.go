package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hello world")

	bytes, err := os.ReadFile("day9-input.txt")
	check(err)

	text := string(bytes)

	lines := strings.Split(text, "\r\n")

	total := 0

	for l := 0; l < len(lines); l++ {
		strvalues := strings.Split(lines[l], " ")
		values := make([]int, len(strvalues))

		for i := 0; i < len(strvalues); i++ {
			number, err := strconv.Atoi(strvalues[i])
			check(err)
			values[i] = number
		}

		total += calculateNext_b(values)
	}

	fmt.Println(total)
}

func calculateNext_b(values []int) int {

	allZero := true
	for i := 0; i < len(values); i++ {
		if values[i] != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		return 0
	}

	nextSequence := make([]int, len(values)-1)

	for i := 0; i < len(nextSequence); i++ {
		a := values[i]
		b := values[i+1]

		nextSequence[i] = b - a
	}

	return values[0] - calculateNext_b(nextSequence)
}
