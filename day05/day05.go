package main

import (
	"aoc2022/utils"
	"fmt"
	"strings"
)

// I am not going to try and parse that input from the file. Sorry!
func getStacks() [][]rune {
	return [][]rune{
		{'B', 'Z', 'T'},
		{'V', 'H', 'T', 'D', 'N'},
		{'B', 'F', 'M', 'D'},
		{'T', 'J', 'G', 'W', 'V', 'Q', 'L'},
		{'W', 'D', 'G', 'O', 'V', 'F', 'Q', 'M'},
		{'V', 'Z', 'Q', 'G', 'H', 'F', 'S'},
		{'Z', 'S', 'N', 'R', 'L', 'T', 'C', 'W'},
		{'Z', 'H', 'W', 'D', 'J', 'N', 'R', 'M'},
		{'M', 'Q', 'L', 'F', 'D', 'S'},
	}
}

func main() {
	var data = utils.GetData("./day05/day05-input")

	var stacks = getStacks()

	for _, e := range data[10:] {
		var move = utils.StringToInt(strings.Split(strings.Split(e, " from ")[0], "move ")[1])
		var from = utils.StringToInt(strings.Split(strings.Split(e, " from ")[1], " to ")[0]) - 1
		var to = utils.StringToInt(strings.Split(strings.Split(e, " from ")[1], " to ")[1]) - 1

		for i := 0; i != move; i++ {
			// Get the last element of the from stack
			var fromStack = stacks[from]
			var fromElement = fromStack[len(fromStack)-1]

			// Push it onto the to stack
			stacks[to] = append(stacks[to], fromElement)

			// Delete the last element of the from stack
			stacks[from] = stacks[from][:len(stacks[from])-1]
		}
	}

	solutionPt1 := ""
	for _, e := range stacks {
		solutionPt1 += string(e[len(e)-1])
	}

	fmt.Println("Day 5 Part 1 solution: " + solutionPt1)

	// Reset the stacks
	stacks = getStacks()

	for _, e := range data[10:] {
		var move = utils.StringToInt(strings.Split(strings.Split(e, " from ")[0], "move ")[1])
		var from = utils.StringToInt(strings.Split(strings.Split(e, " from ")[1], " to ")[0]) - 1
		var to = utils.StringToInt(strings.Split(strings.Split(e, " from ")[1], " to ")[1]) - 1

		// Get the appropriate number of elements of the from stack
		var fromStack = stacks[from]
		var fromElements = fromStack[len(fromStack)-move:]

		// Push them onto the to stack, preserving the original ordering
		stacks[to] = append(stacks[to], fromElements...)

		// Delete the appropriate number of elements from the from stack
		stacks[from] = stacks[from][:len(stacks[from])-move]
	}

	solutionPt2 := ""
	for _, e := range stacks {
		solutionPt2 += string(e[len(e)-1])
	}

	fmt.Println("Day 5 Part 2 solution: " + solutionPt2)
}
