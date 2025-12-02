package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(fmt.Errorf("could not read input: %w", err))
	}

	ranges := strings.Split(string(data), ",")

	var total int64 = 0

	for _, r := range ranges {
		startend := strings.Split(r, "-")
		if len(startend) != 2 {
			panic("not start and end")
		}

		start, err := strconv.ParseInt(startend[0], 10, 64)
		if err != nil {
			panic(fmt.Errorf("failed to parse: %w", err))
		}

		end, err := strconv.ParseInt(startend[1], 10, 64)
		if err != nil {
			panic(fmt.Errorf("failed to parse: %w", err))
		}

		for i := start; i <= end; i++ {
			if !isValid(i) {
				total += i
			}
		}
	}

	fmt.Printf("Invalid Id Total is: %d\n", total)
}

func isValid(i int64) bool {
	digits := int64(math.Floor(math.Log10(float64(i)) + 1))

	if digits%2 == 1 {
		return true
	}

	half := int64(math.Pow(10, float64(digits)/2))

	first := i % int64(half)
	second := i / int64(half)

	return first != second
}
