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

	lines := strings.Split(string(data), "\n")

	firstList := []int64{}
	secondList := map[int64]int64{}

	for l, line := range lines {
		parts := strings.Split(line, "   ")

		if len(parts) != 2 {
			panic("not two parts found")
		}

		first, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			panic(fmt.Errorf("first part in line %d could not be parsed to int: %w", l, err))
		}

		second, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			panic(fmt.Errorf("second part in line %d could not be parsed to int: %w", l, err))
		}

		firstList = append(firstList, first)

		secondList[second] = secondList[second] + 1
	}

	slices.Sort(firstList)

	var sum int64

	for _, i := range firstList {
		sum += i * secondList[i]
	}

	fmt.Printf("Total distance is: %d\n", sum)
}
