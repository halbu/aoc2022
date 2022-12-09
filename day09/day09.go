package main

import (
	"aoc2022/utils"
	"fmt"
	"strconv"
	"strings"
)

var dirs map[string][]int = map[string][]int{
	"R": {1, 0}, "L": {-1, 0}, "U": {0, -1}, "D": {0, 1},
}

func isMoveRequired(head []int, tail []int) bool {
	return (utils.IAbs(tail[0]-head[0]) > 1 || utils.IAbs(tail[1]-head[1]) > 1)
}

func getMoveDir(head []int, tail []int) []int {
	var returnDir = []int{0, 0}
	if tail[0] < head[0] {
		returnDir[0] = 1
	} else if tail[0] > head[0] {
		returnDir[0] = -1
	}
	if tail[1] < head[1] {
		returnDir[1] = 1
	} else if tail[1] > head[1] {
		returnDir[1] = -1
	}
	return returnDir
}

func main() {
	var data = utils.GetData("./day09/day09-input")
	var head = []int{0, 0}
	var tail = []int{0, 0}
	var locs = []string{"0, 0"}

	for _, e := range data {
		var dir = strings.Split(e, " ")[0]
		var dist = utils.StringToInt(strings.Split(e, " ")[1])
		var totalHeadMovement = []int{dirs[dir][0], dirs[dir][1]}
		head[0] += (totalHeadMovement[0] * dist)
		head[1] += (totalHeadMovement[1] * dist)

		for isMoveRequired(head, tail) {
			var tailMovDir = getMoveDir(head, tail)
			tail[0] += tailMovDir[0]
			tail[1] += tailMovDir[1]
			locs = append(locs, strconv.Itoa(tail[0])+", "+strconv.Itoa(tail[1]))
		}
	}

	fmt.Println("Day 9 Part 1 solution: " + strconv.Itoa(len(utils.RemoveDuplicates(locs))))

	var rope = [][]int{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	locs = []string{"0, 0"}

	for _, e := range data {
		var dir = strings.Split(e, " ")[0]
		var dist = utils.StringToInt(strings.Split(e, " ")[1])

		// Need to move one step at a time now rather than jumping `head` directly
		// to its destination as this will influence how the subsequent knots move
		for m := 0; m != dist; m++ {
			rope[0][0] += dirs[dir][0]
			rope[0][1] += dirs[dir][1]
			for i := 0; i < len(rope)-1; i++ {
				for isMoveRequired(rope[i], rope[i+1]) {
					var tailMovDir = getMoveDir(rope[i], rope[i+1])
					rope[i+1][0] += tailMovDir[0]
					rope[i+1][1] += tailMovDir[1]
				}
			}
			locs = append(locs, strconv.Itoa(rope[9][0])+", "+strconv.Itoa(rope[9][1]))
		}
	}

	fmt.Println("Day 9 Part 2 solution: " + strconv.Itoa(len(utils.RemoveDuplicates(locs))))
}
