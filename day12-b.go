package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	AssumeGood = 1
	AssumeBad  = 2
)

func main_12b() {
	fmt.Println("hello world")

	// bytes, err := os.ReadFile("day12-input-test.txt")
	bytes, err := os.ReadFile("day12-input.txt")
	check(err)

	text := string(bytes)

	lines := strings.Split(text, "\r\n")

	total := 0
	totalCalls := 0

	for l := 0; l < len(lines); l++ {
		line := lines[l]

		parts := strings.Split(line, " ")

		row := parts[0]
		// x5
		row = fmt.Sprintf("%s?%s?%s?%s?%s", row, row, row, row, row)
		damaged := strings.Split(fmt.Sprintf("%s,%s,%s,%s,%s", parts[1], parts[1], parts[1], parts[1], parts[1]), ",")
		damagedSprings := make([]int, len(damaged))
		for d := 0; d < len(damaged); d++ {
			n, err := strconv.Atoi(damaged[d])
			check(err)
			damagedSprings[d] = n
		}

		fmt.Println(row)
		fmt.Println(damagedSprings)

		results, calls := inspectSprings_b(&row, 0, Unspecified, Unspecified, &damagedSprings, 0, 0)

		fmt.Printf("Results %d, Calls: %d\n", results, calls)

		total += results
		totalCalls += calls
	}

	fmt.Printf("%d %d", total, totalCalls)
}

// Recursive search the length of the spring to see which combinations are possible.
// ???.###
// [1 1 3]
func inspectSprings_b(rowRef *string, i int, nextRequirement int, assumption int, damangedSpringsRef *[]int, damagedSection int, damageFound int) (int, int) {

	row := *rowRef
	damangedSprings := *damangedSpringsRef
	// fmt.Printf("%d %d %d %d %d\n", i, nextRequirement, assumption, damagedSection, damageFound)
	for {
		// good springs
		for ; i < len(row) && (row[i] == '.' || assumption == AssumeGood); i++ {
			assumption = Unspecified
			if nextRequirement == RequireBad {
				// fmt.Printf("%s %d %d %o %d Missing Bad\n", row, i, nextRequirement, damangedSprings, d)
				return 0, 1
			}
			nextRequirement = Unspecified
			damageFound = 0
		}

		// bad spring
		for ; i < len(row) && (row[i] == '#' || assumption == AssumeBad); i++ {
			assumption = Unspecified
			// We can't have two ## segments without a . in between
			if nextRequirement == RequireGood {
				// fmt.Printf("%s %d %d %o %d Missing Gap\n", row, i, nextRequirement, damangedSprings, d)
				return 0, 1
			}
			if damagedSection == len(damangedSprings) || damangedSprings[damagedSection] == 0 {
				// There shouldn't have been another damaged spring here
				// fmt.Printf("%s %d %d %o %d Too many bad springs\n", row, i, nextRequirement, damangedSprings, d)
				return 0, 1
			}
			damageFound++
			// end of a segment
			if damangedSprings[damagedSection] == damageFound {
				damagedSection++
				nextRequirement = RequireGood
			} else {
				nextRequirement = RequireBad
			}
		}

		// End of the row and we're not expecting any more damaged springs
		if i == len(row) {
			if damagedSection == len(damangedSprings) ||
				damagedSection == len(damangedSprings)-1 && damangedSprings[damagedSection] == damageFound {
				// fmt.Printf("%s %d %d %o %d Valid\n", row, i, nextRequirement, damangedSprings, d)
				return 1, 1
			}
			// fmt.Printf("%s %d %d %o %d Invalid\n", row, i, nextRequirement, damangedSprings, d)
			return 0, 1
		}

		// stop if we find uncertainty
		if row[i] == '?' {
			break
		}
	}

	// Count '.', '?', '#', and expected bad springs remaining
	good := strings.Count(row[i:], ".")
	// bad := strings.Count(row[i:], "#")
	// unknown := strings.Count(row[i:], "?")
	expectedBad := 0
	for d := damagedSection; d < len(damangedSprings); d++ {
		expectedBad += damangedSprings[d]
	}
	minRemaining := expectedBad - damageFound + good
	if minRemaining > len(row)-i {
		return 0, 1
	}
	minRemaining = expectedBad - damageFound + len(damangedSprings) - damagedSection - 1 // Account for spacing
	if minRemaining > len(row)-i {
		return 0, 1
	}

	total := 0
	totalCalls := 1
	if nextRequirement == Unspecified || nextRequirement == RequireGood {
		// Try '.'
		result, calls := inspectSprings_b(rowRef, i, nextRequirement, AssumeGood, damangedSpringsRef, damagedSection, damageFound)
		total += result
		totalCalls += calls
	}
	if nextRequirement == Unspecified || nextRequirement == RequireBad {
		// Try '#'
		result, calls := inspectSprings_b(rowRef, i, nextRequirement, AssumeBad, damangedSpringsRef, damagedSection, damageFound)
		total += result
		totalCalls += calls
	}

	return total, totalCalls
}
