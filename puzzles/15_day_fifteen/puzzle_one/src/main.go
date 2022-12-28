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
		return true
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

func main() {
	input := ReadFile("config/real_input.txt")
	var sensors []sensor
	for input.Scan() {
		sensor := parseLines(input.Text())
		// get all our sensor/beacon pairs
		sensors = append(sensors, sensor)
	}

	fleft := 0
	fright := 0
	for _, s := range sensors {

		if fleft > (s.sx - s.manhattan) {
			fleft = (s.sx - s.manhattan)
		}
		if fright < (s.sx + s.manhattan) {
			fright = (s.sx + s.manhattan)
		}
	}

	count := 0
	for i := fleft; i <= fright; i++ {
		valid := true
		for _, sensor := range sensors {
			valid = withinManhattan(sensor, i, 2000000)
			if !valid {
				// fmt.Println("Position ", i, 10, "is not valid due to sensor", sensor.sx, sensor.sy)
				break
			}
		}
		if !valid {
			count += 1
		}

	}
	fmt.Println(count)
}
