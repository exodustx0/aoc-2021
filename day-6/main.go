package main

import (
	"os"
	"strconv"
	"strings"
)

type FishTimers [9]int

func (f *FishTimers) tick(times int) {
	for i := 0; i < times; i++ {
		newFish := f[0]
		for j := 0; j < 8; j++ {
			f[j] = f[j+1]
		}
		f[8] = newFish
		f[6] += newFish
	}
}

func (f *FishTimers) numFishes() int {
	count := 0
	for i := 0; i < 9; i++ {
		count += f[i]
	}
	return count
}

func getInput() (fishTimers *FishTimers) {
	// content, err := os.ReadFile("example.txt")
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err.Error())
	}

	fishTimers = new(FishTimers)
	for _, timerStr := range strings.Split(string(content), ",") {
		timer, err := strconv.Atoi(timerStr)
		if err != nil {
			panic(err.Error())
		}

		fishTimers[timer]++
	}

	return
}

func main() {
	fishTimers := getInput()
	fishTimers.tick(80)
	println("Part one:", fishTimers.numFishes())
	fishTimers.tick(256 - 80)
	println("Part two:", fishTimers.numFishes())
}
