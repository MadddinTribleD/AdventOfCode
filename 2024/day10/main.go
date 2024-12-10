package main

import (
	"bytes"
	"fmt"
	"os"
)

type Coord struct {
	x int
	y int
}

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(fmt.Errorf("could not read input: %w", err))
	}

	grid := bytes.Split(data, []byte{'\n'})

	start := []Coord{}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '0' {
				start = append(start, Coord{x: x, y: y})
			}
		}
	}

	var sum int = 0
	for _, c0 := range start {

		for _, c1 := range getNeighbors(grid, c0, '1') {
			for _, c2 := range getNeighbors(grid, c1, '2') {
				for _, c3 := range getNeighbors(grid, c2, '3') {
					for _, c4 := range getNeighbors(grid, c3, '4') {
						for _, c5 := range getNeighbors(grid, c4, '5') {
							for _, c6 := range getNeighbors(grid, c5, '6') {
								for _, c7 := range getNeighbors(grid, c6, '7') {
									for _, c8 := range getNeighbors(grid, c7, '8') {
										for range getNeighbors(grid, c8, '9') {
											sum++
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	fmt.Printf("Trail-Score is: %d", sum)
}

func getNeighbors(grid [][]byte, c Coord, height byte) []Coord {
	result := []Coord{}
	if c.y-1 >= 0 && grid[c.y-1][c.x] == height {
		result = append(result, Coord{y: c.y - 1, x: c.x})
	}
	if c.y+1 < len(grid) && grid[c.y+1][c.x] == height {
		result = append(result, Coord{y: c.y + 1, x: c.x})
	}

	if c.x-1 >= 0 && grid[c.y][c.x-1] == height {
		result = append(result, Coord{y: c.y, x: c.x - 1})
	}
	if c.x+1 < len(grid[c.y]) && grid[c.y][c.x+1] == height {
		result = append(result, Coord{y: c.y, x: c.x + 1})
	}

	return result
}
