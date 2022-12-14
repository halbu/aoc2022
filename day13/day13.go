package main

import (
	"aoc2022/utils"
	"strconv"
	"unicode"
)

type Packet struct {
	i *int
	l []Packet
}

func isInt(p Packet) bool      { return p.i != nil }
func toList(p Packet) []Packet { return []Packet{{p.i, []Packet{}}} }

func main() {
	data := utils.GetData("./day13/day13-input")

	correctIndices := []int{}
	for i := 0; i < len(data); i += 3 {
		left, right := data[i], data[i+1]
		leftPacket, _ := packetize((left[1 : len(left)-1]))
		rightPacket, _ := packetize((right[1 : len(right)-1]))

		correctness := compare(leftPacket.l, rightPacket.l)
		correctIndices = append(correctIndices, utils.IFloor(correctness, 0)*((i/3)+1))
	}

	utils.Log("Day 13 Part 1 solution: " + strconv.Itoa(utils.Sum(correctIndices)))
}

// Takes a string and converts it to a Packet. Any substring beginning with `[`
// is packetized recursively by this function and the result appended to the
// list element of our Packet. Returns the Packet, and the number of characters
// consumed in the process of converting the string to the Packet.
func packetize(s string) (Packet, int) {
	if s == "" {
		return Packet{nil, []Packet{}}, 0
	} else if string(s[0]) == "," {
		s = s[1:]
	} else if string(s[0]) == "]" {
		return Packet{nil, []Packet{}}, 1
	}

	token := ""
	out := Packet{}

	for i := 0; i < len(s); i++ {
		if string(s[i]) == "[" {
			packet, jumpIndex := packetize(s[i+1:])
			out.l = append(out.l, packet)
			i += jumpIndex
		} else if string(s[i]) == "]" {
			if token != "" {
				iVal := utils.StringToInt(token)
				out.l = append(out.l, Packet{&iVal, []Packet{}})
			}
			return out, i
		} else if unicode.IsDigit(rune(s[i])) {
			token += string(s[i])
		} else if string(s[i]) == "," {
			iVal := utils.StringToInt(token)
			out.l = append(out.l, Packet{&iVal, []Packet{}})
			token = ""
		}
	}
	iVal := utils.StringToInt(token)
	out.l = append(out.l, Packet{&iVal, []Packet{}})

	return out, 0
}

// Compares two lists of packets. Returns 1 (correct order), -1 (incorrect order), 0 (equal)
func compare(l []Packet, r []Packet) int {

	for i := 0; i < utils.Max([]int{len(l), len(r)}); i++ {

		if i >= len(l) || i >= len(r) {
			return utils.ISign(len(r) - len(l))
		}

		left, right := l[i], r[i]

		if isInt(left) && isInt(right) {
			if utils.ISign(*right.i-*left.i) != 0 {
				return utils.ISign(*right.i - *left.i)
			}
		} else {
			result := 0
			if !isInt(left) && !isInt(right) {
				result = compare(left.l, right.l)
			} else {
				if isInt(left) {
					result = compare(toList(left), right.l)
				} else {
					result = compare(left.l, toList(right))
				}
			}
			if result != 0 {
				return result
			}
		}
	}

	return 0
}
