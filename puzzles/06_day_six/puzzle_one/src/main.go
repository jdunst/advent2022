package main

import (
	"bufio"
	"fmt"
	"os"
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

type positions struct {
	one   string
	two   string
	three string
}

func checkTriplet(p positions) bool {
	return p.one == p.two || p.two == p.three || p.three == p.one
}

func stringIterator(instruction string) int {

	var res int = 4
	var pos = positions{one: string(instruction[0]), two: string(instruction[1]), three: string(instruction[2])}

	for idx, char := range instruction[3:] {
		if string(char) == pos.one || string(char) == pos.two || string(char) == pos.three || checkTriplet(pos) {
			pos.one = pos.two
			pos.two = pos.three
			pos.three = string(char)
			fmt.Println(pos)
		} else {
			res += idx
			break
		}
	}
	return res
}

func main() {

	instructions := ReadFile("config/real_input.txt")

	for instructions.Scan() {
		newLine := instructions.Text()
		res := stringIterator(newLine)
		fmt.Println(res)
	}

}
