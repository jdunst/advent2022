package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
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

func splitString(line string) ([]rune, []rune) {
	halfway := utf8.RuneCountInString(line) / 2
	return []rune(line)[0:halfway], []rune(line)[halfway:]
}

func sharedCharacter(left []rune, right []rune) (repeatedCharacter rune) {

	for _, leftChar := range left {
		for _, rightChar := range right {
			if leftChar == rightChar {
				repeatedCharacter = leftChar
				break
			}
		}

	}
	return repeatedCharacter
}

func main() {
	fileContents := read_file("config/real_input.txt")
	score := 0
	for fileContents.Scan() {
		newLine := fileContents.Text()
		repeatValue := sharedCharacter(splitString(newLine))
		// uppercase letters start at runic 65, so we subtract 38
		// lowercase letters start at runic 97, so we subtract 96
		if repeatValue >= 97 {
			score += (int(repeatValue) - 96)
		} else {
			score += (int(repeatValue) - 38)
		}

	}
	fmt.Println(score)
}
