package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(fmt.Errorf("could not read input: %w", err))
	}

	parts := strings.Split(string(data), "\n\n")
	freshRangesString := strings.Split(parts[0], "\n")
	ingredients := strings.Split(parts[1], "\n")

	freshRanges := parseFreshRanges(freshRangesString)

	totalFresh := 0

	for _, ingredientString := range ingredients {
		ingredient, err := strconv.ParseInt(ingredientString, 10, 64)
		if err != nil {
			panic(err)
		}

		if isFresh(freshRanges, ingredient) {
			totalFresh++
		}
	}

	fmt.Printf("Total fresh ingredients are: %d\n", totalFresh)
}

func isFresh(freshRanges [][]int64, ingredient int64) bool {
	for _, freshRange := range freshRanges {
		if freshRange[0] <= ingredient && ingredient <= freshRange[1] {
			return true
		}
	}

	return false
}

func parseFreshRanges(ranges []string) [][]int64 {
	result := make([][]int64, len(ranges))

	for i, line := range ranges {
		split := strings.Split(line, "-")

		start, err := strconv.ParseInt(split[0], 10, 64)
		if err != nil {
			panic(err)
		}
		end, err := strconv.ParseInt(split[1], 10, 64)
		if err != nil {
			panic(err)
		}

		result[i] = []int64{start, end}
	}

	return result
}
