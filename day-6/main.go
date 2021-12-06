package main

import (
	"os"
	"strconv"
	"strings"
)

type FishTimers [9]int

func (f *FishTimers) tick() {
	newFish := f[0]
	for i := 0; i < 8; i++ {
		f[i] = f[i+1]
	}
	f[8] = newFish
	f[6] += newFish
}

func (f *FishTimers) count() int {
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

func partOne() {
	fishTimers := getInput()
	for i := 0; i < 80; i++ {
		fishTimers.tick()
	}

	println("Part one:", fishTimers.count())
}

func partTwo() {
	fishTimers := getInput()
	for i := 0; i < 256; i++ {
		fishTimers.tick()
	}

	println("Part two:", fishTimers.count())
}

func main() {
	partOne()
	partTwo()
}
