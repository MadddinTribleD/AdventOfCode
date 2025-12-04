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

	places := bytes.Split(data, []byte("\n"))

	const freePlacesThreshold = 4

	removedRoles := [][]int{}
	foundNewRoles := true

	for foundNewRoles {
		foundNewRoles = false
		for y := 0; y < len(places); y++ {
			for x := 0; x < len(places[y]); x++ {
				if places[y][x] != occupied {
					continue
				}
				if findNeighborCount(places, y, x) < freePlacesThreshold {
					removedRoles = append(removedRoles, []int{y, x})
					foundNewRoles = true
				}
			}
		}

		for _, place := range removedRoles {
			places[place[0]][place[1]] = 'x'
		}
	}

	fmt.Printf("Total removed roles  is: %d\n", len(removedRoles))
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
