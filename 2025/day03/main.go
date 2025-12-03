package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(fmt.Errorf("could not read input: %w", err))
	}

	total := 0

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		total += getJoltage(line)
	}

	fmt.Printf("Total Joltage is: %d\n", total)
}

func getJoltage(line string) int {
	const offset rune = '0'

	firstLargest := 0
	secondLargest := 0

	firstLargestIndex := -1

	for index, c := range line {
		value := int(c - offset)

		if value > firstLargest {
			firstLargest = value
			firstLargestIndex = index
		}
	}

	var isFirstBeforeSecond bool
	var searchForSecondNumber string
	if firstLargestIndex+1 < len(line) {
		searchForSecondNumber = line[firstLargestIndex+1:]
		isFirstBeforeSecond = true
	} else {
		searchForSecondNumber = line[:firstLargestIndex]
		isFirstBeforeSecond = false
	}

	for _, c := range searchForSecondNumber {
		value := int(c - offset)

		if value > secondLargest {
			secondLargest = value
		}
	}

	if isFirstBeforeSecond {
		return firstLargest*10 + secondLargest
	}

	return firstLargest + secondLargest*10
}
