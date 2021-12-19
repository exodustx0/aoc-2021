package main

import (
	"bufio"
	"math"
	"os"
	"strings"
)

type Digit []rune

func (d *Digit) equals(d2 *Digit) bool {
	length := len(*d)
	if length != len(*d2) {
		return false
	}
	if length != 5 && length != 6 {
		return true
	}
	for _, segment := range *d {
		if !d2.includes(segment) {
			return false
		}
	}
	return true
}

func (d *Digit) includes(segment rune) bool {
	for _, s := range *d {
		if s == segment {
			return true
		}
	}
	return false
}

func (d *Digit) fits(d2 *Digit) bool {
	for _, s := range *d2 {
		if !d.includes(s) {
			return false
		}
	}

	return true
}

type Display struct {
	digits [10]*Digit
	output [4]*Digit
}

func (d *Display) decode() {
	unsortedDigits := d.digits

	var one, four, seven, eight *Digit
	for _, digit := range unsortedDigits {
		switch len(*digit) {
		case 2:
			one = digit
		case 3:
			seven = digit
		case 4:
			four = digit
		case 7:
			eight = digit
		}
	}

	var zero, six, nine *Digit
	for _, digit := range unsortedDigits {
		if len(*digit) != 6 {
			continue
		}

		if !digit.fits(one) {
			six = digit
		} else if !digit.fits(four) {
			zero = digit
		} else {
			nine = digit
		}
	}

	var topRight rune
	for _, s := range "abcdefg" {
		if !six.includes(s) {
			topRight = s
			break
		}
	}

	var two, three, five *Digit
	for _, digit := range unsortedDigits {
		if len(*digit) != 5 {
			continue
		}

		if digit.fits(one) {
			three = digit
		} else if digit.includes(topRight) {
			two = digit
		} else {
			five = digit
		}
	}

	d.digits = [10]*Digit{zero, one, two, three, four, five, six, seven, eight, nine}
}

func (d *Display) value() (value int) {
	for i, digit := range d.output {
		var number int
		for j, digit2 := range d.digits {
			if digit == digit2 {
				number = j
				break
			}
		}

		value += number * int(math.Pow(10, float64(3-i)))
	}

	return
}

type Displays []Display

func (d *Displays) countEasyDigits() (count int) {
	for _, display := range *d {
		for _, digitPtr := range display.output {
			if digitPtr == display.digits[1] ||
				digitPtr == display.digits[4] ||
				digitPtr == display.digits[7] ||
				digitPtr == display.digits[8] {
				count++
			}
		}
	}

	return
}

func (d *Displays) countValues() (count int) {
	for _, display := range *d {
		count += display.value()
	}

	return
}

func newDisplays(filename string) *Displays {
	f, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	var displays Displays
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		var digits [10]*Digit
		for i, digitStr := range line[:10] {
			digit := Digit(digitStr)
			digits[i] = &digit
		}

		var output [4]*Digit
		for i, digitStr := range line[11:] {
			digit := Digit(digitStr)
			for _, digitPtr := range digits {
				if digit.equals(digitPtr) {
					output[i] = digitPtr
				}
			}
		}

		display := Display{digits, output}
		display.decode()
		displays = append(displays, display)
	}

	return &displays
}

func main() {
	for _, filename := range []string{"example.txt", "input.txt"} {
		println(filename)

		displays := newDisplays(filename)
		println("\tPart one:", displays.countEasyDigits())
		println("\tPart two:", displays.countValues())
	}
}
