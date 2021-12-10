package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	// filename = "example.txt"
	filename = "input.txt"

	open   = "([{<"
	closed = ")]}>"
)

type NavigationSubsystem [][]rune

func (ns *NavigationSubsystem) parse() (errorScore, middleScore int) {
	autoCompleteScores := make([]int, 0)
lineLoop:
	for _, line := range *ns {
		chunks := make([]rune, 0)
		for _, c := range line {
			switch {
			case strings.ContainsRune(open, c):
				chunks = append(chunks, c)
			case strings.ContainsRune(closed, c):
				index := strings.IndexRune(closed, c)
				if len(chunks) == 0 || strings.IndexRune(open, chunks[len(chunks)-1]) != index {
					// Corrupted
					switch index {
					case 0:
						errorScore += 3
					case 1:
						errorScore += 57
					case 2:
						errorScore += 1197
					case 3:
						errorScore += 25137
					}
					continue lineLoop
				}
				chunks = chunks[:len(chunks)-1]
			default:
				panic(fmt.Sprintf("Illegal character %c", c))
			}
		}

		if len(chunks) != 0 {
			// Incomplete
			var autoCompleteScore int
			for i := len(chunks) - 1; i >= 0; i-- {
				autoCompleteScore *= 5
				autoCompleteScore += strings.IndexRune(open, chunks[i]) + 1
			}
			autoCompleteScores = append(autoCompleteScores, autoCompleteScore)
		}
	}

	sort.Slice(autoCompleteScores, func(i, j int) bool { return autoCompleteScores[i] < autoCompleteScores[j] })
	middleScore = autoCompleteScores[len(autoCompleteScores)/2]

	return
}

func getInput() (lines *NavigationSubsystem) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	lines = new(NavigationSubsystem)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		*lines = append(*lines, []rune(scanner.Text()))
	}

	return
}

func main() {
	lines := getInput()
	errorScore, middleScore := lines.parse()
	println("Part one:", errorScore)
	println("Part two:", middleScore)
}
