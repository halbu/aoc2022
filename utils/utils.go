package utils

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func GetData(s string) []string {
	file, _ := os.Open(s)
	scanner := bufio.NewScanner(file)
	output := []string{}

	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	return output
}

func Log(s string) {
	fmt.Println(s)
}

func ALog(arr []string) {
	for i := 0; i != len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func IAbs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func IFloor(i int, f int) int {
	if i < f {
		return f
	}
	return i
}

func ISign(i int) int {
	if i > 0 {
		return 1
	} else if i < 0 {
		return -1
	} else {
		return 0
	}
}

// Convenience functions for num array manipulation that I suspect I will be
// needing a lot

func Max(arr []int) int {
	max := 0
	for _, e := range arr {
		if e > max {
			max = e
		}
	}
	return max
}

func Min(arr []int) int {
	min := int(^uint(0) >> 1)
	for _, e := range arr {
		if e < min {
			min = e
		}
	}
	return min
}

func Sum(arr []int) int {
	sum := 0
	for _, e := range arr {
		sum += e
	}
	return sum
}

func Product(arr []int) int {
	prod := 0
	for i, e := range arr {
		if i == 0 {
			prod = e
		} else {
			prod *= e
		}
	}
	return prod
}

func IntSort(arr []int, desc bool) []int {
	if desc {
		sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	} else {
		sort.Ints(arr)
	}
	return arr
}

func Remove(arr []int, target int) []int {
	out := []int{}
	for _, e := range arr {
		if e != target {
			out = append(out, e)
		}
	}
	return out
}

func Count[T int | string](arr []T, v T) int {
	count := 0
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
	ix := 0
	l := len(substr)
	found := false

	for !found && ix < len(str) {
		if str[ix:ix+l] == substr {
			found = true
		} else {
			ix++
		}
	}

	if found {
		return str[:ix] + str[ix+l:]
	}
	return str
}

func StringToInt(s string) int {
	integer, _ := strconv.Atoi(s)
	return integer
}

func RemoveDuplicates[T int | string](arr []T) []T {
	keys := make(map[T]bool)
	out := []T{}
	for _, e := range arr {
		if _, v := keys[e]; !v {
			keys[e] = true
			out = append(out, e)
		}
	}
	return out
}
