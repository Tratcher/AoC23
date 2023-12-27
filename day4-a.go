package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main_4a() {
	var total = 0
	fmt.Println("hello world")

	bytes, err := os.ReadFile("day4-input.txt")
	check(err)

	var text = string(bytes)

	var lines = strings.Split(text, "\r\n")

	for r := 0; r < len(lines); r++ {
		total += scoreCard(lines[r])
	}

	fmt.Print("\n")
	fmt.Print(total)
}

func scoreCard(card string) int {
	var parts = strings.Split(strings.Split(card, ":")[1], "|")
	var winningNumbers = slices.DeleteFunc(strings.Split(parts[0], " "), func(n string) bool {
		return n == ""
	})
	var ourNumbers = slices.DeleteFunc(strings.Split(parts[1], " "), func(n string) bool {
		return n == ""
	})

	var score = 0

	for v := 0; v < len(ourNumbers); v++ {
		if slices.Contains(winningNumbers, ourNumbers[v]) {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}

	fmt.Printf("%s; %d\n", card, score)

	return score
}
