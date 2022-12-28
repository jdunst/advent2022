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

// if the lower bound is less than or equal to the corresponding lowerbound
// AND the upper bound is also greater than or equal to the corresponding lowerbound
// then they overlap
// example
// if there are pairs 5-7 and 6-10
// then 5 is less than or equal to 6
// and 7 is greater than or equal to 6
// then we know 6 occurs at or between 5 and 7
// if there are pairs 2-6 and 4-8
// then 2 is less than or equal to 4
// and 6 is greater than or equal to 4

func evaluateBounds(left []string, right []string) bool {
	leftLower, _ := strconv.Atoi(left[0])
	leftUpper, _ := strconv.Atoi(left[1])
	rightLower, _ := strconv.Atoi(right[0])
	return leftLower <= rightLower && leftUpper >= rightLower
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
