package main

import (
	"bufio"
	"fmt"
	"math"
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
	value, _ := strconv.Atoi(fields[1])
	return fields[0], value
}

type head struct {
	x  int
	y  int
	px int
	py int
}

type tail struct {
	x       int
	y       int
	visited map[string]int
}

func (h *head) updatePrevious() {
	h.py = h.y
	h.px = h.x
}

func (h *head) moveUp() {
	h.updatePrevious()
	h.x += 1
}

func (h *head) moveRight() {
	h.updatePrevious()
	h.y += 1
}

func (h *head) moveLeft() {
	h.updatePrevious()
	h.y -= 1
}

func (h *head) moveDown() {
	h.updatePrevious()
	h.x -= 1
}

func (t *tail) determinePosition() string {
	x := fmt.Sprint(t.x)
	y := fmt.Sprint(t.y)
	return "r" + x + "c" + y
}

func (t *tail) shouldIMove(h *head) bool {
	if math.Abs(float64(h.y)-float64(t.y)) == 1 && math.Abs(float64(h.x)-float64(t.x)) == 1 {
		return false
	} else if t.x == h.x && math.Abs(float64(h.y)-float64(t.y)) == 1 {
		return false
	} else if t.y == h.y && math.Abs(float64(h.x)-float64(t.x)) == 1 {
		return false
	} else {
		return true
	}
}

func (t *tail) butShouldIMoveDiagonally(h *head) bool {
	if math.Abs(float64(h.x)-float64(t.x)) > 1 || math.Abs(float64(h.y)-float64(t.y)) > 1 {
		return true
	} else {
		return false
	}
}

func (t *tail) moveQuoteDiagonallyUnquote(h *head) {
	t.x = h.px
	t.y = h.py
	t.visited[t.determinePosition()] = 1
}

func main() {

	input := ReadFile("config/real_input.txt")

	head := head{x: 0, y: 0}
	tail := tail{x: 0, y: 0, visited: make(map[string]int)}

	for input.Scan() {
		instruction, value := splitLine(input.Text())
		for x := 0; x < value; x++ {
			switch instruction {
			case "R":
				head.moveRight()
			case "L":
				head.moveLeft()
			case "U":
				head.moveUp()
			case "D":
				head.moveDown()
			}
			if tail.shouldIMove(&head) {
				if tail.butShouldIMoveDiagonally(&head) {
					tail.moveQuoteDiagonallyUnquote(&head)
				}
			}
		}
	}
	fmt.Println(len(tail.visited) + 1)

}
