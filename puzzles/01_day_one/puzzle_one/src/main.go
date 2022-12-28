package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func read_file(path string) (bpos int) {
	readFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	currentPosition := 1
	biggestPosition := 1
	biggestWeight := 0
	currentWeight := 0

	for fileScanner.Scan() {
		content := fileScanner.Text()
		if content == "" {
			//fmt.Println(currentWeight)
			fmt.Println("Current elf: %s Their total weight: %s, current biggest positionholder: %s", currentPosition, currentWeight, biggestPosition)
			if currentWeight > biggestWeight {
				biggestWeight = currentWeight
				biggestPosition = currentPosition
				currentWeight = 0

			} else {
				currentWeight = 0
			}
			currentPosition += 1

		} else {
			val, _ := strconv.Atoi(content)
			currentWeight += val

		}
		//fmt.Println(content)

	}

	return biggestPosition
}

func main() {
	bPos := read_file("config/elves.txt")
	fmt.Println(bPos)
}
