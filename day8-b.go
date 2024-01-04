package main

import (
	"fmt"
	"os"
	"strings"
)

func main_8b() {
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

	// Find all nodes that end in "A"
	startNodes := make([]string, 0)
	for k := range nodes {
		if k[2] == 'A' {
			startNodes = append(startNodes, k)
		}
	}

	current := make([]string, len(startNodes))
	for n := 0; n < len(current); n++ {
		current[n] = startNodes[n]
	}

	totalDistances := make([]int64, len(startNodes))
	sprints := make([]int64, len(startNodes))
	offsets := make([]int64, len(startNodes))
	cycleFound := make([]bool, len(startNodes))
	totalCycles := 0

	fmt.Println(current)

	// Follow the map
	var total int64 = 0
	var offset int64 = 0
	for {
		direction := directions[offset]

		zCount := 0
		for n := 0; n < len(current); n++ {

			node := nodes[current[n]]
			if direction == 'L' {
				current[n] = node.left
			} else {
				current[n] = node.right
			}

			if current[n][2] == 'Z' {
				zCount++

				sprint := total + 1 - totalDistances[n]
				totalDistances[n] = total + 1

				if !cycleFound[n] && sprint == sprints[n] {
					cycleFound[n] = true
					totalCycles++
				} else if sprint != sprints[n] {
					if sprints[n] > 0 {
						offsets[n] = sprints[n] - sprint
					}
					sprints[n] = sprint
				}
			}
		}

		total++
		offset = total % int64(len(directions))

		// Stop when all current nodes end in Z

		if zCount == len(current) || totalCycles == len(current) {
			break
		}

		if zCount > 2 {

			fmt.Printf("Zs: %d; %d\n", zCount, total)
			fmt.Println(current)
			fmt.Println(sprints)
		}
	}

	fmt.Println(current)
	fmt.Println(offsets)
	fmt.Println(sprints)
	fmt.Println(total)

	total = 277
	for i := 0; i < len(sprints); i++ {
		total *= sprints[i] / 277
	}

	fmt.Println(total)

}
