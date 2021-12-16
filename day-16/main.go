package main

import (
	"fmt"
	"math"
	"os"
)

func sum(values []uint) (sum uint) {
	for _, value := range values {
		sum += value
	}

	return
}

func product(values []uint) (product uint) {
	product = 1
	for _, value := range values {
		product *= value
	}

	return
}

func min(values []uint) (min uint) {
	min = math.MaxUint
	for _, value := range values {
		if value < min {
			min = value
		}
	}

	return
}

func max(values []uint) (max uint) {
	for _, value := range values {
		if value > max {
			max = value
		}
	}

	return
}

type TransmissionParser struct {
	transmission []byte

	bytePosition, bitPosition uint
}

func (t *TransmissionParser) readBit() bool {
	b := t.transmission[t.bytePosition]
	bitmask := byte(1 << (7 - t.bitPosition))

	t.bitPosition++
	if t.bitPosition == 8 {
		t.bitPosition = 0
		t.bytePosition++
	}

	return (b & bitmask) != 0
}

func (t *TransmissionParser) readBits(numBits uint) (value uint) {
	for numBits != 0 {
		value <<= 1
		if t.readBit() {
			value |= 1
		}
		numBits--
	}

	return
}

type PacketType uint

const (
	PT_SUM PacketType = iota
	PT_PRODUCT
	PT_MIN
	PT_MAX
	PT_LITERAL
	PT_GREATER
	PT_LESS
	PT_EQUAL
)

type Packet struct {
	version uint
	t       PacketType

	value      uint
	subPackets []*Packet
}

func (p *Packet) getVersionSum() uint {
	sum := p.version

	for _, subPacket := range p.subPackets {
		sum += subPacket.getVersionSum()
	}

	return sum
}

func (p *Packet) computeValue() uint {
	if p.t == PT_LITERAL {
		return p.value
	}

	var values []uint
	for _, subPacket := range p.subPackets {
		values = append(values, subPacket.computeValue())
	}

	switch p.t {
	case PT_SUM:
		return sum(values)
	case PT_PRODUCT:
		return product(values)
	case PT_MIN:
		return min(values)
	case PT_MAX:
		return max(values)
	case PT_GREATER:
		if values[0] > values[1] {
			return 1
		}
	case PT_LESS:
		if values[0] < values[1] {
			return 1
		}
	case PT_EQUAL:
		if values[0] == values[1] {
			return 1
		}
	}

	return 0
}

func (t *TransmissionParser) readPacket() *Packet {
	packet := Packet{
		version: t.readBits(3),
		t:       PacketType(t.readBits(3)),
	}

	switch packet.t {
	case PT_LITERAL:
		packet.value = t.readPacketLiteral()
	default:
		packet.subPackets = t.readPacketOperator()
	}

	return &packet
}

func (t *TransmissionParser) readPacketOperator() []*Packet {
	var subPackets []*Packet
	if t.readBit() {
		numSubPackets := int(t.readBits(11))
		subPackets = make([]*Packet, numSubPackets)
		for i := 0; i < numSubPackets; i++ {
			subPackets[i] = t.readPacket()
		}
	} else {
		numBits := t.readBits(15)
		for numBits != 0 {
			startByte := t.bytePosition
			startBit := t.bitPosition

			subPackets = append(subPackets, t.readPacket())

			numBits -= 8*(t.bytePosition-startByte) + (t.bitPosition - startBit)
		}
	}

	return subPackets
}

func (t *TransmissionParser) readPacketLiteral() uint {
	var value uint
	for {
		last := !t.readBit()
		value <<= 4
		value |= t.readBits(4)
		if last {
			break
		}
	}

	return value
}

func getInput(transmission []byte) *Packet {
	parser := TransmissionParser{
		transmission: make([]byte, len(transmission)/2),
	}
	for _, hexChar := range transmission {
		var nybble byte
		if hexChar >= '0' && hexChar <= '9' {
			nybble = hexChar - '0'
		} else if hexChar >= 'A' && hexChar <= 'F' {
			nybble = hexChar - 'A' + 10
		}

		if parser.bitPosition == 0 {
			parser.bitPosition += 4
			parser.transmission[parser.bytePosition] = nybble << 4
		} else {
			parser.bitPosition = 0
			parser.transmission[parser.bytePosition] |= nybble
			parser.bytePosition++
		}
	}

	parser.bytePosition = 0

	return parser.readPacket()
}

type Expectation struct {
	input  string
	output uint
}

func main() {
	println("Version tests:")
	for _, e := range []Expectation{
		{"8A004A801A8002F478", 16},
		{"620080001611562C8802118E34", 12},
		{"C0015000016115A2E0802F182340", 23},
		{"A0016C880162017C3686B18A3D4780", 31},
	} {
		fmt.Printf("\t%s: %t\n", e.input, getInput([]byte(e.input)).getVersionSum() == e.output)
	}

	println("Value tests:")
	for _, e := range []Expectation{
		{"C200B40A82", 3},
		{"04005AC33890", 54},
		{"880086C3E88112", 7},
		{"CE00C43D881120", 9},
		{"D8005AC2A8F0", 1},
		{"F600BC2D8F", 0},
		{"9C005AC2F8F0", 0},
		{"9C0141080250320F1802104A08", 1},
	} {
		fmt.Printf("\t%s: %t\n", e.input, getInput([]byte(e.input)).computeValue() == e.output)
	}

	println("input.txt:")
	transmission, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err.Error())
	}

	packet := getInput(transmission)
	println("\tPart one:", packet.getVersionSum())
	println("\tPart two:", packet.computeValue())
}
