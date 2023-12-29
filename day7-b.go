package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main_7b() {
	fmt.Println("hello world")

	bytes, err := os.ReadFile("day7-input.txt")
	check(err)

	text := string(bytes)

	lines := strings.Split(text, "\r\n")

	hands := make([]hand, len(lines))

	for l := 0; l < len(lines); l++ {
		// 32T3K 765
		line := lines[l]
		parts := strings.Split(line, " ")

		hands[l].input = parts[0]

		bet, err := strconv.Atoi(parts[1])
		check(err)

		hands[l].bet = bet
		hands[l].sorted = sortInputB(hands[l].input)
		hands[l].rank = rankHandB(hands[l].sorted)

		fmt.Printf("%s %s %d %d\n", hands[l].input, hands[l].sorted, hands[l].bet, hands[l].rank)
	}

	sort.Slice(hands, func(a, b int) bool {
		if hands[a].rank != hands[b].rank {
			return hands[a].rank < hands[b].rank
		}

		for i := 0; i < len(hands[a].input); i++ {
			if hands[a].input[i] != hands[b].input[i] {
				return compareCardsB(hands[a].input[i], hands[b].input[i])
			}
		}

		return false
	})

	fmt.Println(hands)

	total := 0

	// Now, you can determine the total winnings of this set of hands by adding up the result of multiplying each hand's bid with its rank (765 * 1 + 220 * 2 + 28 * 3 + 684 * 4 + 483 * 5).
	for h := 0; h < len(hands); h++ {
		hand := hands[h]
		total += hand.bet * (h + 1)
	}

	fmt.Println(total)
}

// 32T3K -> 233TK
func sortInputB(input string) string {
	sorted := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		sorted[i] = input[i]
	}

	sort.Slice(sorted, func(a, b int) bool {
		return compareCardsB(sorted[a], sorted[b])
	})

	return string(sorted)
}

func compareCardsB(a, b byte) bool {
	// 2-9,T,J,Q,K,A (but ASCII)
	if a == 'J' {
		return a != b // Jokers are lowest
	}
	if b == 'J' {
		return false // A isn't a joker and can't be less than one
	}
	if a <= '9' || b <= '9' {
		return a < b
	}
	switch a {
	case 'T':
		return b != 'T'
	case 'Q':
		return b == 'K' || b == 'A'
	case 'K':
		return b == 'A'
	case 'A':
		return false
	}
	return false
}

// 233TK -> Type_Pair
func rankHandB(input string) int {

	jokers := strings.Count(input, "J")

	if input[0] == input[1] &&
		input[0] == input[2] &&
		input[0] == input[3] &&
		input[0] == input[4] {
		return Type_FiveOfAKind
	}

	// They're sorted so the middle three must match eachother and one of the ends
	if input[1] == input[2] &&
		input[1] == input[3] &&
		(input[1] == input[0] ||
			input[1] == input[4]) {
		if jokers == 1 || jokers == 4 {
			return Type_FiveOfAKind
		}
		return Type_FourOfAKind
	}

	// Full house: 11222 or 11122
	if input[0] == input[1] &&
		input[2] == input[3] &&
		input[2] == input[4] {
		if jokers == 2 || jokers == 3 {
			return Type_FiveOfAKind
		}
		return Type_FullHouse
	}
	if input[0] == input[1] &&
		input[0] == input[2] &&
		input[3] == input[4] {
		if jokers == 2 || jokers == 3 {
			return Type_FiveOfAKind
		}
		return Type_FullHouse
	}

	// Three of a kind. The center must be one of them. 11123 12223 12333
	if input[2] == input[0] ||
		input[2] == input[4] ||
		input[2] == input[1] && input[2] == input[3] {
		if jokers == 1 || jokers == 3 {
			return Type_FourOfAKind
		}
		return Type_ThreeOfAKind
	}

	// Two pair 11223 11233 12233
	if input[0] == input[1] &&
		(input[2] == input[3] ||
			input[3] == input[4]) {
		if jokers == 1 {
			return Type_FullHouse
		}
		if jokers == 2 {
			return Type_FourOfAKind
		}
		return Type_TwoPair
	}
	if input[1] == input[2] &&
		input[3] == input[4] {
		if jokers == 1 {
			return Type_FullHouse
		}
		if jokers == 2 {
			return Type_FourOfAKind
		}
		return Type_TwoPair
	}

	// One pair 11234 12234 12334 12344
	if input[0] == input[1] ||
		input[1] == input[2] ||
		input[2] == input[3] ||
		input[3] == input[4] {
		if jokers == 1 || jokers == 2 {
			return Type_ThreeOfAKind
		}
		return Type_Pair
	}

	if jokers == 1 {
		return Type_Pair
	}

	return Type_HighCard
}
