package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(fmt.Errorf("could not read input: %w", err))
	}
	lines := bytes.Split(data, []byte("\n"))

	wirings := parse(lines)

	for _, wiring := range wirings {
		fmt.Println(wiring.TrySolve())
	}

	fmt.Printf("Largest area is: %d\n", wirings)
}

type ButtonWiring struct {
	lightsInt  int
	wiringsInt []int

	Lights   []bool
	Wirings  [][]bool
	Joltages []int64
}

func (b ButtonWiring) TrySolve() int {
	light := b.lightsInt
	cnt := 0

	smaller := slices.Clone(b.wiringsInt)
	for light > 0 {
		smaller = findSmaller(smaller, light)
		if len(smaller) == 0 {
			panic("aaahh")
		}

		cnt++
		newLights := light ^ smaller[0]
		light = newLights
	}

	return cnt
}

func findSmaller(values []int, search int) []int {
	results := make([]int, 0, len(values))

	for _, v := range values {
		if v <= search {
			results = append(results, v)
		}
	}

	slices.Sort(results)
	slices.Reverse(results)

	return results
}

func NewButtonWiring(
	lights []bool,
	wirings [][]bool,
	joltages []int64,
) ButtonWiring {
	wiringsInt := make([]int, len(wirings))

	for i, wiring := range wirings {
		wiringsInt[i] = boolToBitMask(wiring)
	}

	return ButtonWiring{
		lightsInt:  boolToBitMask(lights),
		wiringsInt: wiringsInt,

		Lights:   lights,
		Wirings:  wirings,
		Joltages: joltages,
	}
}

func boolToBitMask(bools []bool) int {
	result := 0
	for i := 0; i < len(bools); i++ {
		if bools[i] {
			result += 1 << i
		}
	}

	return result
}

func parse(lines [][]byte) []ButtonWiring {
	buttonWirings := make([]ButtonWiring, len(lines))

	for i, line := range lines {
		parts := bytes.Split(line, []byte(" "))

		lights := make([]bool, len(parts[0])-2)
		wirings := make([][]bool, len(parts)-2)

		for l, light := range parts[0][1 : len(parts[0])-1] {
			lights[l] = light == '#'
		}

		for w, wiringString := range parts[1 : len(parts)-1] {
			numberStrings := bytes.Split(wiringString[1:len(wiringString)-1], []byte(","))

			wiring := make([]bool, len(parts[0])-1)

			for _, numberStrings := range numberStrings {
				number, err := strconv.ParseInt(string(numberStrings), 10, 64)
				if err != nil {
					panic(err)
				}

				wiring[number] = true
			}

			wirings[w] = wiring
		}

		joltageString := parts[len(parts)-1]
		numberStrings := bytes.Split(joltageString[1:len(joltageString)-1], []byte(","))

		joltage := make([]int64, len(lights))

		for j, numberStrings := range numberStrings {
			number, err := strconv.ParseInt(string(numberStrings), 10, 64)
			if err != nil {
				panic(err)
			}

			joltage[j] = number
		}

		buttonWirings[i] = NewButtonWiring(lights, wirings, joltage)
	}

	return buttonWirings
}
