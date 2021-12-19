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

func (c *Crabs) minMaxPositions() (min, max int) {
	min = math.MaxInt
	for _, pos := range *c {
		if pos < min {
			min = pos
		}
		if pos > max {
			max = pos
		}
	}

	return
}

func (c *Crabs) leastFuelNeededToAlign(min, max int, calc func(x int) int) int {
	leastFuelNeeded := math.MaxInt
	for target := min; target <= max; target++ {
		fuelNeeded := 0
		for _, pos := range *c {
			fuelNeeded += calc(abs(target - pos))
		}

		if fuelNeeded < leastFuelNeeded {
			leastFuelNeeded = fuelNeeded
		}
	}

	return leastFuelNeeded
}

func newCrabs(filename string) *Crabs {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}

	var crabs Crabs
	for _, posStr := range strings.Split(string(content), ",") {
		pos, _ := strconv.Atoi(posStr)
		crabs = append(crabs, pos)
	}

	return &crabs
}

func main() {
	for _, filename := range []string{"example.txt", "input.txt"} {
		println(filename)

		crabs := newCrabs(filename)
		min, max := crabs.minMaxPositions()
		println("\tPart one:", crabs.leastFuelNeededToAlign(min, max, func(x int) int { return x }))
		println("\tPart two:", crabs.leastFuelNeededToAlign(min, max, func(x int) int { return (x * (x + 1)) / 2 }))
	}
}
