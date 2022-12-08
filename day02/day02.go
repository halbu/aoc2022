package main

import (
	"aoc2022/utils"
	"fmt"
	"strconv"
)

func main() {
	var data = utils.GetData("./day02/day02-input")

	// Let's hardcode the outcomes. Seems simplest
	outcomes := map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}

	var points int = 0
	for _, e := range data {
		var result int = outcomes[e]
		points += result
	}

	fmt.Println("Day 2 Part 1 solution: " + strconv.Itoa(points))

	// And let's hardcode the outcomes again
	newStrategyOutcomes := map[string]int{
		"A X": 3,
		"A Y": 4,
		"A Z": 8,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 2,
		"C Y": 6,
		"C Z": 7,
	}

	var newStrategyPoints int = 0
	for _, e := range data {
		var result int = newStrategyOutcomes[e]
		newStrategyPoints += result
	}

	fmt.Println("Day 2 Part 2 solution: " + strconv.Itoa(newStrategyPoints))
}
