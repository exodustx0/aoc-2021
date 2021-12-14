package main

import (
	"bufio"
	"math"
	"os"
)

type Pair [2]byte

type InsertionRule struct {
	pair Pair
	to   byte
}

type Polymer struct {
	start, end byte
	pairs      map[Pair]int
	rules      []InsertionRule
}

func (p *Polymer) step(times int) {
	for i := 0; i < times; i++ {
		oldPairs := p.pairs
		p.pairs = make(map[Pair]int)
		for pair, count := range oldPairs {
			for _, rule := range p.rules {
				if pair != rule.pair {
					continue
				}

				p.pairs[Pair{pair[0], rule.to}] += count
				p.pairs[Pair{rule.to, pair[1]}] += count
				break
			}
		}
	}
}

func (p *Polymer) getElementCountRange() int {
	elementCounts := make(map[byte]int)
	for pair, pairCount := range p.pairs {
		for _, element := range pair {
			elementCounts[element] += pairCount
		}
	}

	elementCounts[p.start]--
	elementCounts[p.end]--

	for element := range elementCounts {
		elementCounts[element] /= 2
	}

	elementCounts[p.start]++
	elementCounts[p.end]++

	min := math.MaxInt
	max := 0
	for _, count := range elementCounts {
		if count < min {
			min = count
		}
		if count > max {
			max = count
		}
	}

	return max - min
}

func getInput(filename string) (polymer *Polymer) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	chain := scanner.Text()
	chainLength := len(chain)

	polymer = &Polymer{
		start: chain[0],
		end:   chain[chainLength-1],
		pairs: make(map[Pair]int),
	}
	for i := 0; i < chainLength-1; i++ {
		polymer.pairs[Pair{chain[i], chain[i+1]}]++
	}

	for scanner.Scan() {
		pairStr := scanner.Text()
		scanner.Scan() // Arrow
		scanner.Scan()
		pair := Pair{pairStr[0], pairStr[1]}
		to := scanner.Text()[0]

		polymer.rules = append(polymer.rules, InsertionRule{pair, to})
	}

	return
}

func main() {
	for _, filename := range []string{"example.txt", "input.txt"} {
		println(filename)

		polymer := getInput(filename)
		polymer.step(10)
		println("\tPart one:", polymer.getElementCountRange())
		polymer.step(40 - 10)
		println("\tPart two:", polymer.getElementCountRange())
	}
}
