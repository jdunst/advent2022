package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func splitLine(input string) (string, int) {
	fields := strings.Fields(input)
	if fields[0] == "addx" {
		value, _ := strconv.Atoi(fields[1])
		return fields[0], value
	} else {
		return fields[0], 0
	}
}

func main() {

	input := ReadFile("config/real_input.txt")
	signals := make(map[int]int)
	cycles := 1
	signals[cycles] = 1
	for input.Scan() {
		instruction, value := splitLine(input.Text())
		if instruction == "noop" {
			cycles += 1
			signals[cycles] = signals[cycles-1]
			fmt.Println("it is cycle: ", cycles, "and the current value is", signals[cycles])

		} else {
			cycles += 1
			signals[cycles] = signals[cycles-1]
			cycles += 1
			signals[cycles] = signals[cycles-1] + value
			fmt.Println("it is cycle: ", cycles, "and the current value is", signals[cycles])

		}

	}
	signalsToTry := [6]int{20, 60, 100, 140, 180, 220}
	sum := 0
	for _, v := range signalsToTry {
		sum += signals[v] * v
	}
	fmt.Println(sum)
}
