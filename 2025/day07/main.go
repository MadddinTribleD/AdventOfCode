package main

import (
	"bytes"
	"fmt"
	"os"
)

const (
	Start    byte = 'S'
	Splitter byte = '^'
	Beam     byte = '|'
	Void     byte = '.'
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(fmt.Errorf("could not read input: %w", err))
	}
	lines := bytes.Split(data, []byte("\n"))

	for i := 0; i < len(lines[0]); i++ {
		if lines[0][i] == Start {
			lines[0][i] = Beam
		}
	}

	splits := 0

	for i := 1; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			isSplitter := lines[i][j] == Splitter

			if lines[i-1][j] == Beam {

				if isSplitter {
					if j > 0 && lines[i][j-1] == Void {
						lines[i][j-1] = Beam
					}
					if j < len(lines[i])-1 && lines[i][j+1] == Void {
						lines[i][j+1] = Beam
					}
					splits++
				} else {
					lines[i][j] = Beam
				}
			}
		}

		print(lines)
	}

	fmt.Printf("Tachyons splits: %d\n", splits)
}

func print(lines [][]byte) {
	for i := 1; i < len(lines); i++ {
		fmt.Println(string(lines[i]))
	}
	fmt.Println()
}
