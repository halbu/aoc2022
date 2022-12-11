package main

import (
	"aoc2022/utils"
	"fmt"
	"strconv"
	"strings"
)

type Monkey struct {
	items         []int
	op            string
	test          int
	trueMonkeyIx  int
	falseMonkeyIx int
}

func main() {
	fmt.Println("Day 11 part 1 solution: " + strconv.Itoa(processMonkeyBusiness(false, 20)))
	fmt.Println("Day 11 part 2 solution: " + strconv.Itoa(processMonkeyBusiness(true, 10000)))
}

func processMonkeyBusiness(pt2 bool, rounds int) int {
	var data = utils.GetData("./day11/day11-input")

	monkeys, inspections := []Monkey{}, []int{}

	for i := 0; i < len(data); i += 7 {
		inspections = append(inspections, 0)
		var monkeyItemArray = []int{}
		for _, e := range strings.Split(strings.TrimLeft(data[i+1], "Starting items: "), ", ") {
			monkeyItemArray = append(monkeyItemArray, utils.StringToInt(e))
		}
		monkeys = append(monkeys, Monkey{
			items:         monkeyItemArray,
			op:            strings.Split(data[i+2], "Operation: new = ")[1],
			test:          utils.StringToInt(strings.Split(data[i+3], "Test: divisible by ")[1]),
			trueMonkeyIx:  utils.StringToInt(strings.Split(data[i+4], "throw to monkey ")[1]),
			falseMonkeyIx: utils.StringToInt(strings.Split(data[i+5], "throw to monkey ")[1]),
		})
	}

	for round := 0; round < rounds; round++ {
		for i := 0; i < len(monkeys); i++ {
			for len(monkeys[i].items) > 0 {
				inspections[i]++

				operator, operand := strings.Split(monkeys[i].op, " ")[1], strings.Split(monkeys[i].op, " ")[2]
				iOperand := 0
				if operand == "old" {
					iOperand = monkeys[i].items[0]
				} else {
					iOperand = utils.StringToInt(strings.Split(monkeys[i].op, " ")[2])
				}

				if operator == "+" {
					monkeys[i].items[0] = monkeys[i].items[0] + iOperand
				} else if operator == "*" {
					monkeys[i].items[0] = monkeys[i].items[0] * iOperand
				}

				if pt2 {
					monkeys[i].items[0] = monkeys[i].items[0] % 9699690
				} else {
					monkeys[i].items[0] = (monkeys[i].items[0] / 3)
				}

				if monkeys[i].items[0]%monkeys[i].test == 0 {
					monkeys[monkeys[i].trueMonkeyIx].items = append(monkeys[monkeys[i].trueMonkeyIx].items, monkeys[i].items[0])
				} else {
					monkeys[monkeys[i].falseMonkeyIx].items = append(monkeys[monkeys[i].falseMonkeyIx].items, monkeys[i].items[0])
				}
				monkeys[i].items = monkeys[i].items[1:]
			}
		}
	}

	return (utils.Product(utils.IntSort(inspections, true)[0:2]))
}
