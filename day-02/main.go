package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	for _, filename := range []string{"example.txt", "input.txt"} {
		println(filename)

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
			scanner.Scan()
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

		println("\tPart one:", simpleDepth*pos)
		println("\tPart two:", depth*pos)
	}
}
