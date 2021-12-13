package main

import (
	"bufio"
	"math"
	"os"
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

func (d *Digit) invert() Digit {
	var inverted Digit
	for _, s := range "abcdefg" {
		if !d.includes(s) {
			inverted = append(inverted, s)
		}
	}
	return inverted
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

	var two, three, five *Digit
	topRight := six.invert()[0]
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

func (d *Display) getValue() int {
	var value int
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

	return value
}

type Displays []Display

func (d *Displays) countEasyDigits() int {
	var count int
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

	return count
}

func (d *Displays) countValues() int {
	var count int
	for _, display := range *d {
		count += display.getValue()
	}

	return count
}

func getInput(filename string) (displays *Displays) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	displays = new(Displays)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		digits := [10]*Digit{}
		output := [4]*Digit{}
		for i := 0; i < 15; i++ {
			if i != 0 && !scanner.Scan() {
				panic("Incomplete sequence of digits")
			}

			if i == 10 {
				if scanner.Text() != "|" {
					panic("Malformed input")
				}

				continue
			}

			digit := Digit(scanner.Text())
			if i < 10 {
				digits[i] = &digit
			} else {
				for _, digitPtr := range digits {
					if digit.equals(digitPtr) {
						output[i-11] = digitPtr
						break
					}
				}
			}
		}

		display := Display{digits, output}
		display.decode()
		*displays = append(*displays, display)
	}

	return
}

func main() {
	for _, filename := range []string{"example.txt", "input.txt"} {
		println(filename)

		displays := getInput(filename)
		println("\tPart one:", displays.countEasyDigits())
		println("\tPart two:", displays.countValues())
	}
}
