package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Fold struct {
	position int
	axis     rune
}

type Paper [][]bool

func (p *Paper) fold(fold Fold) {
	for y, row := range *p {
		for x, cell := range row {
			if !cell {
				continue
			}

			if fold.axis == 'x' && x > fold.position {
				diff := abs(x - fold.position)
				(*p)[y][fold.position-diff] = true
			} else if fold.axis == 'y' && y > fold.position {
				diff := abs(y - fold.position)
				(*p)[fold.position-diff][x] = true
			}
		}
	}

	if fold.axis == 'x' {
		for y := range *p {
			(*p)[y] = (*p)[y][:fold.position]
		}
	} else {
		*p = (*p)[:fold.position]
	}
}

func (p *Paper) countDots() (count int) {
	for _, row := range *p {
		for _, cell := range row {
			if cell {
				count++
			}
		}
	}
	return
}

func (p *Paper) print() {
	for _, row := range *p {
		print("\t")
		for _, cell := range row {
			if cell {
				print("█")
			} else {
				print("░")
			}
		}
		print("\n")
	}
}

type Dot struct {
	x, y int
}

func getInput(filename string) (*Paper, *[]Fold) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	var width, height int
	var folds []Fold
	var dots []Dot
	scanner := bufio.NewScanner(f)
	gettingFolds := false
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			gettingFolds = true
			continue
		}

		if !gettingFolds {
			xy := strings.Split(line, ",")

			x, _ := strconv.Atoi(xy[0])
			y, _ := strconv.Atoi(xy[1])
			dots = append(dots, Dot{x, y})

			if x > width {
				width = x
			}
			if y > height {
				height = y
			}

			continue
		}

		line = line[11:] // Slice off "fold along "
		position, _ := strconv.Atoi(line[2:])
		axis := rune(line[0])
		folds = append(folds, Fold{position, axis})
	}

	width++
	height++

	paper := make(Paper, height)
	for y := 0; y < height; y++ {
		paper[y] = make([]bool, width)
	}

	for _, dot := range dots {
		paper[dot.y][dot.x] = true
	}

	return &paper, &folds
}

func main() {
	for _, filename := range []string{"example.txt", "input.txt"} {
		println(filename)

		paper, folds := getInput(filename)
		paper.fold((*folds)[0])
		println("\tPart one:", paper.countDots())

		for _, fold := range (*folds)[1:] {
			paper.fold(fold)
		}
		println("\tPart two:")
		paper.print()
	}
}
