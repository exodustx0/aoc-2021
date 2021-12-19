package main

import (
	"bufio"
	"os"
)

type Octopus struct {
	start, level byte
	flashes      int
	neighbors    []*Octopus
}

func (o *Octopus) flash() {
	if o.level < 10 {
		o.level++
	}
	if o.level != 10 {
		return
	}

	o.level++
	o.flashes++
	for _, neighbor := range o.neighbors {
		neighbor.flash()
	}
}

type Grid [10][10]*Octopus

func (g *Grid) forEachOctopus(f func(octopus *Octopus)) {
	for _, row := range *g {
		for _, octopus := range row {
			f(octopus)
		}
	}
}

func (g *Grid) reset() {
	g.forEachOctopus(func(octopus *Octopus) { octopus.level = octopus.start })
}

func (g *Grid) step(times int) {
	for i := 0; i < times; i++ {
		g.forEachOctopus(func(octopus *Octopus) { octopus.flash() })
		g.forEachOctopus(func(octopus *Octopus) {
			if octopus.level > 9 {
				octopus.level = 0
			}
		})
	}
}

func (g *Grid) countFlashes() (count int) {
	g.forEachOctopus(func(octopus *Octopus) { count += octopus.flashes })

	return
}

// Used for debugging
func (g *Grid) print() {
	for _, row := range *g {
		for _, octopus := range row {
			print(octopus.level)
		}
		print("\n")
	}
}

func newGrid(filename string) *Grid {
	f, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	var y byte
	var grid Grid
	scanner := bufio.NewScanner(f)
	for ; scanner.Scan(); y++ {
		for x, level := range scanner.Text() {
			grid[y][x] = &Octopus{start: byte(level - '0')}
		}
	}

	for y, row := range grid {
		for x, octopus := range row {
			if x != 0 {
				octopus.neighbors = append(octopus.neighbors, grid[y][x-1])
				if y != 0 {
					octopus.neighbors = append(octopus.neighbors, grid[y-1][x-1])
				}
			}
			if y != 0 {
				octopus.neighbors = append(octopus.neighbors, grid[y-1][x])
				if x != 9 {
					octopus.neighbors = append(octopus.neighbors, grid[y-1][x+1])
				}
			}
			if x != 9 {
				octopus.neighbors = append(octopus.neighbors, grid[y][x+1])
				if y != 9 {
					octopus.neighbors = append(octopus.neighbors, grid[y+1][x+1])
				}
			}
			if y != 9 {
				octopus.neighbors = append(octopus.neighbors, grid[y+1][x])
				if x != 0 {
					octopus.neighbors = append(octopus.neighbors, grid[y+1][x-1])
				}
			}
		}
	}

	return &grid
}

func main() {
	for _, filename := range []string{"example.txt", "input.txt"} {
		println(filename)

		grid := newGrid(filename)
		grid.reset()
		grid.step(100)
		println("\tPart one:", grid.countFlashes())

		var i int
		grid.reset()
	loop:
		i++
		grid.step(1)
		for _, row := range grid {
			for _, octopus := range row {
				if octopus.level != 0 {
					goto loop
				}
			}
		}

		println("\tPart two:", i)
	}
}
