package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instructions struct {
	quantity int
	from     int
	to       int
}

func read_file(path string) *bufio.Scanner {
	readFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return fileScanner
}

func instructionSplit(line string) instructions {
	columns := strings.Fields(line)
	quant, _ := strconv.Atoi(columns[1])
	from, _ := strconv.Atoi(columns[3])
	to, _ := strconv.Atoi(columns[5])
	return instructions{quantity: quant, from: from, to: to}
}

// [J]             [F] [M]
// [Z] [F]     [G] [Q] [F]
// [G] [P]     [H] [Z] [S] [Q]
// [V] [W] [Z] [P] [D] [G] [P]
// [T] [D] [S] [Z] [N] [W] [B] [N]
// [D] [M] [R] [J] [J] [P] [V] [P] [J]
// [B] [R] [C] [T] [C] [V] [C] [B] [P]
// [N] [S] [V] [R] [T] [N] [G] [Z] [W]
//  1   2   3   4   5   6   7   8   9

func followInstructions(i instructions, containers map[int][]string) {
	positionToStart := len(containers[i.from]) - i.quantity

	containersToMove := containers[i.from][positionToStart:]

	for idx := len(containersToMove) - 1; idx >= 0; idx-- {
		containers[i.to] = append(containers[i.to], containersToMove[idx])
	}
	containers[i.from] = containers[i.from][:positionToStart]
	fmt.Println(containers)
}

func main() {

	var containerMap = map[int][]string{
		1: {"N", "B", "D", "T", "V", "G", "Z", "J"},
		2: {"S", "R", "M", "D", "W", "P", "F"},
		3: {"V", "C", "R", "S", "Z"},
		4: {"R", "T", "J", "Z", "P", "H", "G"},
		5: {"T", "C", "J", "N", "D", "Z", "Q", "F"},
		6: {"N", "V", "P", "W", "G", "S", "F", "M"},
		7: {"G", "C", "V", "B", "P", "Q"},
		8: {"Z", "B", "P", "N"},
		9: {"W", "P", "J"},
	}

	instructions := read_file("config/test_input.txt")

	for instructions.Scan() {
		newLine := instructions.Text()
		c := instructionSplit(newLine)
		fmt.Printf("Instructions: move %s from stack %s to stack %s\n", fmt.Sprint(c.quantity), fmt.Sprint(c.from), fmt.Sprint(c.to))
		followInstructions(c, containerMap)
	}

}

//GFTNRBZPF
