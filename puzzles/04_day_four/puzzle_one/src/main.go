package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func read_file(path string) *bufio.Scanner {
	readFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return fileScanner
}

func commaSplit(line string) []string {
	columns := strings.Split(line, ",")
	return columns
}

func hyphenSplit(line string) []string {
	columns := strings.Split(line, "-")
	return columns
}

func evaluateBounds(left []string, right []string) bool {
	leftLower, _ := strconv.Atoi(left[0])
	leftUpper, _ := strconv.Atoi(left[1])
	rightLower, _ := strconv.Atoi(right[0])
	rightUpper, _ := strconv.Atoi(right[1])
	return leftLower <= rightLower && leftUpper >= rightUpper
}

func main() {
	fileContents := read_file("config/real_input.txt")
	res := 0
	for fileContents.Scan() {
		newLine := fileContents.Text()
		columns := commaSplit((newLine))
		left, right := hyphenSplit(columns[0]), hyphenSplit(columns[1])
		if evaluateBounds(left, right) || evaluateBounds(right, left) {
			res += 1
		}
	}
	fmt.Println(res)
}
