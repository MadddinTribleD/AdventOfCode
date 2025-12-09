package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(fmt.Errorf("could not read input: %w", err))
	}
	lines := bytes.Split(data, []byte("\n"))

	points := parse(lines)

	_ = points

	var maxArea int64 = 0

	for i, a := range points {
		for _, b := range points[i+1:] {
			maxArea = max(a.Area(b), maxArea)
		}
	}

	fmt.Printf("Largest area is: %d\n", maxArea)
}

type Point struct {
	X, Y int64
}

func (p Point) Area(q Point) int64 {
	minX := min(p.X, q.X)
	minY := min(p.Y, q.Y)

	maxX := max(p.X, q.X)
	maxY := max(p.Y, q.Y)

	// +1 because top is inclusive
	diffX := maxX - minX + 1
	diffY := maxY - minY + 1

	return diffX * diffY
}

func parse(lines [][]byte) []Point {
	result := make([]Point, len(lines))

	for l, line := range lines {
		parts := bytes.Split(line, []byte(","))

		if len(parts) != 2 {
			panic("Should be 2")
		}

		coord := make([]int64, 2)
		for i, p := range parts {
			value, err := strconv.ParseInt(string(p), 10, 64)
			if err != nil {
				panic(err)
			}
			coord[i] = value
		}

		result[l] = Point{
			X: coord[0],
			Y: coord[1],
		}
	}

	return result
}
