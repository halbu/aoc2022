package main

import (
	"aoc2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var data = utils.GetData("./day10/day10-input")

	stack := [][]int{} // element format: [ value to add, cycles required for execution ]
	crt := []string{"", "", "", "", "", ""}

	for _, e := range data {
		if strings.Split(e, " ")[0] != "noop" {
			stack = append(stack, []int{utils.StringToInt(strings.Split(e, " ")[1]), 2})
		} else {
			stack = append(stack, []int{0, 1})
		}
	}

	cycle, register, total, x, y := 1, 1, 0, 0, 0

	for len(stack) > 0 {
		if register >= x-1 && register <= x+1 {
			crt[y] = crt[y] + "█"
		} else {
			crt[y] = crt[y] + " "
		}

		if (cycle-20)%40 == 0 {
			total += (cycle * register)
		}

		stack[0][1]--
		if stack[0][1] == 0 {
			register += stack[0][0]
			stack = stack[1:]
		}
		cycle++

		x++
		if x > 39 {
			x -= 40
			y++
		}
	}

	fmt.Println("Day 10 Part 1 solution: " + strconv.Itoa(total))

	fmt.Println("Day 10 Part 2 solution: ")
	for i := 0; i < 6; i++ {
		fmt.Println(crt[i])
	}
}
