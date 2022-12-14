package main

import (
	"aoc2022/utils"
	"sort"
	"strconv"
	"unicode"
)

type Packet struct {
	i   *int
	l   []Packet
	tag bool
}

func isInt(p Packet) bool      { return p.i != nil }
func toList(p Packet) []Packet { return []Packet{{p.i, []Packet{}, false}} }

func main() {
	data := utils.GetData("./day13/day13-input")
	packets := []Packet{}

	correctIndices := []int{}
	for i := 0; i < len(data); i += 3 {
		left, right := data[i], data[i+1]
		leftPacket, _ := packetize((left[1 : len(left)-1]))
		rightPacket, _ := packetize((right[1 : len(right)-1]))
		packets = append(packets, leftPacket)
		packets = append(packets, rightPacket)

		correctness := compare(leftPacket.l, rightPacket.l)
		correctIndices = append(correctIndices, utils.IFloor(correctness, 0)*((i/3)+1))
	}

	utils.Log("Day 13 Part 1 solution: " + strconv.Itoa(utils.Sum(correctIndices)))

	// Part 2 is going to be solved via the extremely naive method of just adding
	// a boolean `tag` field to the Packet struct, and then manually tagging the
	// two divider packets that we want to locate in the sorted array.
	dividerPacketOne, _ := packetize("[[[2]]]")
	dividerPacketTwo, _ := packetize("[[[6]]]")
	dividerPacketOne.tag = true
	dividerPacketTwo.tag = true
	packets = append(packets, []Packet{dividerPacketOne, dividerPacketTwo}...)

	sort.Slice(packets, func(i, j int) bool {
		return compare(packets[i].l, packets[j].l) == 1
	})

	dividerIndices := []int{}
	for i := 0; i < len(packets); i++ {
		if packets[i].tag {
			dividerIndices = append(dividerIndices, i+1)
		}
	}

	utils.Log("Day 13 Part 2 solution: " + strconv.Itoa(utils.Product((dividerIndices))))
}

// Takes a string and converts it to a Packet. Any substring beginning with `[`
// is packetized recursively by this function and the result appended to the
// list element of our Packet. Returns the Packet, and the number of characters
// consumed in the process of converting the string to the Packet.
func packetize(s string) (Packet, int) {
	if s == "" {
		return Packet{nil, []Packet{}, false}, 0
	} else if string(s[0]) == "," {
		s = s[1:]
	} else if string(s[0]) == "]" {
		return Packet{nil, []Packet{}, false}, 1
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
				out.l = append(out.l, Packet{&iVal, []Packet{}, false})
			}
			return out, i + 1
		} else if unicode.IsDigit(rune(s[i])) {
			token += string(s[i])
		} else if string(s[i]) == "," {
			iVal := utils.StringToInt(token)
			out.l = append(out.l, Packet{&iVal, []Packet{}, false})
			token = ""
		}
	}
	iVal := utils.StringToInt(token)
	out.l = append(out.l, Packet{&iVal, []Packet{}, false})

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
