package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type folder struct {
	name     string
	size     int
	parent   *folder
	children map[string]*folder
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

func splitLine(line string) []string {
	return strings.Fields(line)
}

func (f *folder) SizeModify(s int) {
	f.size += s
}

func (f *folder) ChildrenModify(fol *folder) {
	f.children[fol.name] = fol
}

func recursiveAdd(s int, f *folder) {

	if f.parent == nil {
		f.SizeModify(s)
		return
	} else {
		f.SizeModify(s)
		recursiveAdd(s, f.parent)
	}
}

func traverseValues(s *folder, i int, f *folder) *folder {
	if len(s.children) != 0 {
		for _, child := range s.children {
			f = traverseValues(child, i, f)
		}
	}

	if s.size >= i {
		if f == nil {
			return s
		} else if f.size > s.size {
			return s
		}
	}
	return f
}

func main() {

	input := ReadFile("config/real_input.txt")
	dirs := make(map[string]*folder)
	dirs["/"] = &folder{name: "/", size: 0, parent: nil, children: make(map[string]*folder)}
	var cwd = dirs["/"]
	for input.Scan() {
		content := splitLine(input.Text())
		switch content[0] {

		case "$":
			if content[1] == "cd" && content[2] == ".." {
				cwd = cwd.parent
			} else if content[1] == "cd" && content[2] != "/" {
				cwd = cwd.children[content[2]]
			}

		case "dir":
			// if it's a directory, create it inside our current directory, set its parent to the current one
			f := &folder{name: content[1], size: 0, parent: cwd, children: make(map[string]*folder)}
			cwd.ChildrenModify(f)

		default:
			// if it's a file, add its size to itself and recursively add it upwards towards root
			fileSize, _ := strconv.Atoi(content[0])
			recursiveAdd(fileSize, cwd)
		}
	}
	neededSpace := 30000000 - (70000000 - dirs["/"].size)
	smallest := traverseValues(dirs["/"], neededSpace, nil)
	fmt.Println(smallest.size)
}
