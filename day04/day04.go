package main

import (
	"aoc2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var data = utils.GetData("./day04/day04-input")

	var containmentCount = 0
	for _, e := range data {
		var assignmentPair = strings.Split(e, ",")
		// Get two separate assignments as string tuples
		var s1 = strings.Split(assignmentPair[0], "-")
		var s2 = strings.Split(assignmentPair[1], "-")
		// Map them to int tuples
		var a1 = [2]int{utils.StringToInt(s1[0]), utils.StringToInt(s1[1])}
		var a2 = [2]int{utils.StringToInt(s2[0]), utils.StringToInt(s2[1])}
		// Test if assignment `a1` fully contains `a2` or vice versa
		if a1[0] <= a2[0] && a1[1] >= a2[1] || a2[0] <= a1[0] && a2[1] >= a1[1] {
			containmentCount++
		}
	}

	fmt.Println("Day 4 Part 1 solution: " + strconv.Itoa(containmentCount))

	// If I wasn't lazy, or I was more concerned with DRY, I'd extract this out to
	// a more generic function that iterates over the assignments and that accepts
	// a `testFunction` parameter which it then tests against each assignment pair
	var overlapCount = 0
	for _, e := range data {
		var assignmentPair = strings.Split(e, ",")
		// Get two separate assignments as string tuples
		var s1 = strings.Split(assignmentPair[0], "-")
		var s2 = strings.Split(assignmentPair[1], "-")
		// Map them to int tuples
		var a1 = [2]int{utils.StringToInt(s1[0]), utils.StringToInt(s1[1])}
		var a2 = [2]int{utils.StringToInt(s2[0]), utils.StringToInt(s2[1])}
		// Test if assignment `a1` overlaps `a2` or vice versa
		if a1[0] <= a2[0] && a1[1] >= a2[0] || a2[0] <= a1[0] && a2[1] >= a1[0] {
			overlapCount++
		}
	}

	fmt.Println("Day 4 Part 2 solution: " + strconv.Itoa(overlapCount))
}
