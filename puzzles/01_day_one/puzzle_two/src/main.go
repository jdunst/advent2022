package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func read_file(path string) (allWeights []int) {
	readFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var elfWeights []int
	currentWeight := 0

	for fileScanner.Scan() {
		content := fileScanner.Text()
		if content == "" {
			elfWeights = append(elfWeights, currentWeight)
			currentWeight = 0
		} else {
			val, _ := strconv.Atoi(content)
			currentWeight += val

		}
	}

	return elfWeights
}

func main() {
	bPos := read_file("config/elves.txt")
	sort.Sort(sort.Reverse(sort.IntSlice(bPos)))
	fmt.Println(bPos)
}
