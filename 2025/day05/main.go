package main

import (
	"fmt"
	"os"
	"slices"
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

	freshRanges := parseFreshRanges(freshRangesString)

	slices.SortFunc(freshRanges, func(first, second []int64) int {
		return int(first[0] - second[0])
	})

	freshRanges = reduceOverlaps(freshRanges)

	var totalFreshIds int64 = 0

	for _, freshRange := range freshRanges {
		totalFreshIds += freshRange[1] - freshRange[0] + 1
	}

	fmt.Printf("Total fresh ingredients Ids are: %d\n", totalFreshIds)
}

func reduceOverlaps(ranges [][]int64) [][]int64 {
	// How?
	if len(ranges) < 1 {
		return nil
	}

	mergedRanges := [][]int64{}

	for currentIndex := 0; currentIndex < len(ranges); currentIndex++ {
		current := ranges[currentIndex]

		for nextIndex := currentIndex + 1; nextIndex < len(ranges); nextIndex++ {
			// next start is in out current range
			if current[1] >= ranges[nextIndex][0] {
				// move our end
				current[1] = max(current[1], ranges[nextIndex][1])
				currentIndex = nextIndex
			} else {
				// skip all until nextIndex
				currentIndex = nextIndex - 1
				break
			}
		}

		mergedRanges = append(mergedRanges, current)
	}

	return mergedRanges
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
