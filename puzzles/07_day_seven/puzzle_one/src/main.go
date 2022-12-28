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

func traverseValues(s *folder, res int) int {

	fmt.Printf("Traversing folders: %s\n", s.name)

	if len(s.children) != 0 {
		for _, child := range s.children {
			res = traverseValues(child, res)
		}
	}

	if s.size <= 100000 {
		res += s.size
	}
	return res
}

func main() {

	input := ReadFile("config/real_input.txt")
	dirs := make(map[string]*folder)
	dirs["/"] = &folder{name: "/", size: 0, parent: nil, children: make(map[string]*folder)}
	var cwd = dirs["/"]

	for input.Scan() {
		content := splitLine(input.Text())
		//fmt.Println(content)
		switch content[0] {

		case "$":
			if content[1] == "cd" && content[2] == ".." {

				//fmt.Printf("Current directory is %s\n", cwd.name)
				cwd = cwd.parent
			} else if content[1] == "cd" && content[2] != "/" {

				//fmt.Printf("Current directory is %s\n", cwd.name)
				cwd = cwd.children[content[2]]
			}

		case "dir":
			// if it's a directory, create it inside our current directory, set its parent to the current one
			//fmt.Printf("Creating directory %s with parent %s\n", content[1], cwd.name)
			f := &folder{name: content[1], size: 0, parent: cwd, children: make(map[string]*folder)}
			cwd.ChildrenModify(f)

		default:
			// if it's a file, add its size to itself and recursively add it upwards towards root
			//fmt.Printf("Beginning recursive file add\n")
			fileSize, _ := strconv.Atoi(content[0])
			recursiveAdd(fileSize, cwd)
		}
	}

	res := traverseValues(dirs["/"], 0)
	fmt.Println(res)
}
