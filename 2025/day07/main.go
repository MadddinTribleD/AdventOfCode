package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
)

const (
	Start    byte = 'S'
	Splitter byte = '^'
	Beam     byte = '|'
	Void     byte = '.'
)

var allVoidRegExp = regexp.MustCompile(`^\.*$`)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(fmt.Errorf("could not read input: %w", err))
	}
	lines := bytes.Split(data, []byte("\n"))

	splits := make([]map[int]int, 0, len(lines))
	for i := 0; i < len(lines[0]); i++ {
		if lines[0][i] == Start {
			lines[0][i] = Beam

			splits = append(splits, map[int]int{
				i: 1,
			})
			break
		}
	}

	for i := len(lines) - 1; i >= 0; i-- {
		if allVoidRegExp.Match(lines[i]) {
			lines = append(lines[:i], lines[i+1:]...)
		}
	}

	for i := 1; i < len(lines); i++ {
		splits = append(splits, map[int]int{})
		for j := 0; j < len(lines[i]); j++ {
			isSplitter := lines[i][j] == Splitter

			previous := splits[i-1][j]
			if lines[i-1][j] == Beam {
				if isSplitter {
					if j > 0 && lines[i][j-1] != Splitter {
						lines[i][j-1] = Beam
						splits[i][j-1] += previous
					}
					if j < len(lines[i])-1 && lines[i][j+1] != Splitter {
						lines[i][j+1] = Beam
						splits[i][j+1] += previous
					}

				} else {
					lines[i][j] = Beam
					splits[i][j] += previous
				}
			}
		}
	}
	print(lines)

	timelines := 0

	for _, v := range splits[len(splits)-1] {
		timelines += v
	}

	fmt.Printf("Total timelines: %d\n", timelines)
}

func print(lines [][]byte) {
	for i := 1; i < len(lines); i++ {
		fmt.Println(string(lines[i]))
	}
	fmt.Println()
}
