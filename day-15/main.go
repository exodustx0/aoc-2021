package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"time"
)

type Node struct {
	x, y, risk int
	neighbours []*Node

	global  float64
	local   int
	visited bool
	parent  *Node
}

type Grid [][]*Node

// Implements the A* algorithm... poorly.
func (g *Grid) lowestTotalRisk() (totalRisk int) {
	for _, row := range *g {
		for _, node := range row {
			node.global = math.MaxFloat64
			node.local = math.MaxInt
			node.visited = false
			node.parent = nil
		}
	}

	height := len(*g)
	width := height
	start := (*g)[0][0]
	end := (*g)[height-1][width-1]
	heuristic := func(n *Node) float64 {
		return 1
	}

	start.local = 0
	start.global = heuristic(start)

	var current *Node
	testList := []*Node{start}
	for len(testList) != 0 {
		sort.Slice(testList, func(i, j int) bool { return testList[i].global < testList[j].global })

		current = testList[0]
		current.visited = true

		for _, neighbour := range current.neighbours {
			if !neighbour.visited {
				testList = append(testList, neighbour)
			}

			newLocal := current.local + neighbour.risk
			if newLocal < neighbour.local {
				neighbour.parent = current
				neighbour.global = float64(newLocal) + heuristic(neighbour)
				neighbour.local = newLocal
			}
		}

		var i int
		for ; i < len(testList); i++ {
			if !testList[i].visited {
				break
			}
		}

		testList = testList[i:]
	}

	current = end
	for current.parent != nil {
		totalRisk += current.risk
		current = current.parent
	}

	return
}

func (g *Grid) expand() {
	tileHeight := len(*g)
	tileWidth := tileHeight
	width := tileWidth * 5
	height := tileHeight * 5

	for y := 0; y < height; y++ {
		var row []*Node
		if y < tileHeight {
			row = (*g)[y]
		}

		for x := 0; x < width; x++ {
			if x < tileWidth && y < tileHeight {
				continue
			}

			var previousTileNode *Node
			if y < tileHeight {
				previousTileNode = row[x-tileWidth]
			} else {
				previousTileNode = (*g)[y-tileHeight][x]
			}

			node := &Node{
				risk: (previousTileNode.risk % 9) + 1,
				x:    x,
				y:    y,
			}

			if x > 0 {
				node.neighbours = append(node.neighbours, row[x-1])
				row[x-1].neighbours = append(row[x-1].neighbours, node)
			}
			if y > 0 {
				node.neighbours = append(node.neighbours, (*g)[y-1][x])
				(*g)[y-1][x].neighbours = append((*g)[y-1][x].neighbours, node)
			}

			row = append(row, node)
		}

		if y < tileHeight {
			(*g)[y] = row
		} else {
			(*g) = append((*g), row)
		}
	}
}

func (g *Grid) visitedPercentage() float64 {
	var visited int
	for _, row := range *g {
		for _, node := range row {
			if node.visited {
				visited++
			}
		}
	}

	height := len(*g)
	width := height
	return 100 * float64(visited/(width*height))
}

func newGrid(filename string) *Grid {
	f, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	var grid Grid
	var lastRow []*Node
	scanner := bufio.NewScanner(f)
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		lineLength := len(line)

		var lastNode *Node
		row := make([]*Node, lineLength)
		for x, risk := range line {
			node := &Node{
				risk: int(risk - '0'),
				x:    x,
				y:    y,
			}

			if x > 0 {
				node.neighbours = append(node.neighbours, lastNode)
				lastNode.neighbours = append(lastNode.neighbours, node)
			}
			if y > 0 {
				node.neighbours = append(node.neighbours, lastRow[x])
				lastRow[x].neighbours = append(lastRow[x].neighbours, node)
			}

			row[x] = node
			lastNode = node
		}
		grid = append(grid, row)
		lastRow = row
	}

	return &grid
}

func main() {
	for _, filename := range []string{"example.txt", "input.txt"} {
		println(filename)

		start := time.Now()
		grid := newGrid(filename)
		println("\tPart one:", grid.lowestTotalRisk())
		grid.expand()
		println("\tPart two:", grid.lowestTotalRisk())
		fmt.Printf("\tVisited: %f%%\n", grid.visitedPercentage())
		fmt.Printf("\tDuration: %fs\n", time.Since(start).Seconds())
	}
}
