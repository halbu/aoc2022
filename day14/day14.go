package main

import (
	"aoc2022/utils"
	"strconv"
	"strings"
)

var grid, floor = [][]rune{}, 0

func main() {
	data := utils.GetData("./day14/day14-input")

	for i := 0; i != 1000; i++ {
		grid = append(grid, []rune{})
		for j := 0; j != 1000; j++ {
			grid[i] = append(grid[i], '.')
		}
	}

	for _, e := range data {
		path := strings.Split(e, " -> ")

		for i := 0; i < len(path)-1; i++ {
			origin, target := strings.Split(path[i], ","), strings.Split(path[i+1], ",")
			x, y := utils.StringToInt(origin[0]), utils.StringToInt(origin[1])
			tx, ty := utils.StringToInt(target[0]), utils.StringToInt(target[1])

			floor = utils.Max([]int{floor, y + 2, ty + 2})

			for x != tx || y != ty {
				grid[x][y] = '#'
				x += utils.ISign(tx - x)
				y += utils.ISign(ty - y)
			}
			grid[x][y] = '#'
		}
	}

	hitVoid := false

	for i := 0; i != int(^uint(0)>>1); i++ {
		settled := false
		sx, sy := 500, 0

		for !settled {
			if sy == floor-1 {
				settled = true
				if !hitVoid {
					utils.Log("Day 14 Part 1 solution: " + strconv.Itoa(i))
					hitVoid = true
				}
			} else if grid[sx][sy+1] == '.' {
				sy++
			} else if grid[sx-1][sy+1] == '.' {
				sy++
				sx--
			} else if grid[sx+1][sy+1] == '.' {
				sy++
				sx++
			} else {
				settled = true
			}
		}
		grid[sx][sy] = 'o'

		if sx == 500 && sy == 0 {
			utils.Log("Day 14 Part 2 solution: " + strconv.Itoa(i+1))
			return
		}
	}
}
