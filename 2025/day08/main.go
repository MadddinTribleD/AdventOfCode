package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"

	"github.com/samber/lo"
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(fmt.Errorf("could not read input: %w", err))
	}
	lines := bytes.Split(data, []byte("\n"))

	junctionBoxes := parse(lines)

	allConnections := []Connection{}

	for i := 0; i < len(junctionBoxes); i++ {
		a := junctionBoxes[i]
		for j := i + 1; j < len(junctionBoxes); j++ {
			b := junctionBoxes[j]

			allConnections = append(allConnections, NewConnection(a, b))
		}
	}

	slices.SortFunc(allConnections, func(a, b Connection) int {
		switch {
		case a.Distance < b.Distance:
			return -1
		case a.Distance > b.Distance:
			return 1
		case a.Distance == b.Distance:
			return 0
		}
		panic("how?")
	})

	for i := 0; i < 10; i++ {
		fmt.Printf("%f\n", allConnections[i].Distance)
	}

	gridLookup := map[*JunctionBox]*Grid{}
	allGrids := []*Grid{}

	for i := 0; i < 1000; i++ {

		if i > 800-2 && i < 801 {
			fmt.Printf("line:%d\n", i+1)

			sizes := lo.GroupBy(allGrids, func(g *Grid) int { return g.Size() })

			for size, grid := range sizes {
				fmt.Printf("Grid: %d, Size: %d\n", len(grid), size)
			}
			fmt.Println()
		}

		connection := allConnections[i]

		gridA, hasA := gridLookup[connection.A]
		gridB, hasB := gridLookup[connection.B]

		if hasA && !hasB {
			// A has a grid => Attach to A
			gridA.Connections = append(gridA.Connections, connection)
			gridLookup[connection.B] = gridA
		} else if !hasA && hasB {
			// B has a grid => Attach to B
			gridB.Connections = append(gridB.Connections, connection)
			gridLookup[connection.A] = gridB
		} else if !hasA && !hasB {
			// Both have no grid => create a new one
			grid := &Grid{
				Connections: []Connection{connection},
			}
			allGrids = append(allGrids, grid)
			gridLookup[connection.A] = grid
			gridLookup[connection.B] = grid
		} else if gridA != gridB {
			gridA.Connections = append(gridA.Connections, gridB.Connections...)

			for _, c := range gridB.Connections {
				gridLookup[c.A] = gridA
				gridLookup[c.B] = gridA
			}
			gridB.Connections = []Connection{}
		}
	}

	slices.SortFunc(allGrids, func(a, b *Grid) int {
		return b.Size() - a.Size()
	})

	sizes := make([]int, len(allGrids))
	for i := 0; i < len(allGrids); i++ {
		sizes[i] = allGrids[i].Size()
	}

	totalSize := sizes[0] * sizes[1] * sizes[2]

	fmt.Printf("Count of connections: %d\n", len(allConnections))
	fmt.Printf("Count of grids: %d\n", len(allGrids))
	fmt.Printf("Total size: %d\n", totalSize)
}

type Grid struct {
	Connections []Connection
}

func (g Grid) Size() int {
	junctionBoxesInGrid := map[*JunctionBox]bool{}

	for j := 0; j < len(g.Connections); j++ {
		c := g.Connections[j]
		junctionBoxesInGrid[c.A] = true
		junctionBoxesInGrid[c.B] = true
	}

	return len(junctionBoxesInGrid)
}

func NewConnection(a, b *JunctionBox) Connection {
	return Connection{
		Distance: a.Distance(*b),
		A:        a,
		B:        b,
	}
}

type Connection struct {
	Distance float64
	A, B     *JunctionBox
}

type JunctionBox struct {
	X, Y, Z int64
}

func (a JunctionBox) Distance(b JunctionBox) float64 {
	return math.Sqrt(float64(
		(a.X-b.X)*(a.X-b.X) +
			(a.Y-b.Y)*(a.Y-b.Y) +
			(a.Z-b.Z)*(a.Z-b.Z),
	))
}

func parse(lines [][]byte) []*JunctionBox {
	result := make([]*JunctionBox, len(lines))

	for l, line := range lines {
		parts := bytes.Split(line, []byte(","))

		if len(parts) != 3 {
			panic("Should be 3")
		}

		coord := make([]int64, 3)
		for i, p := range parts {
			value, err := strconv.ParseInt(string(p), 10, 64)
			if err != nil {
				panic(err)
			}
			coord[i] = value
		}

		result[l] = &JunctionBox{
			X: coord[0],
			Y: coord[1],
			Z: coord[2],
		}
	}

	return result
}
