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
	name    string
	x       int
	y       int
	px      int
	py      int
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

func (h *head) determinePosition() string {
	x := fmt.Sprint(h.x)
	y := fmt.Sprint(h.y)
	return "r" + x + "c" + y
}

func (t *head) shouldIMove(h *head) bool {
	if t.x == h.x && t.y == h.y {
		return false
	} else if math.Abs(float64(h.y)-float64(t.y)) == 1 && math.Abs(float64(h.x)-float64(t.x)) == 1 {
		return false
	} else if t.x == h.x && math.Abs(float64(h.y)-float64(t.y)) == 1 {
		return false
	} else if t.y == h.y && math.Abs(float64(h.x)-float64(t.x)) == 1 {
		return false
	} else {
		return true
	}
}

func (t *head) shouldIMoveToPrevious(h *head) bool {
	if t.x == h.x || t.y == h.y {
		return true
	} else {
		return false
	}
}

func (h *head) moveToPrevious(p *head) {
	h.updatePrevious()
	h.x = p.px
	h.y = p.py
	h.visited[h.determinePosition()] = 1
	//fmt.Println(h.name, "moving previously to ", h.determinePosition())
}

func (h *head) moveDiagonally(header *head) {
	h.updatePrevious()
	if header.x > h.x && header.y > h.y {
		h.y += 1
		h.x += 1
	} else if header.x < h.x && header.y > h.y {
		h.y += 1
		h.x -= 1
	} else if header.x < h.x && header.y < h.y {
		h.y -= 1
		h.x -= 1
	} else if header.x > h.x && header.y < h.y {
		h.y -= 1
		h.x += 1
	} else if header.x == h.x && header.y < h.y {
		h.y -= 1
	} else if header.x == h.x && header.y > h.y {
		h.y += 1
	} else if header.y == h.y && header.x < h.x {
		h.x -= 1
	} else {
		h.x += 1
	}
	//fmt.Println(h.name, "moving to ", h.determinePosition())
	h.visited[h.determinePosition()] = 1
}

func main() {

	input := ReadFile("config/real_input.txt")

	header := head{x: 0, y: 0, name: "header"}
	var knots []*head
	for x := 0; x < 9; x++ {
		nameVal := fmt.Sprint(x + 1)
		val := head{x: 0, y: 0, px: 0, py: 0, visited: make(map[string]int), name: nameVal}
		knots = append(knots, &val)
	}

	for input.Scan() {
		instruction, value := splitLine(input.Text())
		fmt.Println(input.Text())
		for x := 0; x < value; x++ {
			switch instruction {
			case "R":
				header.moveRight()
			case "L":
				header.moveLeft()
			case "U":
				header.moveUp()
			case "D":
				header.moveDown()
			}
			for i, v := range knots {

				if i == 0 {
					if v.shouldIMove(&header) {
						v.moveToPrevious(&header)
					}
				} else {
					if v.shouldIMove(knots[i-1]) {
						v.moveDiagonally(knots[i-1])
					}
				}

			}

		}

	}
	fmt.Println(len(knots[8].visited) + 1)
}
