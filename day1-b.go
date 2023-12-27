package main

// https://gobyexample.com/

import (
	"fmt"
	"os"
	"strings"
)

func main_1b() {
	var total = 0
	fmt.Println("hello world")

	bytes, err := os.ReadFile("day1-input-a.txt")
	check(err)

	var text = string(bytes)

	// var text = "two1nine\r\neightwothree\r\nabcone2threexyz\r\nxtwone3four\r\n4nineeightseven2\r\nzoneight234\r\n7pqrstsixteen"
	// fmt.Print(text)

	var lines = strings.Split(text, "\r\n")
	/*
		var lines = [7]string{
			"qvqgppfvrktjncgkfshzvhcxfzvtvgtwo37two",
			"eight2three",
			"abcone2threexyz",
			"xtwone3four",
			"4nineeightseven2",
			"zoneight234",
			"7pqrstsixteen"}
	*/
	for l := 0; l < len(lines); l++ {
		var line = lines[l]
		fmt.Print(l)
		fmt.Print(" " + line + "\n")
		var firstDigit = firstDigit(line)
		var lastDigit = lastDigit(line)

		var value = firstDigit*10 + lastDigit
		fmt.Print("Value: ")
		fmt.Print(value)
		fmt.Print("\n")

		total += value
	}

	fmt.Print(total)
}

func firstDigit(line string) int {

	var index = strings.IndexAny(line, "0123456789")
	var digit = int(line[index]) - 48 // '0'

	var digits = [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for d := 0; d < len(digits); d++ {
		var digitText = digits[d]
		var dIndex = strings.Index(line, digitText)
		if dIndex != -1 && dIndex < index {
			index = dIndex
			digit = d
		}
	}

	return digit
}

func lastDigit(line string) int {

	var index = strings.LastIndexAny(line, "0123456789")
	var digit = int(line[index]) - 48 // '0'

	var digits = [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for d := 0; d < len(digits); d++ {
		var digitText = digits[d]
		var dIndex = strings.LastIndex(line, digitText)
		if dIndex > index {
			index = dIndex
			digit = d
		}
	}

	return digit
}
