package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(fmt.Errorf("could not read input: %w", err))
	}

	// key is stone number, value is count
	knownStones := map[int64]int64{}

	for _, str := range strings.Split(string(data), " ") {
		n, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			panic(fmt.Errorf("failed to parse '%s': %w", str, err))
		}

		knownStones[n]++
	}

	for i := 0; i < 75; i++ {
		newStones := map[int64]int64{}

		for stone, cnt := range knownStones {
			firstNewStone, secondNewStone := blinkStone(stone)

			newStones[firstNewStone] += cnt

			if secondNewStone != nil {
				newStones[*secondNewStone] += cnt
			}
		}

		knownStones = newStones
	}

	var sum int64 = 0

	for _, cnt := range knownStones {
		sum += cnt
	}

	fmt.Printf("Number of stones: %d", sum)
}

func blinkStone(stone int64) (int64, *int64) {

	if stone == 0 {
		return 1, nil
	}

	digits := digits(stone)
	if digits%2 == 0 {
		pow := int64(math.Pow10(int(digits) / 2))

		first := stone / pow
		second := stone % pow

		return first, &second
	}

	return stone * 2024, nil
}

func digits(i int64) int {
	return int(math.Log10(float64(i))) + 1
}
