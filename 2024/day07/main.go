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

	lines := strings.Split(string(data), "\n")

	var sum int64

	for _, line := range lines {
		p := strings.Split(line, ": ")

		result, err := strconv.ParseInt(p[0], 10, 64)
		if err != nil {
			panic(fmt.Errorf("failed to parse result: %w", err))
		}

		inputStrings := strings.Split(p[1], " ")

		inputs := []int64{}

		for _, pp := range inputStrings {
			i, err := strconv.ParseInt(pp, 10, 64)
			if err != nil {
				panic(fmt.Errorf("failed to parse input: %w", err))
			}

			inputs = append(inputs, i)
		}

		operantsMax := int64(1) << (len(inputs) - 1)

		found := false

		for operant := int64(0); operant < operantsMax; operant++ {

			c := int64(inputs[0])

			for i := 0; i < len(inputs)-1; i++ {
				input := inputs[i+1]

				if operant&(int64(1)<<i) == 0 {
					c = c * input
				} else {
					c = c + input
				}

				if c > result {
					break
				}
			}

			if c == result {
				found = true
				break
			}
		}

		if found {
			sum = sum + result
		}
	}

	fmt.Printf("Sum off results: %d\n", sum)
}
