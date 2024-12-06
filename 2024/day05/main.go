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

	parts := strings.Split(string(data), "\n\n")

	firstParts := strings.Split(parts[0], "\n")
	secondParts := strings.Split(parts[1], "\n")

	required := [][]string{}

	for _, p := range firstParts {
		requiredPages := strings.Split(p, "|")
		required = append(required, requiredPages)
	}

	var sum int64

	for _, p := range secondParts {
		pages := strings.Split(p, ",")

		hadAnyWrongOrder := false
		wrongOrder := false

		for i := 0; i < len(pages); i++ {
			for _, r := range required {
				if r[0] == pages[i] {
					for j := 0; j < i; j++ {
						if r[1] == pages[j] {
							if i > j {
								hadAnyWrongOrder = true
								pageI := pages[i]
								pages[i] = pages[j]
								pages[j] = pageI
								wrongOrder = true
								break
							}
						}
					}
					if wrongOrder {
						break
					}
				}
			}
			if wrongOrder {
				wrongOrder = false
				i = -1
			}
		}

		if hadAnyWrongOrder {
			middlePage, err := strconv.ParseInt(pages[len(pages)/2], 10, 64)
			if err != nil {
				panic(fmt.Errorf("could not parse middel page: %w", err))
			}

			sum += middlePage
		}
	}
	fmt.Printf("Correct ordered pages sum: %d\n", sum)
}
