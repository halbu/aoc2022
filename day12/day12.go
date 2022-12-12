package main

import (
	"aoc2022/utils"
	"fmt"
	"sort"
	"strconv"
)

type Node struct {
	x    int
	y    int
	cost int
}

var E, g = []int{}, [][]rune{}

func main() {
	fmt.Println("Day 12 Part 1 solution: " + strconv.Itoa(getShortestPath(false)))
	fmt.Println("Day 12 Part 2 solution: " + strconv.Itoa(getShortestPath(true)))
}

func getShortestPath(startFromAnyLowercaseA bool) int {
	var data = utils.GetData("./day12/day12-input")

	startLocations := [][]int{}
	g = [][]rune{}

	for x := 0; x != len(data[0]); x++ {
		g = append(g, []rune{})
		for y := 0; y != len(data); y++ {
			if rune(data[y][x]) == 83 {
				if !startFromAnyLowercaseA {
					startLocations = append(startLocations, []int{x, y})
				}
				g[x] = append(g[x], rune('a'))
			} else if rune(data[y][x]) == 69 {
				E = []int{x, y}
				g[x] = append(g[x], rune('z'))
			} else {
				g[x] = append(g[x], rune(data[y][x]))
			}
			if rune(data[y][x]) == 97 && startFromAnyLowercaseA {
				startLocations = append(startLocations, []int{x, y})
			}
		}
	}

	shortestPaths := []int{}
	for _, s := range startLocations {
		shortestPaths = append(shortestPaths, findPathFrom(s[0], s[1], E[0], E[1]))
	}
	return utils.Min(utils.Remove(shortestPaths, -1))
}

func findPathFrom(ox int, oy int, tx int, ty int) int {
	open, closed, current := []Node{{x: ox, y: oy, cost: 0}}, map[string]bool{}, Node{}

	for {
		if len(open) == 0 {
			return -1
		}
		sort.Slice(open, func(i, j int) bool { return open[i].cost > open[j].cost })
		current, open = open[len(open)-1], open[:len(open)-1]
		closed[strconv.Itoa(current.x)+":"+strconv.Itoa(current.y)] = true

		if current.x == tx && current.y == ty {
			return current.cost
		}

		for _, m := range viableMovesFrom(current.x, current.y) {
			if !closed[strconv.Itoa(m[0])+":"+strconv.Itoa(m[1])] {
				var foundInOpenList = false
				for _, f := range open {
					if f.x == m[0] && f.y == m[1] {
						foundInOpenList = true
						f.cost = utils.Min([]int{current.cost + 1, f.cost})
					}
				}
				if !foundInOpenList {
					open = append(open, Node{x: m[0], y: m[1], cost: current.cost + 1})
				}
			}
		}
	}
}

func validCell(x int, y int) bool { return x >= 0 && x < len(g) && y >= 0 && y < len(g[0]) }

func viableMovesFrom(x int, y int) [][]int {
	out := [][]int{}
	for _, e := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		if validCell(x+e[0], y+e[1]) && g[x][y]+1 >= g[x+e[0]][y+e[1]] {
			out = append(out, []int{x + e[0], y + e[1]})
		}
	}
	return out
}
