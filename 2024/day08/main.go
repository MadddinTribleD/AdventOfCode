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

func (c Coord) Diff(d Coord) Coord {
	return Coord{
		x: d.x - c.x,
		y: d.y - c.y,
	}
}

func (c Coord) Add(d Coord) Coord {
	return Coord{
		x: c.x + d.x,
		y: c.y + d.y,
	}
}
func (c Coord) Sub(d Coord) Coord {
	return Coord{
		x: c.x - d.x,
		y: c.y - d.y,
	}
}

const (
	empty    byte = '.'
	antinode byte = '#'
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(fmt.Errorf("could not read input: %w", err))
	}

	grid := bytes.Split(data, []byte{'\n'})

	coords := map[byte][]Coord{}

	for y, line := range grid {
		for x, c := range line {
			if c != empty {
				cc, ok := coords[c]
				if !ok {
					cc = []Coord{}
				}

				cc = append(cc, Coord{x: x, y: y})

				coords[c] = cc
			}
		}
	}

	for _, c := range coords {
		// there will be no antinode for a single antenna
		if len(c) < 2 {
			continue
		}

		for s := 0; s < len(c); s++ {
			for e := s + 1; e < len(c); e++ {
				start := c[s]
				end := c[e]

				diff := start.Diff(end)

				top := start.Sub(diff)
				bottom := end.Add(diff)

				if CheckBounds(grid, top) {
					grid[top.y][top.x] = antinode
				}
				if CheckBounds(grid, bottom) {
					grid[bottom.y][bottom.x] = antinode
				}
			}
		}
	}

	var sum int64

	for _, line := range grid {
		for _, c := range line {
			if c == antinode {
				sum++
			}
		}
	}

	fmt.Printf("Sum off antinode: %d\n", sum)
}

func CheckBounds(grid [][]byte, c Coord) bool {
	if c.y >= 0 && c.y < len(grid) {
		if c.x >= 0 && c.x < len(grid[0]) {
			return true
		}
	}

	return false
}
