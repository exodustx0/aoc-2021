package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var bits []int
	var numValues, numBits int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if numValues == 0 {
			numBits = len(scanner.Text())
			bits = make([]int, numBits)
		}
		for i, bit := range scanner.Text() {
			if bit == '1' {
				bits[i]++
			}
		}
		numValues++
	}

	var gamma, epsilon int
	half := numValues / 2
	for i, count := range bits {
		if count > half {
			gamma |= 1 << (numBits - i - 1)
		} else if count < half {
			epsilon |= 1 << (numBits - i - 1)
		}
	}

	fmt.Println("Part one:", gamma*epsilon)
}

func partTwo() {
	f := getInput()
	defer f.Close()

	var values []int
	var numValues, numBits int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if numValues == 0 {
			numBits = len(scanner.Text())
		}

		value, err := strconv.ParseUint(scanner.Text(), 2, 0)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		values = append(values, int(value))
		numValues++
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})

	var oxygen, co2 []int
	mask := 1 << (numBits - 1)
	// The loop for the most-signficant bit will get answers for both values, so it's done separate from the rest of the loops
	for i, value := range values {
		if (value & mask) == 0 {
			continue
		}

		if i > numValues-i {
			oxygen = values[:i]
			co2 = values[i:]
		} else {
			co2 = values[:i]
			oxygen = values[i:]
		}
		break
	}

	for mask >>= 1; mask > 0; mask >>= 1 {
		if len(oxygen) > 1 {
			for i, value := range oxygen {
				if (value & mask) == 0 {
					continue
				}

				if i > len(oxygen)-i {
					oxygen = oxygen[:i]
				} else {
					oxygen = oxygen[i:]
				}
				break
			}
		}
		if len(co2) > 1 {
			for i, value := range co2 {
				if (value & mask) == 0 {
					continue
				}

				if i == 0 || i > len(co2)-i {
					co2 = co2[i:]
				} else {
					co2 = co2[:i]
				}
				break
			}
		}
	}

	fmt.Println("Part two:", oxygen[0]*co2[0])
}

func main() {
	partOne()
	partTwo()
}
