package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	starting       []int
	operation      string
	operationValue int
	testValue      int
	trueRecipient  int
	falseRecipient int
	inspections    int
}

func convertToInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func ReadFile(path string) *bufio.Scanner {
	readFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return fileScanner
}

func parseStarters(numbers []string) []int {
	var starters []int

	// values := strings.ReplaceAll(numbers, ",", "")
	// starterStrings = strings.Fields(values)
	for _, v := range numbers {
		number := strings.Replace(v, ",", "", -1)
		starters = append(starters, convertToInt(number))
	}
	return starters
}

func createMonkey(instructions []string) *Monkey {
	i := instructions
	monkey := &Monkey{}
	for k, v := range i {
		values := strings.Fields(v)
		switch k {
		case 0:
			continue
		case 1:
			monkey.starting = parseStarters(values[2:])
		case 2:
			monkey.operation = values[4]
			monkey.operationValue = convertToInt(values[5])
		case 3:
			monkey.testValue = convertToInt(values[3])
		case 4:
			monkey.trueRecipient = convertToInt(values[5])
		case 5:
			monkey.falseRecipient = convertToInt(values[5])
		}
	}
	monkey.inspections = 0
	return monkey
}

func performOperation(start int, op string, opVal int) int {
	if opVal == 0 {
		opVal = start
	}

	if op == "*" {
		return start * opVal
	} else {
		return start + opVal
	}
}

func checkRemainder(l int, r int) bool {
	return math.Mod(float64(l), float64(r)) == 0
}

func (m *Monkey) monkeyBusiness(monkeys []*Monkey) {
	for _, v := range m.starting {
		worryLevel := performOperation(v, m.operation, m.operationValue)
		worryLevel /= 437
		if checkRemainder(worryLevel, m.testValue) {
			monkeys[m.trueRecipient].starting = append(monkeys[m.trueRecipient].starting, worryLevel)
		} else {
			monkeys[m.falseRecipient].starting = append(monkeys[m.falseRecipient].starting, worryLevel)
		}
		m.starting = m.starting[:len(m.starting)-1]
		m.inspections += 1
	}
}

func main() {

	input := ReadFile("config/test_input.txt")
	monkeyInstructions := []string{"", "", "", "", "", "", ""}
	var monkeys []*Monkey
	i := 0
	for input.Scan() {
		monkeyInstructions[i] = input.Text()
		i += 1
		if i == 6 {
			monkeys = append(monkeys, createMonkey(monkeyInstructions))
		}

		if i == 7 {
			i = 0
			continue
		}

	}
	fmt.Println(monkeys)
	for x := 0; x < 1000; x++ {
		for _, m := range monkeys {
			m.monkeyBusiness(monkeys)
		}

	}
	for _, v := range monkeys {
		fmt.Println(v)
	}
}
