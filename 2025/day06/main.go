package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Operator byte

const (
	Multiply Operator = '*'
	Add      Operator = '+'
)

var (
	whiteSpaceRegEx = regexp.MustCompile(`\s+`)
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(fmt.Errorf("could not read input: %w", err))
	}
	lines := bytes.Split(data, []byte("\n"))

	allValues := make([][]int64, len(lines)-1)
	var operators []Operator

	for lineIndex := range lines {
		line := bytes.TrimSpace(lines[lineIndex])
		line = whiteSpaceRegEx.ReplaceAll(line, []byte(","))
		parts := bytes.Split(line, []byte(","))

		if lineIndex == len(lines)-1 {
			operators = make([]Operator, len(parts))
			for i, p := range parts {
				operators[i] = Operator(p[0])
			}
		} else {
			values := make([]int64, len(parts))

			for i, p := range parts {
				number, err := strconv.ParseInt(string(p), 10, 64)
				if err != nil {
					panic(err)
				}

				values[i] = number
			}

			allValues[lineIndex] = values
		}
	}

	var grandTotal int64 = 0

	for i, operator := range operators {
		value := allValues[0][i]

		for lineIndex := 1; lineIndex < len(allValues); lineIndex++ {
			switch operator {
			case Multiply:
				value *= allValues[lineIndex][i]
			case Add:
				value += allValues[lineIndex][i]
			default:
				panic("Unknown Operator")
			}
		}

		grandTotal += value
	}

	fmt.Printf("Grand total is: %d\n", grandTotal)
}
