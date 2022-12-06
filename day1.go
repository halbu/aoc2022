package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getData(s string) []string {
	file, _ := os.Open(s)

	scanner := bufio.NewScanner(file)
	var output []string

	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	return output
}

func stringToInt(s string) int {
	integer, _ := strconv.Atoi(s)
	return integer
}

func max(arr []int) int {
	var max = 0
	for _, e := range arr {
		if e > max {
			max = e
		}
	}
	return max
}

func main() {
	var output = getData("./day1-input")
	var values []int
	var elfIndex int = 0

	for _, s := range output {
		if s == "" {
			elfIndex++
		} else {
			if len(values) < (elfIndex + 1) {
				values = append(values, 0)
			}

			values[elfIndex] += stringToInt(s)
		}
	}

	fmt.Println("Day 1 Part 1 solution: " + strconv.Itoa(max(values)))

	maxes := []int{0, 0, 0}

	for _, e := range values {
		var foundSmallerElement bool = false
		for i := range maxes {
			if e > maxes[i] && !foundSmallerElement {
				// If this value is bigger than any of the values in the maxes array...
				// Insert this value into the array at the appropriate point and pop the smallest value off the end
				maxes = append(maxes[:i+1], maxes[i:]...)
				maxes[i] = e
				maxes = maxes[:3]
				foundSmallerElement = true
			}
		}
	}

	var result int = 0
	for _, e := range maxes {
		result += e
	}

	fmt.Println("Day 1 Part 2 solution: " + strconv.Itoa(result))
}
