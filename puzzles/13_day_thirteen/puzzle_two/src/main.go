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

type nestedArray struct {
	parent *nestedArray
	values []any
}

func (n *nestedArray) appendArray(c any) {
	n.values = append(n.values, c)
}

func createIntArray(i int) *nestedArray {
	arr := nestedArray{}
	arr.values = append(arr.values, i)
	return &arr
}

func convertToNestedArray(line string) nestedArray {

	var root nestedArray
	cwd := &root
	var currentString string

	for idx, char := range line {
		if idx == 0 {
			continue
		}
		switch string(char) {
		case "[":
			newArray := nestedArray{parent: cwd}
			cwd.appendArray(&newArray)
			cwd = &newArray
		case "]":
			if currentString != "" {
				charVal, _ := strconv.Atoi(currentString)
				cwd.appendArray(charVal)
				currentString = ""
			}

			if len(cwd.values) == 0 {
				cwd.appendArray("A")
			}

			cwd = cwd.parent
		case ",":
			if currentString != "" {
				charVal, _ := strconv.Atoi(currentString)
				cwd.appendArray(charVal)
				currentString = ""
			}
		default:
			currentString += string(char)
		}
	}

	return root
}

func compareLines(left any, right any) int {
	diff := 0

	switch l := left.(type) {
	case int:
		switch r := right.(type) {
		case int:
			return l - r
		case *nestedArray:
			lo := createIntArray(l)
			diff = compareLines(lo, r)
		case string:
			return 1
		}
	case *nestedArray:
		switch r := right.(type) {
		case int:
			ro := createIntArray(r)
			diff = compareLines(l, ro)
		case string:
			return 1
		case *nestedArray:
			for i, v := range l.values {
				diff = compareLines(v, r.values[i])
				if diff != 0 {
					return diff
				} else if i == len(r.values)-1 {
					return len(l.values) - len(r.values)
				} else if i == len(l.values)-1 {
					return len(l.values) - len(r.values)
				}
			}
		}
	case string:
		_, ok := right.(string)
		if ok {
			return 0
		} else {
			return -1
		}
	}

	return diff
}

func main() {
	input := ReadFile("config/real_input.txt")
	firstInput := convertToNestedArray("[[2]]")
	secondInput := convertToNestedArray("[[6]]")
	firstCount := 0
	secondCount := 0
	iterator := 1
	for input.Scan() {

		switch iterator % 3 {
		case 1:
			line := convertToNestedArray(input.Text())
			if compareLines(&line, &firstInput) < 0 {
				firstCount += 1
			}
			if compareLines(&line, &secondInput) < 0 {
				secondCount += 1
			}

		case 2:
			line := convertToNestedArray(input.Text())
			if compareLines(&line, &firstInput) < 0 {
				firstCount += 1
			}
			if compareLines(&line, &secondInput) < 0 {
				secondCount += 1
			}
		}
		iterator += 1
	}
	fmt.Println((firstCount + 1) * (secondCount + 2))
}
