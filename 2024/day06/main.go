package main

import (
	"bytes"
	"fmt"
	"os"
)

type Cord struct {
	x int
	y int
}

const (
	start    byte = '^'
	obstacle byte = '#'
	visited  byte = 'X'
)

var nextDirections = map[Cord]Cord{
	{x: 0, y: -1}: {x: 1, y: 0},
	{x: 1, y: 0}:  {x: 0, y: 1},
	{x: 0, y: 1}:  {x: -1, y: 0},
	{x: -1, y: 0}: {x: 0, y: -1},
}

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(fmt.Errorf("could not read input: %w", err))
	}

	grid := bytes.Split(data, []byte{'\n'})
	_ = grid

	current := Cord{x: 0, y: 0}
	direction := Cord{x: 0, y: -1}

	for y := range grid {
		found := false
		for x := range grid[y] {
			if grid[y][x] == byte('^') {
				current = Cord{x: x, y: y}
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	defer func() {

		cnt := 0
		for _, y := range grid {
			for _, x := range y {
				if x == visited {
					cnt++
				}
			}
		}

		printGrid(grid)

		fmt.Printf("Guard visited spaces: %d\n", cnt)
		os.Exit(0)
	}()

	grid[current.y][current.x] = visited

	for {
		next := Cord{
			x: current.x + direction.x,
			y: current.y + direction.y,
		}

		if grid[next.y][next.x] == obstacle {
			direction = nextDirections[direction]
		} else {
			current = next
			grid[current.y][current.x] = visited
		}
	}

	panic("How???")
}

func printGrid(grid [][]byte) {
	for _, y := range grid {
		fmt.Println(string(y))
	}
	fmt.Println()
}
