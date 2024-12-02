package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxIncrease int64 = 3

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(fmt.Errorf("could not read input: %w", err))
	}

	lines := strings.Split(string(data), "\n")

	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	values := [][]int64{}

	for l, line := range lines {
		parts := strings.Split(line, " ")

		lineValues := []int64{}

		for i, part := range parts {
			value, err := strconv.ParseInt(part, 10, 64)
			if err != nil {
				panic(fmt.Errorf("could not parse line %d part %d: %w", l, i, err))
			}

			lineValues = append(lineValues, value)
		}

		values = append(values, lineValues)
	}

	var safe int64 = 0

	for _, lineValues := range values {
		diffs := []int64{}

		for i := 1; i < len(lineValues); i++ {
			diffs = append(diffs, lineValues[i]-lineValues[i-1])
		}

		isSafe := true
		increasing := false

		if diffs[0] > 0 {
			increasing = true
		}

		for i := 0; i < len(diffs); i++ {
			if diffs[i] == 0 {
				isSafe = false
				break
			}

			if increasing && (diffs[i] > maxIncrease || diffs[i] <= 0) {
				isSafe = false
				break
			} else if !increasing && (diffs[i] < -maxIncrease || diffs[i] >= 0) {
				isSafe = false
				break
			}

		}

		if isSafe {
			safe++
		}
	}

	fmt.Printf("Safe Reports: %d\n", safe)
}
