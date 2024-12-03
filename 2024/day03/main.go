package main

import (
	"fmt"
	"math"
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

	mulRegex, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	if err != nil {
		panic(fmt.Errorf("mul regex did not compile: %w", err))
	}

	doRegex, err := regexp.Compile(`do\(\)`)
	if err != nil {
		panic(fmt.Errorf("do regex did not compile: %w", err))
	}

	dontRegex, err := regexp.Compile(`don't\(\)`)
	if err != nil {
		panic(fmt.Errorf("don't regex did not compile: %w", err))
	}

	doMul := true
	var sum int64 = 0

	for {
		doIndex := doRegex.FindStringIndex(line)
		dontIndex := dontRegex.FindStringIndex(line)
		mulIndex := mulRegex.FindStringIndex(line)

		if mulIndex == nil {
			break
		}

		if len(doIndex) == 0 {
			doIndex = []int{math.MaxInt, math.MaxInt}
		}
		if len(dontIndex) == 0 {
			dontIndex = []int{math.MaxInt, math.MaxInt}
		}

		if mulIndex[0] < doIndex[0] && mulIndex[0] < dontIndex[0] {
			if doMul {
				mul := line[mulIndex[0]+4 : mulIndex[1]-1]
				sum += multiply(mul)
			}
			line = line[mulIndex[1]:]

		} else {
			if doIndex[0] < dontIndex[0] {
				line = line[doIndex[1]:]
				doMul = true
			} else {
				line = line[dontIndex[1]:]
				doMul = false
			}
		}
	}

	fmt.Printf("The sum is: %d\n", sum)
}

func multiply(text string) int64 {
	parts := strings.Split(text, ",")

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

	return first * second
}
