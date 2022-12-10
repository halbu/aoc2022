package main

import (
	"aoc2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var data = utils.GetData("./day10/day10-input")

	buffer := [][]int{} // element format: [ value to add, cycles required for execution ]

	for _, e := range data {
		if strings.Split(e, " ")[0] != "noop" {
			buffer = append(buffer, []int{utils.StringToInt(strings.Split(e, " ")[1]), 2})
		} else {
			buffer = append(buffer, []int{0, 1})
		}
	}

	cycle, register, total := 1, 1, 0

	for len(buffer) > 0 {
		if (cycle-20)%40 == 0 {
			total += (cycle * register)
		}

		buffer[0][1]--
		if buffer[0][1] == 0 {
			register += buffer[0][0]
			buffer = buffer[1:]
		}
		cycle++
	}

	fmt.Println("Day 10 Part 1 solution: " + strconv.Itoa(total))
}
