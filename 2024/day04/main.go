package main

import (
	"fmt"
	"os"
	"strings"
)

type direction struct {
	x int
	y int
}

var searchWord = "XMAS"

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(fmt.Errorf("could not read input: %w", err))
	}

	lines := strings.Split(string(data), "\n")

	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	directions := []direction{
		{x: -1, y: -1},
		{x: 0, y: -1},
		{x: 1, y: -1},

		{x: -1, y: 0},
		{x: 1, y: 0},

		{x: -1, y: 1},
		{x: 0, y: 1},
		{x: 1, y: 1},
	}

	sum := 0

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if lines[y][x:x+1] == searchWord[0:1] {
				for _, d := range directions {
					found := true
					for l := 1; l < len(searchWord); l++ {

						if y+d.y*l < 0 || y+d.y*l >= len(lines) {
							found = false
							break
						}
						if x+d.x*l < 0 || x+d.x*l >= len(lines[y+d.y*l]) {
							found = false
							break
						}

						if lines[y+d.y*l][x+d.x*l:x+d.x*l+1] != searchWord[l:l+1] {
							found = false
							break
						}
					}

					if found {
						sum++
					}
				}
			}
		}
	}

	fmt.Printf("Found '%s' times: %d\n", searchWord, sum)
}
