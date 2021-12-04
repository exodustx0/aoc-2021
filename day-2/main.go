package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getInput() *os.File {
	// f, err := os.Open("example.txt")
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}

	return f
}

func partOne() {
	f := getInput()
	defer f.Close()

	var depth, pos int
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
		case "down":
			depth += a
		case "up":
			depth -= a
		default:
			panic(fmt.Sprintf("Invalid op %s!", op))
		}
	}

	fmt.Println("Part one:", depth*pos)
}

func partTwo() {
	f := getInput()
	defer f.Close()

	var depth, pos, aim int
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
			aim += a
		case "up":
			aim -= a
		default:
			panic(fmt.Sprintf("Invalid op %s!", op))
		}
	}

	fmt.Println("Part two:", depth*pos)
}

func main() {
	partOne()
	partTwo()
}
