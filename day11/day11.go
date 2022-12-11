package main

import (
	"aoc2022/utils"
	"fmt"
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

	for round := 1; round < 21; round++ {
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

				monkeys[i].items[0] = (monkeys[i].items[0] / 3)

				if monkeys[i].items[0]%monkeys[i].test == 0 {
					monkeys[monkeys[i].trueMonkeyIx].items = append(monkeys[monkeys[i].trueMonkeyIx].items, monkeys[i].items[0])
				} else {
					monkeys[monkeys[i].falseMonkeyIx].items = append(monkeys[monkeys[i].falseMonkeyIx].items, monkeys[i].items[0])
				}
				monkeys[i].items = monkeys[i].items[1:]
			}
		}
	}

	utils.IntSort(inspections, true)
	fmt.Println(inspections[0] * inspections[1])
}
