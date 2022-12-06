package utils

import (
	"bufio"
	"os"
	"strconv"
)

func GetData(s string) []string {
	file, _ := os.Open(s)
	scanner := bufio.NewScanner(file)
	var output []string

	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	return output
}

func StringToInt(s string) int {
	integer, _ := strconv.Atoi(s)
	return integer
}

/*
*  Convenience functions for num array manipulation that I suspect I will be
*  needing a lot
 */
func Max(arr []int) int {
	var max = 0
	for _, e := range arr {
		if e > max {
			max = e
		}
	}
	return max
}

func Min(arr []int) int {
	var min = int(^uint(0) >> 1)
	for _, e := range arr {
		if e < min {
			min = e
		}
	}
	return min
}

func Sum(arr []int) int {
	var sum int = 0
	for _, e := range arr {
		sum += e
	}
	return sum
}
