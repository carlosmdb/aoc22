package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(fileName string) ([]string, error) {
	output := make([]string, 0)
	file, err := os.Open(fileName)
	if err != nil {
		return output, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	err = file.Close()
	if err != nil {
		return output, err
	}
	return output, nil
}

const (
	DIR  = "dir"
	FILE = "file"
)

type object struct {
	name      string
	size      int
	objType   string
	children  []*object
	parent    *object
	depth     int
	totalSize int
}

func getTotalSizes(root *object) int {
	queue := make([]*object, 0)
	queue = append(queue, root)
	total := 0
	for len(queue) > 0 {
		currentObj := queue[0]
		queue = queue[1:]
		if currentObj.objType != FILE {
			currentObj.totalSize = calculateDirSize(currentObj)
			if currentObj.totalSize <= 100000 {
				total += currentObj.totalSize
			}
		}
		if len(currentObj.children) > 0 {
			queue = append(queue, currentObj.children...)
		}
	}

	return total
}

func calculateDirSize(root *object) int {
	queue := make([]*object, 0)
	queue = append(queue, root.children...)

	totalSize := root.totalSize
	for len(queue) > 0 {
		currentObj := queue[0]
		queue = queue[1:]

		if currentObj.objType != FILE {
			totalSize += calculateDirSize(currentObj)
		}
	}
	return totalSize
}

func getDirObject(root *object, name string) *object {
	for _, v := range root.children {
		if v.name == name && v.objType == DIR {
			return v
		}
	}
	return nil
}

func printFilesystem(root *object) {
	queue := make([]*object, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		currentObj := queue[0]
		queue = queue[1:]

		for i := 0; i <= currentObj.depth; i++ {
			fmt.Printf(" ")
		}

		if currentObj.objType == FILE {
			fmt.Printf("- %v (%v, size=%v)\n", currentObj.name, currentObj.objType, currentObj.size)
		} else {
			fmt.Printf("- %v (%v, totalSize=%v)\n", currentObj.name, currentObj.objType, currentObj.totalSize)
		}

		if len(currentObj.children) > 0 {
			queue = append(queue, currentObj.children...)
		}
	}
}

func main() {
	puzzle, _ := readFile("d7/puzzle_input.txt")

	root := object{
		name:    "/",
		objType: DIR,
	}

	currentDir := &root
	depth := 1

	for _, line := range puzzle {
		//fmt.Println("value is", line)

		if line[0] == '$' {

			if line == "$ cd /" {
				//fmt.Println("go to the root")
				currentDir = &root
				depth = 1
			} else if line == "$ cd .." {
				//fmt.Println("go to parent")
				if currentDir.parent != nil {
					currentDir = currentDir.parent
					depth--
				} else {
					fmt.Println("trying to go up from root!!")
				}
			} else {
				if strings.Contains(line, "cd") {
					dirName := strings.Split(line, " ")[2]
					//fmt.Println("go to dir", dirName)
					currentDir = getDirObject(currentDir, dirName)
					depth++
				} else {
					//fmt.Println("ls")
				}
			}
		} else {
			if strings.Contains(line, "dir") {
				//fmt.Println("directory")
				dir := object{
					name:    strings.Split(line, " ")[1],
					objType: DIR,
					parent:  currentDir,
					depth:   depth,
				}
				currentDir.children = append(currentDir.children, &dir)
			} else {
				//fmt.Println("file")
				size, _ := strconv.Atoi(strings.Split(line, " ")[0])
				file := object{
					name:    strings.Split(line, " ")[1],
					objType: FILE,
					size:    size,
					parent:  currentDir,
					depth:   depth,
				}
				currentDir.totalSize += size
				currentDir.children = append(currentDir.children, &file)
			}
		}
	}

	printFilesystem(&root)
	fmt.Println(getTotalSizes(&root))
}
