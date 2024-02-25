package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Unspecified = 0
	RequireGood = 1
	RequireBad  = 2
)

func main() {
	fmt.Println("hello world")

	bytes, err := os.ReadFile("day12-input.txt")
	check(err)

	text := string(bytes)

	lines := strings.Split(text, "\r\n")

	total := 0

	for l := 0; l < len(lines); l++ {
		line := lines[l]

		parts := strings.Split(line, " ")

		row := parts[0]
		damaged := strings.Split(parts[1], ",")
		damagedSprings := make([]int, len(damaged))
		for d := 0; d < len(damaged); d++ {
			n, err := strconv.Atoi(damaged[d])
			check(err)
			damagedSprings[d] = n
		}

		fmt.Println(row)
		fmt.Println(damagedSprings)

		i := 0
		d := 0
		results := inspectSprings_a(row, i, Unspecified, damagedSprings, d)

		fmt.Printf("Results %d\n", results)

		total += results
	}

	fmt.Println(total)
}

// Recursive search the length of the spring to see which combinations are possible.
// ???.###
// [1 1 3]
func inspectSprings_a(row string, i int, nextRequirement int, damangedSprings []int, d int) int {

	// fmt.Printf("%s %d %d %o %d\n", row, i, nextRequirement, damangedSprings, d)
	for {
		segmentStart := i

		// good springs
		for ; i < len(row) && row[i] == '.'; i++ {
			if nextRequirement == RequireBad {
				// fmt.Printf("%s %d %d %o %d Missing Bad\n", row, i, nextRequirement, damangedSprings, d)
				return 0
			}
			nextRequirement = Unspecified
		}

		// bad spring
		for ; i < len(row) && row[i] == '#'; i++ {
			// We can't have two ## segments without a . in between
			if nextRequirement == RequireGood {
				// fmt.Printf("%s %d %d %o %d Missing Gap\n", row, i, nextRequirement, damangedSprings, d)
				return 0
			}
			if d == len(damangedSprings) || damangedSprings[d] == 0 {
				// There shouldn't have been another damaged spring here
				// fmt.Printf("%s %d %d %o %d Too many bad springs\n", row, i, nextRequirement, damangedSprings, d)
				return 0
			}
			damangedSprings[d]--
			// end of a segment
			if damangedSprings[d] == 0 {
				d++
				nextRequirement = RequireGood
			} else {
				nextRequirement = RequireBad
			}
		}

		// End of the row and we're not expecting any more damaged springs
		if i == len(row) {
			if d == len(damangedSprings) ||
				d == len(damangedSprings)-1 && damangedSprings[d] == 0 {
				// fmt.Printf("%s %d %d %o %d Valid\n", row, i, nextRequirement, damangedSprings, d)
				return 1
			}
			// fmt.Printf("%s %d %d %o %d Invalid\n", row, i, nextRequirement, damangedSprings, d)
			return 0
		}

		// stop if we stop making progress
		if i == segmentStart {
			break
		}
	}

	// must be a ?
	if row[i] != '?' {
		panic("Not a ???")
	}

	total := 0
	// Try '.'
	cpy := make([]int, len(damangedSprings))
	copy(cpy, damangedSprings)
	rowa := row[0:i] + "." + row[i+1:]
	total += inspectSprings_a(rowa, i, nextRequirement, cpy, d)

	// Try '#'
	cpy = make([]int, len(damangedSprings))
	copy(cpy, damangedSprings)
	rowb := row[0:i] + "#" + row[i+1:]
	total += inspectSprings_a(rowb, i, nextRequirement, cpy, d)

	return total
}
