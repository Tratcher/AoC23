package main

import (
	"fmt"
	"os"
	"strings"
)

func main_8a() {
	fmt.Println("hello world")

	bytes, err := os.ReadFile("day8-input.txt")
	check(err)

	text := string(bytes)

	lines := strings.Split(text, "\r\n")

	// First line is directions: LLR
	directions := lines[0]
	nodes := make(map[string]node)

	// Skip the directions and the blank line
	lines = lines[2:]

	// Build a map... AAA = (BBB, CCC)
	for l := 0; l < len(lines); l++ {
		line := lines[l]

		nodes[line[0:3]] = node{line[7:10], line[12:15]}
	}

	fmt.Println(nodes)

	// Follow the map
	total := 0

	current := "AAA"
	offset := 0
	for {
		direction := directions[offset%len(directions)]
		node := nodes[current]
		if direction == 'L' {
			current = node.left
		} else {
			current = node.right
		}
		total++
		offset++
		if current == "ZZZ" {
			break
		}
	}

	fmt.Println(total)
}

type node struct {
	left  string
	right string
}
