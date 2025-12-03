package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(fmt.Errorf("could not read input: %w", err))
	}

	var total int64 = 0

	lines := bytes.Split(data, []byte("\n"))
	for _, line := range lines {
		joltage := getJoltage(line)
		// fmt.Printf("%d\n", joltage)
		total += joltage
	}

	fmt.Printf("Total Joltage is: %d\n", total)
}

func getJoltage(line []byte) int64 {
	const offset byte = '0'
	const expectedBatteries = 12

	startIndex := 1
	for len(line) > expectedBatteries {
		for i := startIndex; i < len(line); i++ {
			previousValue := int(line[i-1] - offset)
			value := int(line[i] - offset)

			if previousValue < value || previousValue == value && i == len(line)-1 {
				line = removeAt(line, i-1)

				startIndex = 0
				break
			} else if previousValue > value && i == len(line)-1 {
				line = removeAt(line, i)

				startIndex = 0
			}
		}
		startIndex++
	}

	joltage, err := strconv.ParseInt(string(line), 10, 64)
	if err != nil {
		panic(fmt.Errorf("failed to parse: %w", err))
	}

	return joltage
}

func removeAt(line []byte, i int) []byte {
	removed := slices.Clone(line)
	removed = append(removed[:i], removed[i+1:]...)
	return removed
}
