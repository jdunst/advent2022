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
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return fileScanner
}

func assembleTreelines(t []string) ([]string, []string) {
	var leftRight []string
	var upDown []string
	for rowIndex, tree := range t[1 : len(t)-1] {
		leftRight = append(leftRight, tree)
		var vertTree string
		for columnIndex, _ := range tree {
			// when we iterate over a tree, column and row switch because we are moving from left
			// to right instead of from top to bottom
			vertTree += string(t[columnIndex][rowIndex+1])
		}
		upDown = append(upDown, vertTree)
	}
	return leftRight, upDown
}

func reverseTree(tree string) string {
	var revTree string
	for x := len(tree) - 1; x >= 0; x-- {
		revTree += string(tree[x])
	}

	return revTree
}

func convertRune(r rune) int {
	value, _ := strconv.Atoi(string(r))
	return value
}

func createPosition(row int, col int, flag string) string {
	if flag == "ud" {
		return "r" + fmt.Sprint(col) + "c" + fmt.Sprint(row)
	}
	return "r" + fmt.Sprint(row) + "c" + fmt.Sprint(col)
}

func iterateTree(index int, tree string, d map[string][]int, flag string) {
	revTree := reverseTree(tree)
	high := convertRune(rune(tree[0]))
	revHigh := convertRune(rune(revTree[0]))

	d[createPosition(index+1, 0, "lr")] = append(d[createPosition(index+1, 0, "lr")], 1)
	d[createPosition(index+1, len(tree)-1, "lr")] = append(d[createPosition(index+1, len(tree)-1, "lr")], 1)
	d[createPosition(0, index+1, "lr")] = append(d[createPosition(0, index+1, "lr")], 1)
	d[createPosition(len(tree)-1, index+1, "lr")] = append(d[createPosition(len(tree)-1, index+1, "lr")], 1)
	for idx, val := range tree {
		if convertRune(val) > high {
			high = convertRune(val)
			d[createPosition(index+1, idx, flag)] = append(d[createPosition(index+1, idx, flag)], 1)
		}
	}
	for idx, val := range revTree {
		if convertRune(val) > revHigh {
			revHigh = convertRune(val)
			d[createPosition(index+1, len(tree)-idx-1, flag)] = append(d[createPosition(index+1, len(tree)-idx-1, flag)], 1)
		}
	}
}

func main() {

	input := ReadFile("config/real_input.txt")
	var trees []string
	for input.Scan() {
		trees = append(trees, input.Text())
	}

	visible := make(map[string][]int)
	l, u := assembleTreelines(trees)

	for idx, tree := range l {
		iterateTree(idx, tree, visible, "lr")
	}
	for idx, tree := range u {
		iterateTree(idx, tree, visible, "ud")
	}
	fmt.Println(len(visible) + 4)
}
