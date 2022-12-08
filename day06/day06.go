package main

import (
	"aoc2022/utils"
	"fmt"
	"strconv"
)

func indexOfFirstDistinctNChars(n int, data string) int {
	for i := 0; i != len(data)-n; i++ {
		var str = data[i : i+n]
		// If our commonality function reports n common characters between two
		// instances of the string, then every character in the packet is distinct
		if len(utils.Commonalities(str, str)) == n {
			return i + n
		}
	}
	return -1
}

func main() {
	var data = utils.GetData("./day06/day06-input")[0]
	fmt.Println("Day 6 Part 1 solution: " + strconv.Itoa(indexOfFirstDistinctNChars(4, data)))
	fmt.Println("Day 6 Part 2 solution: " + strconv.Itoa(indexOfFirstDistinctNChars(14, data)))
}
