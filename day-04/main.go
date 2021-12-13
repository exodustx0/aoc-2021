package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Cell struct {
	value  byte
	marked bool
}
type Board struct {
	cells  [5][5]Cell
	marks  byte
	hasWon bool
}

func (b *Board) mark(value byte) {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if b.cells[y][x].value == value {
				b.cells[y][x].marked = true
				b.marks++
				if b.marks >= 5 {
					b.checkWin()
				}
				return
			}
		}
	}
}

func (b *Board) checkWin() {
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if !b.cells[y][x].marked {
				break
			}

			if y == 4 {
				b.hasWon = true
				return
			}
		}
	}
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if !b.cells[y][x].marked {
				break
			}

			if x == 4 {
				b.hasWon = true
				return
			}
		}
	}
}

func (b *Board) unmarkedSum() (sum int) {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if !b.cells[y][x].marked {
				sum += int(b.cells[y][x].value)
			}
		}
	}
	return
}

func getInput(filename string) (calls *[]byte, boards *[]*Board) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	calls = new([]byte)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	for _, valueStr := range strings.Split(scanner.Text(), ",") {
		value, _ := strconv.Atoi(valueStr)
		*calls = append(*calls, byte(value))
	}

	boards = new([]*Board)
	for scanner.Scan() {
		board := Board{}

		for y := 0; y < 5; y++ {
			for x := 0; x < 5; x++ {
				if (x != 0 || y != 0) && !scanner.Scan() {
					panic("Incomplete board")
				}
				value, _ := strconv.Atoi(scanner.Text())
				board.cells[y][x].value = byte(value)
			}
		}

		*boards = append(*boards, &board)
	}

	return
}

func partOne(filename string) {
	calls, boards := getInput(filename)

	for _, value := range *calls {
		for _, board := range *boards {
			board.mark(value)
			if board.hasWon {
				println("\tPart one:", board.unmarkedSum()*int(value))
				return
			}
		}
	}
}

func partTwo(filename string) {
	calls, boards := getInput(filename)

	var lastBoard *Board
	var lastCall int
	for _, board := range *boards {
		for call, value := range *calls {
			board.mark(value)
			if board.hasWon {
				if call > lastCall {
					lastBoard = board
					lastCall = call
				}
				break
			}
		}
	}

	println("\tPart two:", lastBoard.unmarkedSum()*int((*calls)[lastCall]))
}

func main() {
	for _, filename := range []string{"example.txt", "input.txt"} {
		println(filename)
		partOne(filename)
		partTwo(filename)
	}
}
