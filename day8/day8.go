package main

import (
	"aoc2022/utils"
	"fmt"
	"strconv"
)

var treeMap [][]int = [][]int{}

func valid(x int, y int) bool {
	return x >= 0 && x < len(treeMap) && y >= 0 && y < len(treeMap[0])
}

func testVisibleFromEdge(x int, y int, dx int, dy int) bool {
	var height = treeMap[x][y]
	var tx = x + dx
	var ty = y + dy
	for valid(tx, ty) {
		if treeMap[tx][ty] >= height {
			return false
		}
		tx += dx
		ty += dy
	}
	return true
}

func calculateViewDistance(x int, y int, dx int, dy int) int {
	var height = treeMap[x][y]
	var dist = 0
	var tx = x + dx
	var ty = y + dy
	for valid(tx, ty) {
		dist++
		if treeMap[tx][ty] >= height {
			break
		}
		tx += dx
		ty += dy
	}
	return dist
}

func main() {
	var data = utils.GetData("./day8/day8-input")

	// Dump our tree height values into a 2d array
	for i := range data {
		treeMap = append(treeMap, []int{})
		for j := range data[i] {
			var element = utils.StringToInt(string(data[i][j]))
			treeMap[i] = append(treeMap[i], element)
		}
	}

	var visibleCount = 0

	for i := 0; i != len(treeMap); i++ {
		for j := 0; j != len(treeMap[0]); j++ {
			var visDirs []bool = []bool{true, true, true, true}
			visDirs[0] = testVisibleFromEdge(i, j, -1, 0)
			visDirs[1] = testVisibleFromEdge(i, j, 1, 0)
			visDirs[2] = testVisibleFromEdge(i, j, 0, -1)
			visDirs[3] = testVisibleFromEdge(i, j, 0, 1)

			if visDirs[0] || visDirs[1] || visDirs[2] || visDirs[3] {
				visibleCount++
			}
		}
	}

	fmt.Println("Day 8 Part 1 solution: " + strconv.Itoa(visibleCount))

	var scenicValues = []int{}

	for i := 0; i != len(treeMap); i++ {
		for j := 0; j != len(treeMap[0]); j++ {
			var viewDistances []int = []int{0, 0, 0, 0}
			viewDistances[0] = calculateViewDistance(i, j, -1, 0)
			viewDistances[1] = calculateViewDistance(i, j, 1, 0)
			viewDistances[2] = calculateViewDistance(i, j, 0, -1)
			viewDistances[3] = calculateViewDistance(i, j, 0, 1)
			scenicValues = append(scenicValues, utils.Product(viewDistances))
		}
	}

	fmt.Println("Day 8 Part 2 solution: " + strconv.Itoa(utils.Max(scenicValues)))
}
