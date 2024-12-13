package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
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

	originalGrid := bytes.Split(slices.Clone(data), []byte{'\n'})
	grid := bytes.Split(data, []byte{'\n'})

	plots := [][]Coord{}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == byte(0) {
				continue
			}

			plot := searchPlot(grid, Coord{x: x, y: y})
			plots = append(plots, plot)
		}
	}

	sum := 0

	for _, plot := range plots {
		fence := getFence(originalGrid, plot)
		sum += fence * len(plot)
	}

	fmt.Printf("Total fence price: %d", sum)
}

func getFence(grid [][]byte, plot []Coord) int {
	fence := 0

	for _, c := range plot {
		neighbors := getNeighbors(grid, c, (grid)[c.y][c.x])

		fence += 4 - len(neighbors)
	}

	return fence
}

func searchPlot(grid [][]byte, start Coord) []Coord {
	searchPlot := (grid)[start.y][start.x]
	openList := map[Coord]any{
		start: true,
	}

	found := []Coord{}

	for len(openList) > 0 {
		var c Coord
		for cc := range openList {
			c = cc
		}
		delete(openList, c)

		found = append(found, c)
		(grid)[c.y][c.x] = 0

		for _, n := range getNeighbors(grid, c, searchPlot) {
			openList[n] = true
		}
	}

	return found
}

func getNeighbors(grid [][]byte, c Coord, searchPlot byte) []Coord {
	result := []Coord{}
	if c.y-1 >= 0 && grid[c.y-1][c.x] == searchPlot {
		result = append(result, Coord{y: c.y - 1, x: c.x})
	}
	if c.y+1 < len(grid) && grid[c.y+1][c.x] == searchPlot {
		result = append(result, Coord{y: c.y + 1, x: c.x})
	}

	if c.x-1 >= 0 && grid[c.y][c.x-1] == searchPlot {
		result = append(result, Coord{y: c.y, x: c.x - 1})
	}
	if c.x+1 < len(grid[c.y]) && grid[c.y][c.x+1] == searchPlot {
		result = append(result, Coord{y: c.y, x: c.x + 1})
	}

	return result
}
