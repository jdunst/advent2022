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

type sensor struct {
	sx        int
	sy        int
	bx        int
	by        int
	manhattan int
}

func (s *sensor) manhattanDistance() {
	xDist := s.bx - s.sx
	yDist := s.by - s.sy
	if xDist < 0 {
		xDist *= -1
	}
	if yDist < 0 {
		yDist *= -1
	}
	s.manhattan = xDist + yDist
}

func extractValue(strs []string, idx int) int {
	var conditional int
	if idx == 9 {
		conditional = 0
	} else {
		conditional = 1
	}
	val, _ := strconv.Atoi(strs[idx][2 : len(strs[idx])-conditional])
	return val
}

func parseLines(line string) sensor {
	contents := strings.Fields(line)
	newSensor := sensor{
		sx: extractValue(contents, 2),
		sy: extractValue(contents, 3),
		bx: extractValue(contents, 8),
		by: extractValue(contents, 9)}
	// find manhattan distance
	newSensor.manhattanDistance()
	return newSensor
}

func withinManhattan(s sensor, x int, y int) bool {
	// fmt.Println("I am checking sensor", s, "to see if", x, y, "is within its manhattan distance of", s.manhattan)
	if x == s.bx && y == s.by {
		return false
	}

	xDist := s.sx - x
	yDist := s.sy - y
	if xDist < 0 {
		xDist *= -1
	}
	if yDist < 0 {
		yDist *= -1
	}

	if s.manhattan >= (xDist + yDist) {
		return false
	} else {
		return true
	}
}

func calcJump(s sensor, x int, y int) int {
	xDist := s.sx - x
	yDist := s.sy - y
	if xDist < 0 {
		xDist *= -1
	}
	if yDist < 0 {
		yDist *= -1
	}

	return s.manhattan - xDist - yDist
}

func main() {
	input := ReadFile("config/real_input.txt")
	var sensors []sensor
	for input.Scan() {
		sensor := parseLines(input.Text())
		// get all our sensor/beacon pairs
		sensors = append(sensors, sensor)
	}
	bx, by := 0, 0

	for x := 0; x <= 4000000; x++ {
		for y := 0; y <= 4000000; y++ {
			valid := true

			for _, sensor := range sensors {
				valid = withinManhattan(sensor, x, y)
				if !valid {
					// fmt.Println("Position ", x, y, "is not valid due to sensor", sensor.sx, sensor.sy, sensor.manhattan)
					jump := calcJump(sensor, x, y)
					// fmt.Println(x, y, "I would skip to", x, y+jump)
					if jump >= 1 {
						y = (y + jump - 1)
					}
					break
				}
			}
			if valid {
				fmt.Println("Position", x, y, "is valid")
				bx, by = x, y
				break
			}
		}
	}
	fmt.Println(bx*4000000 + by)
}
