package main

import (
	"bufio"
	"os"
	"strconv"
)

const (
	// filename = "example.txt"
	filename = "input.txt"
)

func main() {
	f, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	var simpleDepth, depth, pos, aim int
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		op := scanner.Text()

		if !scanner.Scan() {
			panic("Invalid amount of tokens!")
		}
		a, _ := strconv.Atoi(scanner.Text())

		switch op {
		case "forward":
			pos += a
			depth += a * aim
		case "down":
			simpleDepth += a
			aim += a
		case "up":
			simpleDepth -= a
			aim -= a
		}
	}

	println("Part one:", simpleDepth*pos)
	println("Part two:", depth*pos)
}
