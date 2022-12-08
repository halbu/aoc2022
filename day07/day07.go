package main

import (
	"aoc2022/utils"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Node struct {
	parent *Node
	files  []int
	dirs   []*Node
}

var dirList = []int{}

func findAllDirsWithCondition(n *Node, val int, cond string) int {
	var tot = 0
	for _, e := range n.files {
		tot += e
	}

	for _, e := range n.dirs {
		tot += findAllDirsWithCondition(e, val, cond)
	}

	if (cond == "max" && tot <= val) || (cond == "min" && tot >= val) || cond == "none" {
		dirList = append(dirList, tot) // Push any dir that meets the conditions into `dirList`
	}
	return tot // Return the total filesize of this directory and all its subdirectories
}

func main() {
	var data = utils.GetData("./day07/day07-input")

	var currentDir *Node
	var root = Node{parent: nil, files: []int{}, dirs: []*Node{}}
	currentDir = &root

	// Parse all directories and files and shove them into our crude tree structure
	for _, e := range data {
		if utils.StartsWith(e, "$ cd ") {
			if utils.Remainder(e, "$ cd ") == ".." {
				currentDir = currentDir.parent
			} else {
				var newNode = Node{parent: currentDir, files: []int{}, dirs: []*Node{}}
				currentDir.dirs = append(currentDir.dirs, &newNode)
				currentDir = &newNode
			}
		}

		if unicode.IsDigit(rune(e[0])) {
			var filesize int = utils.StringToInt(strings.Split(e, " ")[0])
			currentDir.files = append(currentDir.files, filesize)
		}
	}

	dirList = []int{}
	findAllDirsWithCondition(&root, 100000, "max")
	fmt.Println("Day 7 Part 1 solution: " + strconv.Itoa(utils.Sum(dirList)))

	var memoryNeeded = 30000000 - (70000000 - findAllDirsWithCondition(&root, -1, "none"))

	dirList = []int{}
	findAllDirsWithCondition(&root, memoryNeeded, "min")
	fmt.Println("Day 7 Part 2 solution: " + strconv.Itoa(utils.Min(dirList)))
}
