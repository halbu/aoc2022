package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

// Convenience functions for num array manipulation that I suspect I will be
// needing a lot

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

func Count(arr []int, v int) int {
	var count int = 0
	for _, e := range arr {
		if e == v {
			count++
		}
	}
	return count
}

// String handling stuff

// Return all common characters shared between two strings
func Commonalities(arr1 string, arr2 string) string {
	output := ""
	for _, e1 := range arr1 {
		for _, e2 := range arr2 {
			if e1 == e2 {
				if !strings.ContainsRune(output, e1) {
					output = output + string(e1)
				}
			}
		}
	}
	return output
}

func StartsWith(str string, substr string) bool {
	if len(str) >= len(substr) {
		if str[:len(substr)] == substr {
			return true
		}
	}
	return false
}

func Remainder(str string, substr string) string {
	var ix = 0
	var l = len(substr)
	var found bool = false

	for !found && ix < len(str) {
		if str[ix:ix+l] == substr {
			found = true
		} else {
			ix++
		}
	}

	if found {
		var newStr string = str[:ix] + str[ix+l:]
		return newStr
	} else {
		return str
	}
}
