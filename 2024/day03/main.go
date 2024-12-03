package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(fmt.Errorf("could not read input: %w", err))
	}

	line := string(data)

	r, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	if err != nil {
		panic(fmt.Errorf("regex did not compile: %w", err))
	}

	results := r.FindAllString(line, -1)

	var sum int64 = 0

	for _, mul := range results {
		mul = strings.TrimLeft(mul, "mul(")
		mul = strings.TrimRight(mul, ")")

		parts := strings.Split(mul, ",")

		if len(parts) != 2 {
			panic("found wrong number of parts")
		}

		first, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			panic(fmt.Errorf("failed to parse part 1: %w", err))
		}

		second, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			panic(fmt.Errorf("failed to parse part 2: %w", err))
		}

		sum += first * second
	}

	fmt.Printf("The sum is: %d\n", sum)
}
