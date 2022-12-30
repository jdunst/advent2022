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

func availablePumps(valves map[string]*valve, opened []*valve) []*valve {
	var available []*valve
	for _, v := range valves {
		open := false
		for _, o := range opened {
			if v == o {
				open = true
			}
		}
		if !open && v.name != "AA" {
			available = append(available, v)
		}
	}
	return available
}

func calcScore(current [2]*valve, valves map[string]*valve, times [2]int, released [2]int, masterOpened []*valve, opens [2][]*valve, flows [2]int, best int, duration int, moves []string) int {
	// check every open valve to calculate current flow
	flows[0], flows[1] = 0, 0
	for index := range flows {
		for _, openedValve := range opens[index] {
			flows[index] += openedValve.flow
		}

	}
	// fmt.Println(flows[0], flows[1])
	// open new valve if needed, add to opened
	// add to time, add to flow since it takes one minute to open
	for index := range current {
		if !checkOpened(*current[index], masterOpened) && current[index].flow != 0 {
			times[index] += 1
			released[index] += flows[index]
			masterOpened = append(masterOpened, current[index])
			opens[index] = append(opens[index], current[index])
			flows[index] += current[index].flow
		}
	}

	if len(valves) == len(masterOpened)+1 || (noValidMoves(*current[0], valves, masterOpened, times[0], duration) && noValidMoves(*current[1], valves, masterOpened, times[1], duration)) {
		if times[0] <= duration && times[1] <= duration {
			for i := 0; i < 2; i++ {
				released[i] += (flows[i] * (duration - times[i]))
			}
			totalReleased := (released[0] + released[1])
			if totalReleased > best {
				best = totalReleased
				fmt.Printf("total released: %v, with times %v and %v\n", totalReleased, times[0], times[1])
			}
		}

		// base case. ends recursion if there are no valid moves or all valves are opened
		// if time is > 30 then we add 0 * currentFlow, filtering out invalid scenarios

	} else if times[0] < duration && times[1] < duration {
		previousV := current[0]
		previousX := current[1]
		available := availablePumps(valves, masterOpened)

		// if there's two available and they both have valid moves
		if len(available) >= 2 && (!noValidMoves(*current[0], valves, masterOpened, times[0], duration) && !noValidMoves(*current[1], valves, masterOpened, times[1], duration)) {
			for _, v := range available {

				vdistance := current[0].paths[v.name].length
				times[0] += vdistance
				released[0] += (flows[0] * vdistance)
				current[0] = v
				for _, x := range available {
					if v != x {
						xdistance := current[1].paths[x.name].length
						times[1] += xdistance
						released[1] += (flows[1] * xdistance)
						current[1] = x
						best = calcScore(current, valves, times, released, masterOpened, opens, flows, best, duration, moves)
						times[1] -= xdistance
						released[1] -= (flows[1] * xdistance)
						current[1] = previousX
					}
				}
				times[0] -= vdistance
				released[0] -= (flows[0] * vdistance)
				current[0] = previousV
			}
			// if just the first runner has valid moves
		} else if !noValidMoves(*current[0], valves, masterOpened, times[0], duration) && noValidMoves(*current[1], valves, masterOpened, times[1], duration) {
			for _, v := range available {
				if v != current[0] {
					vdistance := current[0].paths[v.name].length
					times[0] += vdistance
					released[0] += (flows[0] * vdistance)
					current[0] = v
					best = calcScore(current, valves, times, released, masterOpened, opens, flows, best, duration, moves)
					times[0] -= vdistance
					released[0] -= (flows[0] * vdistance)
					current[0] = previousV
				}

			}
			// if just the second unner has valid moves
		} else if noValidMoves(*current[0], valves, masterOpened, times[0], duration) && !noValidMoves(*current[1], valves, masterOpened, times[1], duration) {
			for _, v := range available {
				vdistance := current[1].paths[v.name].length
				times[1] += vdistance
				released[1] += (flows[1] * vdistance)
				current[1] = v
				best = calcScore(current, valves, times, released, masterOpened, opens, flows, best, duration, moves)
				times[1] -= vdistance
				released[1] -= (flows[1] * vdistance)
				current[1] = previousX
			}
		}
	}

	return best
}

func main() {
	start := time.Now()
	input := ReadFile("config/real_input.txt")
	// input := ReadFile("config/test_input.txt")

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
	var allOpened []*valve
	current := [2]*valve{flowValves["AA"], flowValves["AA"]}
	opens := [2][]*valve{}
	times := [2]int{0, 0}
	flows := [2]int{0, 0}
	released := [2]int{0, 0}
	var moves []string
	best := calcScore(current, flowValves, times, released, allOpened, opens, flows, 0, 26, moves)
	fmt.Println(best)
	fmt.Println(time.Since(start))
	// 2392 too low
}
