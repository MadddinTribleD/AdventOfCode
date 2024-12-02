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

	var safeCnt int64 = 0

	for _, lineValues := range values {
		diffs := getDiffs(lineValues)

		safe := isSafe(diffs)

		if safe {
			safeCnt++
		} else {

			anySafe := false
			for i := 0; i < len(lineValues); i++ {
				tryAgain := append([]int64{}, lineValues[:i]...)
				tryAgain = append(tryAgain, lineValues[i+1:]...)
				diffs = getDiffs(tryAgain)

				if safe = isSafe(diffs); safe {
					anySafe = true
					break
				}
			}
			if anySafe {
				safeCnt++
			}
		}

	}

	fmt.Printf("Safe Reports: %d\n", safeCnt)
}

func getDiffs(values []int64) []int64 {
	diffs := []int64{}

	for i := 1; i < len(values); i++ {
		diffs = append(diffs, values[i]-values[i-1])
	}

	return diffs
}

func isSafe(diffs []int64) bool {
	increasing := false

	if diffs[0] > 0 {
		increasing = true
	}

	for i := 0; i < len(diffs); i++ {
		if diffs[i] == 0 {
			return false
		}

		if increasing && (diffs[i] > maxIncrease || diffs[i] <= 0) {
			return false
		} else if !increasing && (diffs[i] < -maxIncrease || diffs[i] >= 0) {
			return false
		}

	}

	return true
}
