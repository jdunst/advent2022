package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(path string) *bufio.Scanner {
	readFile, _ := os.Open(path)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	return fileScanner
}

func main() {
	input := ReadFile("config/test_input.txt")
	// input := ReadFile("config/real_input.txt")
	for input.Scan() {
		fmt.Println(input.Text())
	}
}
