package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type ValueKind byte

const (
	VK_NUMBER ValueKind = iota
	VK_PAIR
)

// OK, unions would've been nice. Seriously, what's the deal here?
type Value struct {
	kind   ValueKind
	number byte
	pair   [2]*Value

	parent *Value
}

func (v *Value) copy() *Value {
	newValue := &Value{
		kind:   v.kind,
		parent: v.parent,
	}

	if v.kind == VK_PAIR {
		newValue.pair[0] = v.pair[0].copy()
		newValue.pair[1] = v.pair[1].copy()
		newValue.pair[0].parent = newValue
		newValue.pair[1].parent = newValue
	} else {
		newValue.number = v.number
	}

	return newValue
}

type FindDirection byte

const (
	FD_PREVIOUS FindDirection = iota
	FD_NEXT
)

func (v *Value) findNumber(last *Value, direction FindDirection) *Value {
	if v.kind == VK_NUMBER {
		return v
	}

	var first, second *Value
	if direction == FD_NEXT {
		first = v.pair[0]
		second = v.pair[1]
	} else {
		first = v.pair[1]
		second = v.pair[0]
	}

	switch last {
	case v.parent:
		return first.findNumber(v, direction)
	case first:
		return second.findNumber(v, direction)
	case second:
		if v.parent != nil {
			return v.parent.findNumber(v, direction)
		}
	}

	return nil
}

type ReduceMode byte

const (
	RM_PAIR ReduceMode = iota
	RM_NUMBER
)

func (v *Value) reduce(nestLevel byte, mode ReduceMode) bool {
startOver:
	switch v.kind {
	case VK_NUMBER:
		if mode == RM_NUMBER && v.number > 9 {
			roundDown := v.number / 2
			roundUp := v.number - roundDown
			v.kind = VK_PAIR
			v.pair = [2]*Value{
				{kind: VK_NUMBER, number: roundDown, parent: v},
				{kind: VK_NUMBER, number: roundUp, parent: v},
			}
			return true
		}
	case VK_PAIR:
		if mode == RM_PAIR && nestLevel == 4 {
			if previousNumber := v.parent.findNumber(v, FD_PREVIOUS); previousNumber != nil {
				previousNumber.number += v.pair[0].number
			}
			if nextNumber := v.parent.findNumber(v, FD_NEXT); nextNumber != nil {
				nextNumber.number += v.pair[1].number
			}

			v.kind = VK_NUMBER
			v.number = 0

			return true
		}

		switch {
		case v.pair[0].reduce(nestLevel+1, mode):
			fallthrough
		case v.pair[1].reduce(nestLevel+1, mode):
			if v.parent == nil {
				mode = RM_PAIR
				goto startOver
			}

			return true
		}
	}

	if nestLevel == 0 && mode == RM_PAIR {
		mode = RM_NUMBER
		goto startOver
	}

	return false
}

func (v *Value) magnitude() uint {
	if v.kind == VK_NUMBER {
		return uint(v.number)
	}

	return 3*v.pair[0].magnitude() + 2*v.pair[1].magnitude()
}

// Used for debugging
func (v *Value) toString() string {
	if v.kind == VK_NUMBER {
		return strconv.Itoa(int(v.number))
	}

	return fmt.Sprintf("[%s,%s]", v.pair[0].toString(), v.pair[1].toString())
}

type ValuesReader struct {
	input    string
	position int
}

func (ir *ValuesReader) isPair() bool {
	if ir.input[ir.position] == '[' {
		ir.position++
		return true
	}

	return false
}

func (ir *ValuesReader) readNumber() byte {
	number, _ := strconv.Atoi(ir.input[ir.position : ir.position+1])
	ir.position++
	return byte(number)
}

func (hr *ValuesReader) newValue(parent *Value) *Value {
	value := new(Value)
	if hr.isPair() {
		value.kind = VK_PAIR
		value.pair[0] = hr.newValue(value)
		hr.position++ // Comma
		value.pair[1] = hr.newValue(value)
		hr.position++ // Close bracket
	} else {
		value.kind = VK_NUMBER
		value.number = hr.readNumber()
	}

	value.parent = parent
	return value
}

type Values []*Value

func (v *Values) sum() *Value {
	sum := (*v)[0].copy()
	for _, second := range (*v)[1:] {
		first := sum
		second = second.copy()
		sum = &Value{
			kind: VK_PAIR,
			pair: [2]*Value{first, second},
		}

		first.parent = sum
		second.parent = sum
		sum.reduce(0, RM_PAIR)
	}

	return sum
}

func (v *Values) maxMagnitude() (max uint) {
	for _, a := range *v {
		for _, b := range *v {
			if a == b {
				continue
			}

			values := Values{a, b}
			mag := values.sum().magnitude()
			if mag > max {
				max = mag
			}
		}
	}

	return
}

func newValues(filename string) *Values {
	f, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	var values Values
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		reader := ValuesReader{input: scanner.Text()}
		values = append(values, reader.newValue(nil))
	}

	return &values
}

func main() {
	for _, filename := range []string{"example.txt", "input.txt"} {
		println(filename)

		values := newValues(filename)
		println("\tPart one:", values.sum().magnitude())
		println("\tPart two:", values.maxMagnitude())
	}
}
