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

func (hm *Heightmap) riskLevelSum() (sum int) {
	for _, lowPoint := range hm.lowPoints {
		sum += 1 + int(hm.grid[lowPoint.y][lowPoint.x])
	}

	return
}

func (hm *Heightmap) basinNeighbours(c Coordinate) []Coordinate {
	var neighbours []Coordinate
	x := c.x
	y := c.y
	if x != 0 && hm.grid[y][x-1] != 9 {
		neighbours = append(neighbours, Coordinate{x - 1, y})
	}
	if x != hm.width-1 && hm.grid[y][x+1] != 9 {
		neighbours = append(neighbours, Coordinate{x + 1, y})
	}
	if y != 0 && hm.grid[y-1][x] != 9 {
		neighbours = append(neighbours, Coordinate{x, y - 1})
	}
	if y != hm.height-1 && hm.grid[y+1][x] != 9 {
		neighbours = append(neighbours, Coordinate{x, y + 1})
	}

	return neighbours
}

func (hm *Heightmap) addToBasin(c Coordinate, basin *Basin) {
	for _, neighbour := range hm.basinNeighbours(c) {
		if !basin.includes(neighbour) {
			*basin = append(*basin, neighbour)
			hm.addToBasin(neighbour, basin)
		}
	}
}

func (hm *Heightmap) basinSizes() []int {
	var sizes []int
	for _, lowPoint := range hm.lowPoints {
		basin := Basin{lowPoint}
		hm.addToBasin(lowPoint, &basin)
		sizes = append(sizes, len(basin))
	}

	sort.Slice(sizes, func(i, j int) bool { return sizes[i] > sizes[j] })

	return sizes
}

func newHeightmap(filename string) *Heightmap {
	f, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	var heightmap Heightmap
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

	return &heightmap
}

func main() {
	for _, filename := range []string{"example.txt", "input.txt"} {
		println(filename)

		heightmap := newHeightmap(filename)
		println("\tPart one:", heightmap.riskLevelSum())
		println("\tPart two:", product(heightmap.basinSizes()[:3]))
	}
}
