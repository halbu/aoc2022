package main

import (
	"aoc2022/utils"
	"fmt"
)

var sensors, p = [][]int{}, map[int]bool{}

func main() {
	for _, e := range utils.GetData("./day15/day15-input") {
		sensors = append(sensors, utils.GetInts(e))
	}

	for _, e := range sensors {
		sx, sy, bx, by := e[0], e[1], e[2], e[3]
		mht := utils.IAbs(sx-bx) + utils.IAbs(sy-by) // Manhattan distance
		dy := utils.IAbs(sy - 2000000)               // Y-distance from sensor to y-val of interest

		if mht > dy {
			rem := mht - dy
			for x := (sx - rem); x <= (sx + rem); x++ {
				p[x] = true
			}
		}
	}

	for i := 0; i != len(sensors); i++ {
		if sensors[i][3] == 2000000 {
			delete(p, sensors[i][2])
		}
	}

	fmt.Println(len(p))
}
