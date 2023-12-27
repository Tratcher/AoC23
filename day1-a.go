package main

// https://gobyexample.com/

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main_1a() {
	fmt.Println("hello world")

	bytes, err := os.ReadFile("day1-input-a.txt")
	check(err)

	var text = string(bytes)

	// fmt.Print(text)

	var total = 0

	var lines = strings.Split(text, "\r\n")

	for l := 0; l < len(lines); l++ {
		var line = lines[l]
		fmt.Print(l)
		fmt.Print(" " + line + "\n")
		var i = strings.IndexAny(line, "0123456789")
		var firstDigit = int(line[i]) - 48 // '0'
		i = strings.LastIndexAny(line, "0123456789")
		var lastDigit = int(line[i]) - 48 // '0'

		var value = firstDigit*10 + lastDigit
		fmt.Print("Value: ")
		fmt.Print(value)
		fmt.Print("\n")

		total += value
	}

	fmt.Print(total)
}
