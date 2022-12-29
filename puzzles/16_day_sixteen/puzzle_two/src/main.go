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

func interpretValve(line []string, valves map[string]*valve, fValves map[string]*valve) {
	// translates lines into valve types
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
	valves[key] = &valve{name: key, flow: flow, tunnels: tunnels}
	if flow != 0 || key == "AA" {
		fValves[key] = &valve{name: key, flow: flow, tunnels: tunnels, paths: make(map[string]path)}
	}
}

func (v *valve) isAdjacent(dest string) bool {
	// determines if a given valve name is in the current valve type's adjacent tunnels
	for _, valve := range v.tunnels {
		if valve == dest {
			return true
		}
	}
	return false
}

func (v *valve) setPath(target string, distance int) {
	// receiver method to set the paths map to have a name and distance for a given path
	v.paths[target] = path{name: target, length: distance}
}

func checkedTunnels(dest string, checked []string) bool {
	// boolean for our path determinator
	for _, v := range checked {
		if dest == v {
			return true
		}
	}
	return false
}

func checkOpened(v valve, o []*valve) bool {
	// boolean for our best score determinator
	for _, x := range o {
		if x.name == v.name {
			return true
		}
	}
	return false
}

func shortestPath(v *valve, target string, valves map[string]*valve, checked []string, sp int) int {
	// takes a map of valves with flow and a map of all valves and finds the shortest path between
	// valves with flow
	checked = append(checked, v.name)
	if v.isAdjacent(target) {
		if sp > len(checked) || sp == 0 {
			sp = len(checked)
		}
	}
	for _, tunnel := range v.tunnels {
		if !checkedTunnels(tunnel, checked) {
			sp = shortestPath(valves[tunnel], target, valves, checked, sp)

		}
	}
	return sp
}

func checkAllOpened(valves map[string]*valve, opened []*valve) bool {
	return len(opened) == len(valves)-1
}

func noValidMoves(current valve, valves map[string]*valve, opened []*valve, time int, duration int) bool {
	// check to see if remaining valves' distance is greater than the current time remaining.
	// if none are valid to move to, we begin just incrementing the flow in our recursive function
	for _, v := range valves {
		if !checkOpened(*v, opened) && current.paths[v.name].length <= (duration-time) {
			return false
		}
	}
	return true
}

func calcScore(current valve, valves map[string]*valve, time int, released int, opened []*valve, best int, duration int) int {
	// check every open valve to calculate current flow
	currentFlow := 0
	for _, v := range opened {
		currentFlow += v.flow
	}

	// open new valve if needed, add to opened
	// add to time, add to flow since it takes one minute to open
	if !checkOpened(current, opened) && current.flow != 0 {
		time += 1
		released += currentFlow
		opened = append(opened, &current)
		currentFlow += current.flow

	}

	if (checkAllOpened(valves, opened) && time < duration) || noValidMoves(current, valves, opened, time, duration) {
		// base case. ends recursion if there are no valid moves or all valves are opened
		// if time is > 30 then we add 0 * currentFlow, filtering out invalid scenarios
		released += (currentFlow * (duration - time))
		if released > best {
			best = released
		}

	} else if time < duration {
		// try possible moves while time is available
		for _, v := range valves {
			// if it's not the current valve, it's not our start (AA), and it's not opened, explore a path to it
			if v != &current && !checkOpened(*v, opened) && v.name != "AA" {
				distance := current.paths[v.name].length
				best = calcScore(*v, valves, (time + distance), released+(currentFlow*(distance)), opened, best, duration)
			}
		}
	}

	return best
}

func main() {
	start := time.Now()
	// input := ReadFile("config/test_input.txt")
	input := ReadFile("config/real_input.txt")

	// map of all valves
	allValves := make(map[string]*valve)
	// map of valves with flow
	flowValves := make(map[string]*valve)

	for input.Scan() {
		interpretValve(strings.Fields(input.Text()), allValves, flowValves)
	}
	for _, v := range flowValves {
		for _, targetv := range flowValves {
			if v.name == targetv.name {
				continue
			} else {
				var checked []string
				dist := shortestPath(v, targetv.name, allValves, checked, 0)
				v.setPath(targetv.name, dist)
			}
		}

	}
	var opened []*valve
	best := calcScore(*flowValves["AA"], flowValves, 0, 0, opened, 0, 30)
	fmt.Println(best)
	fmt.Println(time.Since(start))
}
