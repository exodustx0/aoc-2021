package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func product(slice []int) int {
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

func (hm *Heightmap) getRiskLevelSum() (sum int) {
	for _, lowPoint := range hm.lowPoints {
		sum += 1 + int(hm.grid[lowPoint.y][lowPoint.x])
	}

	return
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

func getInput(filename string) (heightmap *Heightmap) {
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
			height, _ := strconv.Atoi(string(heightRune))
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
	for _, filename := range []string{"example.txt", "input.txt"} {
		println(filename)

		heightmap := getInput(filename)
		println("\tPart one:", heightmap.getRiskLevelSum())
		println("\tPart two:", product(heightmap.getBasinSizes()[:3]))
	}
}
