package main

import (
	"aoc2022/utils"
	"fmt"
	"strconv"
)

func main() {
	var data = utils.GetData("./day1/day1-input")
	var values []int
	var elfIndex int = 0

	for _, s := range data {
		if s == "" {
			elfIndex++
		} else {
			if len(values) < (elfIndex + 1) {
				values = append(values, 0)
			}

			values[elfIndex] += utils.StringToInt(s)
		}
	}

	fmt.Println("Day 1 Part 1 solution: " + strconv.Itoa(utils.Max(values)))

	maxes := []int{0, 0, 0}

	for _, e := range values {
		var foundSmallerElement bool = false
		for i := range maxes {
			// If this value is bigger than any of the values in the maxes array -
			// Insert this value into the array at the appropriate point and pop the
			// smallest value off the end. Then stop checking the subsequent values in
			// order that we don't insert it again
			if e > maxes[i] && !foundSmallerElement {
				maxes = append(maxes[:i+1], maxes[i:]...)
				maxes[i] = e
				maxes = maxes[:3]
				foundSmallerElement = true
			}
		}
	}

	fmt.Println("Day 1 Part 2 solution: " + strconv.Itoa(utils.Sum(maxes)))
}
