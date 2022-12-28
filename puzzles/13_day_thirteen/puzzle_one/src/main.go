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

func (n *nestedArray) printValues() {
	for _, v := range n.values {
		switch y := v.(type) {
		case int:
			fmt.Println(y)
		case *nestedArray:
			y.printValues()
		}
	}
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
			fmt.Println("comparing", l.values, "and", r.values)
			for i, v := range l.values {
				diff = compareLines(v, r.values[i])
				if diff != 0 {
					return diff
				} else if i == len(r.values)-1 {
					// fmt.Println("right ran out of values")
					return len(l.values) - len(r.values)
				} else if i == len(l.values)-1 {
					// fmt.Println("left ran out of values")
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

	input := ReadFile("config/test_input.txt")
	var left, right nestedArray
	iterator := 1
	count := 0
	for input.Scan() {

		switch iterator % 3 {
		case 1:
			left = convertToNestedArray(input.Text())
			fmt.Println(input.Text())

		case 2:
			right = convertToNestedArray(input.Text())
			fmt.Println(input.Text())
			d := compareLines(&left, &right)
			fmt.Println("index:", (iterator+1)/3)
			fmt.Println(d)
			if d < 0 {
				count += (iterator + 1) / 3
			}
		}
		iterator += 1
	}
	fmt.Println(count)
}
