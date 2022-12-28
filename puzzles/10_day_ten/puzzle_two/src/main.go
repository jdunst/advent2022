package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type splice struct {
	left  int
	mid   int
	right int
}

func ReadFile(path string) *bufio.Scanner {
	readFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return fileScanner
}

func channelMatrix(position int) int {
	switch {
	case position <= 40:
		return 0
	case position <= 80:
		return 1
	case position <= 120:
		return 2
	case position <= 160:
		return 3
	case position <= 200:
		return 4
	default:
		return 5
	}
}

func (s *splice) adjustSplice(position int) {
	s.mid = position
	s.left = position - 1
	s.right = position + 1
}

func splitLine(input string) (string, int) {
	fields := strings.Fields(input)
	if fields[0] == "addx" {
		value, _ := strconv.Atoi(fields[1])
		return fields[0], value
	} else {
		return fields[0], 0
	}
}

func determineCharacter(cycle int, splice splice) string {
	checkValue := (channelMatrix(cycle)) * 40
	comparisonValue := (cycle - checkValue) - 1
	if cycle <= 41 {
		//fmt.Println(comparisonValue)
		fmt.Println(splice)
	}
	if comparisonValue == splice.left || comparisonValue == splice.mid || comparisonValue == splice.right {
		fmt.Println("#")
		fmt.Println(comparisonValue)
		return "#"
	} else {
		fmt.Println(".")
		fmt.Println(comparisonValue)
		return "."
	}
}

func main() {

	input := ReadFile("config/real_input.txt")
	signals := make(map[int]int)
	cycle := 0
	channels := [6]string{"", "", "", "", "", ""}
	splice := splice{left: 0, mid: 1, right: 2}
	for input.Scan() {
		instruction, value := splitLine(input.Text())
		if instruction == "noop" {
			cycle += 1
			signals[cycle] = signals[cycle-1]
			channels[channelMatrix(cycle)] += determineCharacter(cycle, splice)
			//fmt.Println("it is cycle: ", cycle, "and the current value is", signals[cycle])
			// in order:
			// increase the cycle
			// check if cycle in range of splice values
			// if yes, add # to current channel's string
			// if not, add . to current channel's string
			// adjust the value

		} else {
			cycle += 1
			if cycle == 1 {
				signals[cycle] = signals[cycle]
			} else {
				signals[cycle] = signals[cycle-1]
			}

			channels[channelMatrix(cycle)] += determineCharacter(cycle, splice)

			cycle += 1
			signals[cycle] = signals[cycle-1] + value
			channels[channelMatrix(cycle)] += determineCharacter(cycle, splice)
			splice.adjustSplice(signals[cycle] + 1)
			//fmt.Println("it is cycle: ", cycle, "and the current value is", signals[cycle])

		}

	}
	for _, line := range channels {
		fmt.Println(line)
	}
}
