package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// type player map[string]string {
// 	X string
// 	Y string
// 	Z string
// }

// type opponent struct {
// 	A string
// 	B string
// 	C string
// }

func read_file(path string) *bufio.Scanner {
	readFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return fileScanner
}

func splitString(line string) []string {
	columns := strings.Fields(line)
	return columns
}

// A is rock, B is paper, C is scissors
// X is lose, Y is draw, Z is win
// Rock is 1, Paper is 2, Scissors is 3
// Lose is 0, Draw is 3, Win is 6
func main() {
	fileContents := read_file("config/real_input.txt")
	playerMap := map[string]string{"X": "lose", "Y": "draw", "Z": "win"}
	opponentMap := map[string]map[string]int{"A": {"lose": 3, "draw": 4, "win": 8},
		"B": {"lose": 1, "draw": 5, "win": 9},
		"C": {"lose": 2, "draw": 6, "win": 7}}

	score := 0

	for fileContents.Scan() {
		newLine := fileContents.Text()
		inputs := splitString(newLine)
		scenario := playerMap[inputs[1]]
		score += opponentMap[inputs[0]][scenario]
	}
	fmt.Println(score)
}
