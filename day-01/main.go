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

		var a, b, c, d, inc, sumInc int
		scanner := bufio.NewScanner(f)
		for i := 0; scanner.Scan(); i++ {
			a, _ = strconv.Atoi(scanner.Text())

			if i != 0 {
				if a > b {
					inc++
				}
				if i > 2 {
					if a > d {
						sumInc++
					}
				}
			}

			d = c
			c = b
			b = a
		}

		println("\tPart one:", inc)
		println("\tPart two:", sumInc)
	}
}
