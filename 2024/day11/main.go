package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(fmt.Errorf("could not read input: %w", err))
	}

	stoneStrings := strings.Split(string(data), " ")
	stones := []int64{}

	for _, str := range stoneStrings {
		n, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			panic(fmt.Errorf("failed to parse '%s': %w", str, err))
		}

		stones = append(stones, n)
	}

	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}

	fmt.Printf("Number of stones: %d", len(stones))
}

func blink(stones []int64) []int64 {
	newStones := []int64{}

	for _, stone := range stones {
		if stone == 0 {
			newStones = append(newStones, 1)
			continue
		}

		digits, s := countDigits(stone)
		if digits%2 == 0 {
			part1, part2 := splitString(s)
			newStones = append(newStones, part1, part2)
			continue
		}

		newStones = append(newStones, stone*2024)
	}

	return newStones
}

func countDigits(n int64) (int, string) {
	s := strconv.FormatInt(n, 10)

	return len(s), s
}

func splitString(s string) (int64, int64) {
	mid := len(s) / 2
	part1String := s[:mid]
	part2String := s[mid:]

	part1, err := strconv.ParseInt(part1String, 10, 64)
	if err != nil {
		panic(fmt.Errorf("failed to parse '%s': %w", part1String, err))
	}
	part2, err := strconv.ParseInt(part2String, 10, 64)
	if err != nil {
		panic(fmt.Errorf("failed to parse '%s': %w", part2String, err))
	}

	return part1, part2
}
