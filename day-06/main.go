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

func (f *FishTimers) countFish() (count int) {
	for i := 0; i < 9; i++ {
		count += f[i]
	}
	return
}

func getInput(filename string) (fishTimers *FishTimers) {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}

	fishTimers = new(FishTimers)
	for _, timerStr := range strings.Split(string(content), ",") {
		timer, _ := strconv.Atoi(timerStr)
		fishTimers[timer]++
	}

	return
}

func main() {
	for _, filename := range []string{"example.txt", "input.txt"} {
		println(filename)

		fishTimers := getInput(filename)
		fishTimers.tick(80)
		println("\tPart one:", fishTimers.countFish())
		fishTimers.tick(256 - 80)
		println("\tPart two:", fishTimers.countFish())
	}
}
