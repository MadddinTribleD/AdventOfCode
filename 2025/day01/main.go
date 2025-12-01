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

	var countZeroes int64 = 0

	const start int64 = 50
	const max int64 = 100

	current := start

	for _, line := range lines {
		d := line[0]
		line = line[1:]
		c, err := strconv.ParseInt(line, 10, 64)

		if err != nil {
			panic(fmt.Errorf("could not parse: %w", err))
		}

		var direction int64 = 1
		if d == 'L' {
			direction = -1
		}

		for range c {
			current += direction

			if current%max == 0 {
				countZeroes++
			}
		}

	}

	fmt.Printf("Password is: %d\n", countZeroes)
}
