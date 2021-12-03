package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getInput() *os.File {
	// f, err = os.Open("example.txt")
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}

	return f
}

func partOne() {
	f := getInput()
	defer f.Close()

	var a, b int
	inc := 0
	first := true
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		a, _ = strconv.Atoi(scanner.Text())
		if !first {
			if a > b {
				inc++
			}
		} else {
			first = false
		}

		b = a
	}

	fmt.Println("Part one:", inc)
}

func partTwo() {
	f := getInput()
	defer f.Close()

	var a, b, c, d int
	i := 0
	inc := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		a, _ = strconv.Atoi(scanner.Text())
		if i >= 3 {
			if a > d {
				inc++
			}
		}

		d = c
		c = b
		b = a
		i++
	}

	fmt.Println("Part two:", inc)
}

func main() {
	partOne()
	partTwo()
}
