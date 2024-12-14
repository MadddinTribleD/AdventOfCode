package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ClawMachine struct {
	AX int64
	AY int64
	BX int64
	BY int64
	X  int64
	Y  int64
}

type ButtonPresses struct {
	A int64
	B int64
}

func (b *ButtonPresses) Tokens() int64 {
	return b.A*3 + b.B
}

func NewClawMachine(match []string) ClawMachine {
	if len(match) != 7 {
		panic("not 7 matches")
	}

	values := []int64{}

	for _, m := range match[1:] {
		v, err := strconv.ParseInt(m, 10, 64)
		if err != nil {
			panic(fmt.Errorf("could not parse: %w", err))
		}

		values = append(values, v)
	}

	return ClawMachine{
		AX: values[0],
		AY: values[1],
		BX: values[2],
		BY: values[3],
		X:  values[4] + 10000000000000,
		Y:  values[5] + 10000000000000,
	}
}

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(fmt.Errorf("could not read input: %w", err))
	}

	claws := strings.Split(string(data), "\n\n")

	regex := regexp.MustCompile(`X\+(?P<AX>\d+).*Y\+(?P<AY>\d+).*X\+(?P<BX>\d+).*Y\+(?P<BY>\d+).*X=(?P<X>\d+).*Y=(?P<Y>\d+)`)

	clawMachines := []ClawMachine{}

	for _, clawString := range claws {

		match := regex.FindStringSubmatch(strings.Replace(clawString, "\n", "", -1))

		clawMachines = append(clawMachines, NewClawMachine(match))
	}

	sum := int64(0)

	for i, cl := range clawMachines {
		// part1Result := part1(cl)

		sx := (float64(-cl.X)*(float64(cl.BY)/float64(cl.BX)) + float64(cl.Y)) /
			((float64(cl.AY) / float64(cl.AX)) - (float64(cl.BY) / float64(cl.BX)))

		tokens := int64(0)

		if i == 42 {
			fmt.Print("")
		}

		if IsInt(sx) {
			a := int64(math.Round(sx / float64(cl.AX)))
			b := int64(math.Round((float64(cl.X) - sx) / float64(cl.BX)))

			tokens = a*3 + b
		}

		// if part1Result != tokens {
		// 	fmt.Printf("i=%d expected token: %d, got %d\n", i, part1Result, tokens)
		// }

		sum += tokens
	}

	fmt.Printf("Minimum Tokens: %d", sum)
}

func IsInt(f float64) bool {
	intFloat := math.Round(f)

	if f == intFloat {
		return true
	}

	const e = 0.0001
	isInt := f+e > intFloat && f-e < intFloat

	return isInt
}

// func part1(clawMachine ClawMachine) int64 {
// 	results := []ButtonPresses{}

// 	for a := int64(1); a <= 100; a++ {

// 		ax := clawMachine.AX * a
// 		ay := clawMachine.AY * a

// 		if ax > clawMachine.X || ay > clawMachine.Y {
// 			break
// 		}

// 		for b := int64(1); b <= 100; b++ {
// 			x := ax + b*clawMachine.BX

// 			if clawMachine.X == x {
// 				y := ay + b*clawMachine.BY
// 				if clawMachine.Y == y {
// 					results = append(results, ButtonPresses{A: a, B: b})
// 				} else if clawMachine.Y < y {
// 					break
// 				}
// 			} else if clawMachine.X < x {
// 				break
// 			}
// 		}
// 	}

// 	if len(results) > 0 {
// 		tokens := results[0].Tokens()

// 		for i := 1; i < len(results); i++ {
// 			t := results[i].Tokens()
// 			if t < tokens {
// 				tokens = t
// 			}
// 		}

// 		return tokens
// 	}

// 	return 0
// }
