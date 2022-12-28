package main

import (
	"bufio"
	"fmt"
	"os"
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

func sharedCharacters(repeats []rune, newLine []rune) (repeatedCharacters []rune) {
	for _, leftChar := range repeats {
		for _, rightChar := range newLine {
			if leftChar == rightChar {
				repeatedCharacters = append(repeatedCharacters, leftChar)
				break
			}
		}
	}
	return repeatedCharacters
}

func main() {
	fileContents := read_file("config/real_input.txt")
	var repeatValues []rune
	score := 0
	x := 0
	for fileContents.Scan() {
		currentLine := []rune(fileContents.Text())
		if x == 0 {
			repeatValues = currentLine
		} else {
			repeatValues = sharedCharacters(repeatValues, currentLine)
		}
		if x == 2 {
			// uppercase letters start at runic 65, so we subtract 38
			// lowercase letters start at runic 97, so we subtract 96
			if repeatValues[0] >= 97 {
				score += (int(repeatValues[0]) - 96)
			} else {
				score += (int(repeatValues[0]) - 38)
			}
			x = 0
		} else {
			x += 1
		}

	}
	fmt.Println(score)
}
