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

func main() {
	fileContents := read_file("config/real_input.txt")
	playerMap := map[string]string{"X": "rock", "Y": "paper", "Z": "scissors"}
	opponentMap := map[string]string{"A": "rock", "B": "paper", "C": "scissors"}
	scoreSystem := map[string]int{"rock": 1, "paper": 2, "scissors": 3}

	score := 0

	for fileContents.Scan() {
		newLine := fileContents.Text()
		inputs := splitString(newLine)
		opponentChoice, playerChoice := opponentMap[inputs[0]], playerMap[inputs[1]]
		round := 0
		playerScore := scoreSystem[playerChoice]

		if playerChoice == opponentChoice {
			round = 3
		} else if (opponentChoice == "rock" && playerChoice == "scissors") ||
			(opponentChoice == "scissors" && playerChoice == "paper") ||
			(opponentChoice == "paper" && playerChoice == "rock") {
			round = 0
		} else {
			round = 6
		}
		fmt.Printf("Player choice: %s, Opponent choice: %s\n", playerChoice, opponentChoice)

		score += (round + playerScore)
		fmt.Println(score)
	}
	fmt.Println(score)
}
