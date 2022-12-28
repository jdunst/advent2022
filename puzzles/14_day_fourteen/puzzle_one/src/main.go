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
		//fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return fileScanner
}

func lineSplit(input string) []string {
	return strings.Fields(input)
}

func removeArrows(input []string) []string {
	var sanitized []string
	for idx, val := range input {
		if idx == 0 || idx%2 == 0 {
			sanitized = append(sanitized, string(val))
		}
	}
	return sanitized
}

func printCave(cave [][]int) {
	for x := 0; x < 10; x++ {
		fmt.Print(x, ": ")
		fmt.Print(cave[x][493:504])
		fmt.Println()
	}
}

func convertToArrays(input []string) [][]string {
	var coordinates [][]string
	for _, coordinate := range input {
		coordinates = append(coordinates, strings.Split(coordinate, ","))
	}
	return coordinates
}

func interpretRock(moves [][]string, cave [][]int) {
	cy, _ := strconv.Atoi(moves[0][0])
	cx, _ := strconv.Atoi(moves[0][1])

	// for the purposes of this exercise, the incoming values are reversed
	// such that the 500-centric values are the nth number in a row
	// and a row is x
	// therefore 500,6 -> 500,4 is moving from the 500th column, 4th row
	// to 500th column, 6th row
	// x starts at the "top" of our cave object and is 600 0s long

	for x := 1; x < len(moves); x++ {
		ny, _ := strconv.Atoi(moves[x][0])
		nx, _ := strconv.Atoi(moves[x][1])
		if cx != nx {
			if nx-cx > 0 {
				// draw rock up
				for i := cx; i <= nx; i++ {
					cave[i][cy] = 1

				}
			} else {
				// draw rock down
				for i := cx; i >= nx; i-- {
					cave[i][cy] = 1
				}
			}
		} else {
			if ny-cy > 0 {
				// draw rock right
				for i := cy; i <= ny; i++ {
					cave[cx][i] = 1
				}
			} else {
				// draw rock left
				for i := cy; i >= ny; i-- {
					cave[cx][i] = 1
				}
			}
		}
		cy, cx = ny, nx

	}
}

func fallingSand(sx int, sy int, cave [][]int, bot bool) bool {
	if sx == len(cave)-1 {
		bot = true
		return bot
	}
	if cave[sx+1][sy] == 0 {
		bot = fallingSand(sx+1, sy, cave, bot)
	} else if cave[sx+1][sy-1] == 0 {
		bot = fallingSand(sx+1, sy-1, cave, bot)
	} else if cave[sx+1][sy+1] == 0 {
		bot = fallingSand(sx+1, sy+1, cave, bot)
	} else {
		cave[sx][sy] = 1

	}
	return bot
}

func constructCave() [][]int {
	var cave [][]int

	for x := 0; x < 600; x++ {
		var line []int
		for x := 0; x < 600; x++ {
			line = append(line, 0)
		}
		cave = append(cave, line)
	}
	return cave
}

func main() {

	input := ReadFile("config/real_input.txt")
	cave := constructCave()
	for input.Scan() {
		line := convertToArrays(removeArrows(lineSplit(input.Text())))
		interpretRock(line, cave)
	}
	exit := false
	count := 0
	for !exit {
		exit = fallingSand(0, 500, cave, exit)
		count++
	}
	fmt.Println(count - 1)
}
