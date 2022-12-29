package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadFile(path string) *bufio.Scanner {
	readFile, _ := os.Open(path)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	return fileScanner
}

type path struct {
	name   string
	length int
}

type valve struct {
	name    string
	flow    int
	tunnels []string
	paths   map[string]path
}

func interpretValve(line []string, valves map[string]*valve) {
	key := line[1]
	flow, _ := strconv.Atoi(line[4][5 : len(line[4])-1])
	var tunnels []string
	for _, v := range line[9:] {
		if len(v) == 3 {
			tunnels = append(tunnels, v[:2])
		} else {
			tunnels = append(tunnels, v)
		}
	}
	valves[key] = &valve{name: key, flow: flow, tunnels: tunnels, paths: make(map[string]path)}
}

func (v *valve) isAdjacent(dest string) bool {
	for _, valve := range v.tunnels {
		if valve == dest {
			return true
		}
	}
	return false
}

func (v *valve) setPath(target string, distance int) {
	// v.paths = append(v.paths, path{name: target, length: distance})
	v.paths[target] = path{name: target, length: distance}
}

// need to give more generic name
func checkedTunnels(dest string, checked []string) bool {
	for _, v := range checked {
		if dest == v {
			return true
		}
	}
	return false
}

func checkOpened(v valve, o []*valve) bool {
	for _, x := range o {
		if x.name == v.name {
			return true
		}
	}
	return false
}

func shortestPath(v *valve, target string, valves map[string]*valve, count int, checked []string, sp int) int {
	checked = append(checked, v.name)

	count += 1

	if v.isAdjacent(target) {
		if sp > count || sp == 0 {
			sp = count
		}
	}

	for _, tunnel := range v.tunnels {
		if !checkedTunnels(tunnel, checked) {
			sp = shortestPath(valves[tunnel], target, valves, count, checked, sp)

		}
	}

	return sp
}

func checkAllOpened(valves map[string]*valve, opened []*valve) bool {
	allOpened := 0
	for _, v := range valves {
		if v.flow != 0 {
			allOpened += 1
		}
	}
	if len(opened) != allOpened {
		return false
	} else {
		return true
	}
}

func noValidMoves(current valve, valves map[string]*valve, opened []*valve, time int) bool {
	for _, v := range valves {
		if !checkOpened(*v, opened) && v.flow > 0 && current.paths[v.name].length <= (30-time) {
			return false
		}
	}
	return true
}

func calcScore(current valve, valves map[string]*valve, time int, released int, opened []*valve, best int, moves []string) int {
	// check every open valve to calc current flow
	currentFlow := 0
	if len(opened) != 0 {
		for _, v := range opened {
			currentFlow += v.flow
		}
	}

	// add current flow to score
	// open new valve if needed, add to opened
	if !checkOpened(current, opened) && current.flow != 0 {
		time += 1

		released += currentFlow
		moves = append(moves, "Opening valve "+current.name+" at time "+strconv.Itoa(time)+" with released "+strconv.Itoa(released)+" and flow "+strconv.Itoa(currentFlow))
		opened = append(opened, &current)
		currentFlow += current.flow

	}

	// add time. if time is up, return score || base case
	if time == 30 {
		if strings.Fields(moves[1])[2] == "DD" &&
			strings.Fields(moves[2])[2] == "BB" &&
			strings.Fields(moves[6])[2] == "CC" &&
			strings.Fields(moves[4])[2] == "HH" &&
			strings.Fields(moves[5])[2] == "EE" {
			for _, v := range moves {
				fmt.Println(v)
			}
			fmt.Println(released)

		}
		if released > best {
			fmt.Println(released)
			best = released
		}

	} else if checkAllOpened(valves, opened) && time < 30 {
		time += 1
		released += currentFlow
		best = calcScore(current, valves, time, released, opened, best, moves)

		// needs to be a check here for what to do if there are no remaining possible moves, such that the current flow will be incremented
	} else if noValidMoves(current, valves, opened, time) {
		best = calcScore(current, valves, 30, released+(currentFlow)*(30-time), opened, best, moves)
	} else if time < 30 {
		for _, v := range valves {
			// if the valve in consideration is not the current one
			// does not have a flow of 0
			// and has not been opened, then head towards it while the score is under 30
			// and add the distance to it to our time
			if v != &current && !checkOpened(*v, opened) && v.flow != 0 {
				distance := current.paths[v.name].length
				best = calcScore(*v, valves, (time + distance), released+(currentFlow*(distance)), opened, best, moves)
			}
		}
	}

	return best
}

func main() {
	start := time.Now()
	input := ReadFile("config/test_input.txt")
	// input := ReadFile("config/real_input.txt")
	valves := make(map[string]*valve)
	for input.Scan() {
		interpretValve(strings.Fields(input.Text()), valves)
	}
	for _, v := range valves {
		for _, targetv := range valves {
			if v.name == targetv.name {
				continue
			} else {
				var checked []string
				dist := shortestPath(v, targetv.name, valves, 0, checked, 0)
				// fmt.Println("Shortest distance between", v.name, " and ", targetv.name, "is ", dist)
				v.setPath(targetv.name, dist)
			}
		}

	}
	var opened []*valve
	var moves []string
	moves = append(moves, "Starting at A at time 0 with released 0")
	best := calcScore(*valves["AA"], valves, 0, 0, opened, 0, moves)
	fmt.Println(best)
	fmt.Println(time.Since(start))
	// 1586 too low
	// 	for _, v := range valves {
	// 		fmt.Println(v)
	// 	}
}
