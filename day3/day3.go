package main

import (
	"aoc2022/utils"
	"fmt"
	"strconv"
)

func intToPriority(i int) int {
	if i >= 97 {
		return i - 96
	} else {
		return i - 38
	}
}

func main() {
	var data = utils.GetData("./day3/day3-input")

	var priorities = []int{}

	for _, e := range data {
		var l int = len(e) / 2
		var first = e[:l]
		var last = e[l:]

		// Find common element, convert to int, then to priority value, and add it
		// to priorities array
		var foundCommonElement bool = false
		for fi, _ := range first {
			for li, _ := range last {
				if first[fi] == last[li] && !foundCommonElement {
					foundCommonElement = true
					var charAsInt int = int(first[fi])
					priorities = append(priorities, intToPriority(charAsInt))
				}
			}
		}
	}

	fmt.Println("Day 3 Part 1 solution: " + strconv.Itoa(utils.Sum(priorities)))

	var newPriorities = []int{}

	// Iterate over the data, in batches of three, isolating commonalities
	var ix int = 0
	for ix < len(data) {
		var c = utils.Commonalities(data[ix], data[ix+1])
		c = utils.Commonalities(c, data[ix+2])
		newPriorities = append(newPriorities, intToPriority(int(c[0])))
		ix += 3
	}

	fmt.Print("Day 3 Part 2 solution: " + strconv.Itoa(utils.Sum(newPriorities)))
}
