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

func checkVisited(s []string, x int, y int) bool {
	for _, v := range s {
		if v == fmt.Sprint(x)+fmt.Sprint(y) {
			return true
		}
	}
	return false
}

func explore(grid []string, x int, y int, visited []string, count int, trips *[]int, shortest map[string]int) {
	coord := fmt.Sprint(x) + fmt.Sprint(y)
	// if x == 2 && y == 5 {
	if x == 20 && y == 43 {

		*trips = append(*trips, len(visited))
		if shortest[coord] == 0 {
			shortest[coord] = len(visited)
		} else if shortest[coord] > len(visited) {
			shortest[coord] = len(visited)
		}
	} else {

		visited = append(visited, coord)
		if shortest[coord] == 0 {
			shortest[coord] = len(visited)
		} else if shortest[coord] > len(visited) {
			shortest[coord] = len(visited)
		} else {
			return
		}
		//fmt.Println(visited)
		//valid_move := false
		if x != 0 {
			if grid[x][y]+1 >= grid[x-1][y] && !checkVisited(visited, x-1, y) {
				//valid_move = true
				count += 1
				explore(grid, x-1, y, visited, count, trips, shortest)
			}
		}
		if y != 0 {
			if grid[x][y]+1 >= grid[x][y-1] && !checkVisited(visited, x, y-1) {
				//valid_move = true
				count += 1
				explore(grid, x, y-1, visited, count, trips, shortest)
			}
		}
		if x != (len(grid) - 1) {
			if grid[x][y]+1 >= grid[x+1][y] && !checkVisited(visited, x+1, y) {
				//valid_move = true
				count += 1
				explore(grid, x+1, y, visited, count, trips, shortest)
			}
		}
		if y != (len(grid[0]) - 1) {
			if grid[x][y]+1 >= grid[x][y+1] && !checkVisited(visited, x, y+1) {
				//valid_move = true
				count += 1
				explore(grid, x, y+1, visited, count, trips, shortest)
			}
		}
	}
}

func main() {

	input := ReadFile("config/real_input.txt")
	// input := ReadFile("config/test_input.txt")
	var grid []string
	for input.Scan() {
		grid = append(grid, input.Text())
	}

	var previous []string
	var trips []int
	shortestPath := make(map[string]int)

	// 20, 0 for real
	// 0, 0 for test
	explore(grid, 20, 0, previous, 0, &trips, shortestPath)
	// explore(grid, 0, 0, previous, 0, &trips, shortestPath)
	// fmt.Println(shortestPath["25"])
	fmt.Println(shortestPath["2043"])
}
