package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

const (
	// filename = "example.txt"
	filename = "input.txt"
)

func multiplySlice(slice []int) int {
	product := 1
	for _, factor := range slice {
		product *= factor
	}

	return product
}

type Coordinate struct {
	x, y int
}

type Basin []Coordinate

func (b *Basin) includes(c2 Coordinate) bool {
	for _, c := range *b {
		if c == c2 {
			return true
		}
	}

	return false
}

type Heightmap struct {
	grid          [][]byte
	width, height int
	lowPoints     []Coordinate
}

func (hm *Heightmap) getRiskLevelSum() int {
	var riskLevelSum int
	for _, lowPoint := range hm.lowPoints {
		riskLevelSum += int(1 + hm.grid[lowPoint.y][lowPoint.x])
	}

	return riskLevelSum
}

func (hm *Heightmap) getBasinNeighbors(c Coordinate) []Coordinate {
	var neighbors []Coordinate
	x := c.x
	y := c.y
	if x != 0 && hm.grid[y][x-1] != 9 {
		neighbors = append(neighbors, Coordinate{x - 1, y})
	}
	if x != hm.width-1 && hm.grid[y][x+1] != 9 {
		neighbors = append(neighbors, Coordinate{x + 1, y})
	}
	if y != 0 && hm.grid[y-1][x] != 9 {
		neighbors = append(neighbors, Coordinate{x, y - 1})
	}
	if y != hm.height-1 && hm.grid[y+1][x] != 9 {
		neighbors = append(neighbors, Coordinate{x, y + 1})
	}

	return neighbors
}

func (hm *Heightmap) getBasin(c Coordinate, basin *Basin) {
	for _, neighbor := range hm.getBasinNeighbors(c) {
		if !basin.includes(neighbor) {
			*basin = append(*basin, neighbor)
			hm.getBasin(neighbor, basin)
		}
	}
}

func (hm *Heightmap) getBasinSizes() []int {
	var sizes []int
	for _, lowPoint := range hm.lowPoints {
		basin := Basin{lowPoint}
		hm.getBasin(lowPoint, &basin)
		sizes = append(sizes, len(basin))
	}

	sort.Slice(sizes, func(i, j int) bool { return sizes[i] > sizes[j] })

	return sizes
}

func getInput() (heightmap *Heightmap) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	heightmap = new(Heightmap)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var row []byte
		for _, heightRune := range scanner.Text() {
			height, err := strconv.Atoi(string(heightRune))
			if err != nil {
				panic(err.Error())
			}
			row = append(row, byte(height))
		}
		heightmap.grid = append(heightmap.grid, row)
	}

	heightmap.width = len(heightmap.grid[0])
	heightmap.height = len(heightmap.grid)
	for y, row := range heightmap.grid {
		for x, height := range row {
			if x != 0 && heightmap.grid[y][x-1] <= height {
				continue
			}
			if x != heightmap.width-1 && heightmap.grid[y][x+1] <= height {
				continue
			}
			if y != 0 && heightmap.grid[y-1][x] <= height {
				continue
			}
			if y != heightmap.height-1 && heightmap.grid[y+1][x] <= height {
				continue
			}

			heightmap.lowPoints = append(heightmap.lowPoints, Coordinate{x, y})
		}
	}

	return
}

func main() {
	heightmap := getInput()
	println("Part one:", heightmap.getRiskLevelSum())
	println("Part two:", multiplySlice(heightmap.getBasinSizes()[:3]))
}
