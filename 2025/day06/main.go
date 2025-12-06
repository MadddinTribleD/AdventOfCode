package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type Operator byte

const (
	Multiply Operator = '*'
	Add      Operator = '+'
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(fmt.Errorf("could not read input: %w", err))
	}
	lines := bytes.Split(data, []byte("\n"))

	lines = flip(lines)

	var grandTotal int64 = 0

	var operator Operator
	reset := true
	var currentCalculation int64 = 0

	for _, line := range lines {
		if reset {
			operator = Operator(line[len(line)-1])
			line = line[:len(line)-1]
			if operator == Add {
				currentCalculation = 0
			} else {
				currentCalculation = 1
			}
			reset = false
		}

		line = bytes.TrimSpace(line)

		if len(line) == 0 {
			grandTotal += currentCalculation
			reset = true
			continue
		}

		number, err := strconv.ParseInt(string(line), 10, 64)
		if err != nil {
			panic(err)
		}

		switch operator {
		case Multiply:
			currentCalculation *= number
		case Add:
			currentCalculation += number
		default:
			panic("Unknown Operation")
		}
	}

	// last row is not terminated with empty line
	grandTotal += currentCalculation

	fmt.Printf("Grand total is: %d\n", grandTotal)
}

func flip(data [][]byte) [][]byte {
	result := make([][]byte, len(data[0]))

	for column := 0; column < len(data[0]); column++ {
		result[column] = make([]byte, len(data))
	}

	for row := 0; row < len(data); row++ {
		for column := 0; column < len(data[row]); column++ {
			result[column][row] = data[row][column]
		}
	}

	return result
}
