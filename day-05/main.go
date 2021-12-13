package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func (p *Point) parse(input string) {
	xy := strings.Split(input, ",")
	if len(xy) != 2 {
		panic("Malformed coordinate input")
	}

	x, err := strconv.Atoi(xy[0])
	if err != nil {
		panic(err.Error())
	}
	p.x = x

	y, err := strconv.Atoi(xy[1])
	if err != nil {
		panic(err.Error())
	}
	p.y = y
}

type Line struct {
	from, to *Point
}

func (l *Line) maxX() int {
	if l.from.x > l.to.x {
		return l.from.x
	} else {
		return l.to.x
	}
}

func (l *Line) maxY() int {
	if l.from.y > l.to.y {
		return l.from.y
	} else {
		return l.to.y
	}
}

type Grid [][]byte

func (g *Grid) init(width, height int) {
	*g = make([][]byte, height)
	for y := 0; y < height; y++ {
		(*g)[y] = make([]byte, width)
	}
}

func (g *Grid) drawLines(lines *[]Line) {
	for _, line := range *lines {
		y := line.from.y
		x := line.from.x
		if y == line.to.y {
			// Horizontal
			for {
				(*g)[y][x]++

				if x == line.to.x {
					break
				} else if x < line.to.x {
					x++
				} else {
					x--
				}
			}
		} else if x == line.to.x {
			// Vertical
			for {
				(*g)[y][x]++

				if y == line.to.y {
					break
				} else if y < line.to.y {
					y++
				} else {
					y--
				}
			}
		} else {
			// Diagonal (45 degrees)
			for {
				(*g)[y][x]++

				if y == line.to.y && x == line.to.x {
					break
				}

				if y < line.to.y {
					y++
				} else {
					y--
				}

				if x < line.to.x {
					x++
				} else {
					x--
				}
			}
		}
	}
}

func (g *Grid) getOverlapCount() int {
	var overlap int
	for _, row := range *g {
		for _, cell := range row {
			if cell > 1 {
				overlap++
			}
		}
	}

	return overlap
}

// Used for debugging
func (g *Grid) print() {
	for _, row := range *g {
		for _, cell := range row {
			if cell == 0 {
				print(".")
			} else {
				print(cell)
			}
		}
		print("\n")
	}
}

func getInput(filename string, includeDiagonals bool) (lines *[]Line, width, height int) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	lines = new([]Line)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		p1 := Point{}
		p1.parse(scanner.Text())

		if !scanner.Scan() || scanner.Text() != "->" || !scanner.Scan() {
			panic("Malformed input")
		}

		p2 := Point{}
		p2.parse(scanner.Text())

		if includeDiagonals || p1.x == p2.x || p1.y == p2.y {
			var line Line
			if p1.x > p2.x || p1.y > p2.y {
				line.from = &p2
				line.to = &p1
			} else {
				line.from = &p1
				line.to = &p2
			}

			*lines = append(*lines, line)

			if line.maxX() > width {
				width = line.maxX()
			}
			if line.maxY() > height {
				height = line.maxY()
			}
		}
	}

	width++
	height++
	return
}

func partOne(filename string) {
	lines, width, height := getInput(filename, false)

	var grid Grid
	grid.init(width, height)
	grid.drawLines(lines)
	println("\tPart one:", grid.getOverlapCount())
}

func partTwo(filename string) {
	lines, width, height := getInput(filename, true)

	var grid Grid
	grid.init(width, height)
	grid.drawLines(lines)
	println("\tPart two:", grid.getOverlapCount())
}

func main() {
	for _, filename := range []string{"example.txt", "input.txt"} {
		println(filename)
		partOne(filename)
		partTwo(filename)
	}
}
