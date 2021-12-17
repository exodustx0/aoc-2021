package main

import (
	"os"
	"strconv"
	"strings"
)

type Range [2]int

func (r *Range) includes(pos int) bool {
	if pos >= r[0] && pos <= r[1] {
		return true
	}
	return false
}

func newRange(input string) Range {
	input = input[2:]
	ab := strings.Split(input, "..")
	a, _ := strconv.Atoi(ab[0])
	b, _ := strconv.Atoi(ab[1])

	if a < b {
		return Range{a, b}
	} else {
		return Range{b, a}
	}
}

type Target struct {
	x, y Range
}

func (t *Target) includes(x, y int) bool {
	if t.x.includes(x) && t.y.includes(y) {
		return true
	}
	return false
}

func (t *Target) testInitialVelocity(xVelocity, yVelocity int) bool {
	var xPosition, yPosition int

	for {
		xPosition += xVelocity
		yPosition += yVelocity

		if yVelocity < 0 && yPosition < t.y[0] {
			return false
		} else if t.includes(xPosition, yPosition) {
			return true
		}

		if xVelocity > 0 {
			xVelocity--
		} else if xVelocity < 0 {
			xVelocity++
		}
		yVelocity--
	}
}

func (t *Target) findHighestPossiblePosition() int {
	var xInitialStart, yInitial int

	if t.x[1] > 0 {
		xInitialStart = t.x[1]
	} else {
		// This assumes !t.x.includes(0), which is... sorta fine? Both example and input don't have this, so for once I'll refrain from acting on my perfectionism.
		xInitialStart = t.x[0]
	}

	if t.y[1] > 0 {
		yInitial = t.y[1]
	} else {
		// This assumes !t.y.includes(0), which is fine cause the answer would be infinity otherwise.
		yInitial = -t.y[0] - 1
	}

	for yInitial != 0 {
		xInitial := xInitialStart

		for xInitial != 0 {
			if t.testInitialVelocity(xInitial, yInitial) {
				goto found
			}

			if xInitial > 0 {
				xInitial--
			} else {
				xInitial++
			}
		}

		yInitial--
	}

found:
	yPosition := 0
	yVelocity := yInitial
	for yVelocity != 0 {
		yPosition += yVelocity
		yVelocity--
	}

	return yPosition
}

func (t *Target) countPossibleInitialVelocities() (count int) {
	// TODO:

	return
}

func getInput(filename string) *Target {
	inputBuf, err := os.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}

	input := string(inputBuf[13:])
	xy := strings.Split(input, ", ")
	x := newRange(xy[0])
	y := newRange(xy[1])

	return &Target{x, y}
}

func main() {
	for _, filename := range []string{"example.txt", "input.txt"} {
		println(filename)

		target := getInput(filename)
		println("\tPart one:", target.findHighestPossiblePosition())
		println("\tPart two:", target.countPossibleInitialVelocities())
	}
}
