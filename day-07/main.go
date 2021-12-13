package main

import (
	"math"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Crabs []int

func (c *Crabs) leastFuelNeededToAlign(min, max int, f func(x int) int) int {
	leastFuelNeeded := math.MaxInt
	for target := min; target <= max; target++ {
		fuelNeeded := 0
		for _, pos := range *c {
			fuelNeeded += f(abs(target - pos))
		}

		if fuelNeeded < leastFuelNeeded {
			leastFuelNeeded = fuelNeeded
		}
	}

	return leastFuelNeeded
}

func getInput(filename string) (crabs *Crabs, min, max int) {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}

	crabs = new(Crabs)
	for _, posStr := range strings.Split(string(content), ",") {
		pos, err := strconv.Atoi(posStr)
		if err != nil {
			panic(err.Error())
		}

		*crabs = append(*crabs, pos)
	}

	min = math.MaxInt
	max = 0
	for _, pos := range *crabs {
		if pos < min {
			min = pos
		}
		if pos > max {
			max = pos
		}
	}

	return
}

func main() {
	for _, filename := range []string{"example.txt", "input.txt"} {
		println(filename)

		crabs, min, max := getInput(filename)

		println("\tPart one:", crabs.leastFuelNeededToAlign(min, max, func(x int) int {
			return x
		}))

		println("\tPart two:", crabs.leastFuelNeededToAlign(min, max, func(x int) int {
			return (x * (x + 1)) / 2
		}))
	}
}
