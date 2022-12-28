package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadFile(path string) *bufio.Scanner {
	readFile, err := os.Open(path)

	if err != nil {
		//fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return fileScanner
}

func boundaries(y int, x int, boundary int) bool {
	return y == 0 || y == boundary || x == 0 || x == boundary
}

func convertRune(r rune) int {
	value, _ := strconv.Atoi(string(r))
	return value
}

func directionCondition(s string, x int, y int, len int) bool {
	switch s {
	case "right":
		return y == len
	case "down":
		return x == len
	case "left":
		return y == 0
	default: //up, but golang is dumb and demands either a useless return or to include a default
		return x == 0
	}
}

func directionIncrementation(s string, x int, y int) (int, int) {
	switch s {
	case "right":
		return x, y + 1
	case "down":
		return x + 1, y
	case "left":
		return x, y - 1
	default:
		return x - 1, y
	}
}

func directionExplorer(tree int, x int, y int, trees []string, direction string) int {
	var amount int

	if tree > convertRune(rune(trees[x][y])) {
		if directionCondition(direction, x, y, len(trees)-1) {
			return amount + 1
		} else {
			x, y = directionIncrementation(direction, x, y)
			amount += 1 + directionExplorer(tree, x, y, trees, direction)
		}
	} else {
		return amount + 1
	}
	return amount
}

func main() {

	input := ReadFile("config/real_input.txt")
	var trees []string
	for input.Scan() {
		trees = append(trees, input.Text())
	}
	maxViewable := 0
	for y, treeLine := range trees {
		for x, tree := range treeLine {
			if boundaries(x, y, len(trees)-1) {
				continue
			}
			right := directionExplorer(convertRune(tree), y, x+1, trees, "right")
			down := directionExplorer(convertRune(tree), y+1, x, trees, "down")
			left := directionExplorer(convertRune(tree), y, x-1, trees, "left")
			up := directionExplorer(convertRune(tree), y-1, x, trees, "up")
			viewableScore := right * down * left * up
			if viewableScore > maxViewable {
				maxViewable = viewableScore
			}
		}
	}
	fmt.Println(maxViewable)

}
