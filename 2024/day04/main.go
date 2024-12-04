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

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(fmt.Errorf("could not read input: %w", err))
	}

	lines := strings.Split(string(data), "\n")

	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	searchMap := [][]string{
		{
			"M.S",
			".A.",
			"M.S",
		},
		{
			"M.M",
			".A.",
			"S.S",
		},
		{
			"S.M",
			".A.",
			"S.M",
		},
		{
			"S.S",
			".A.",
			"M.M",
		},
	}
	sum := 0

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {

			for _, s := range searchMap {
				found := true
				for yy := 0; yy < len(s) && found; yy++ {
					for xx := 0; xx < len(s[yy]) && found; xx++ {
						if s[yy][xx:xx+1] == "." {
							continue
						}

						if y+yy >= len(lines) {
							found = false
							break
						}
						if x+xx >= len(lines[y+yy]) {
							found = false
							break
						}

						if s[yy][xx:xx+1] != lines[y+yy][x+xx:x+xx+1] {
							found = false
							break
						}
					}
				}

				if found {
					sum++
				}
			}
		}
	}

	fmt.Printf("Found times: %d\n", sum)
}
