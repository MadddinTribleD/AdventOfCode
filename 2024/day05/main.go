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

		wrongOrder := false

		for i, page := range pages {
			for _, r := range required {
				if r[0] == page {
					for j := 0; j < i; j++ {
						if r[1] == pages[j] {
							if i > j {
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
				break
			}
		}

		if !wrongOrder {
			middlePage, err := strconv.ParseInt(pages[len(pages)/2], 10, 64)
			if err != nil {
				panic(fmt.Errorf("could not parse middel page: %w", err))
			}

			sum += middlePage
		}
	}
	fmt.Printf("Correct ordered pages sum: %d\n", sum)
}
