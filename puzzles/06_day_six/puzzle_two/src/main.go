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

func generateMap(instruction string, buildTo int) map[string]int {
	valueMap := make(map[string]int)
	for _, char := range instruction[:buildTo] {
		valueMap[string(char)] += 1
	}
	return valueMap
}

func checkMap(m map[string]int) bool {
	for _, v := range m {
		if v > 1 {
			return false
		}
	}
	return true
}

func main() {

	instructions := ReadFile("config/real_input.txt")
	desiredLen := 14
	var res int

	for instructions.Scan() {
		newLine := instructions.Text()
		valueMap := generateMap(newLine, desiredLen)

		for idx, v := range newLine[desiredLen:] {
			if checkMap(valueMap) {
				res = idx + desiredLen
				break
			}
			valueMap[string(newLine[idx])] -= 1 // subtract from first position
			if valueMap[string(newLine[idx])] == 0 {
				delete(valueMap, string(newLine[idx]))
			}
			valueMap[string(v)] += 1 // add to the new character
		}
		fmt.Println(res)
	}
}
