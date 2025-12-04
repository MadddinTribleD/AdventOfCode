package main

import (
	"bytes"
	"fmt"
	"os"
)

const occupied byte = '@'

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(fmt.Errorf("could not read input: %w", err))
	}

	var total int64 = 0

	places := bytes.Split(data, []byte("\n"))

	total = 0
	const freePlacesThreshold = 4

	for y := 0; y < len(places); y++ {
		for x := 0; x < len(places[y]); x++ {
			if places[y][x] != occupied {
				continue
			}
			if findNeighborCount(places, y, x) < freePlacesThreshold {
				total++
			}
		}
	}

	fmt.Printf("Total accessible roles  is: %d\n", total)
}

func findNeighborCount(places [][]byte, y, x int) int {
	neighbors := 0

	for yy := max(0, y-1); yy < min(len(places), y+2); yy++ {
		for xx := max(0, x-1); xx < min(len(places[yy]), x+2); xx++ {
			if yy == y && xx == x {
				continue
			}

			if places[yy][xx] == occupied {
				neighbors++
			}
		}
	}

	return neighbors
}
