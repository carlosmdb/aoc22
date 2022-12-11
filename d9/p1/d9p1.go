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

type Pos struct {
	x, y int
}

type Rope struct {
	head Pos
	tail Pos
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	puzzle, _ := readFile("d9/puzzle_input.txt")

	m := make(map[Pos]struct{}, 0)
	rope := Rope{
		head: Pos{x: 0, y: 0},
		tail: Pos{x: 0, y: 0},
	}
	m[rope.tail] = struct{}{}

	for _, line := range puzzle {
		fmt.Println("value is", line)
		move := strings.Split(line, " ")[0]
		n, _ := strconv.Atoi(string(strings.Split(line, " ")[1]))

		for i := 0; i < n; i++ {
			oldPos := rope.head
			if move == "U" {
				rope.head.x++

				if rope.head.x > rope.tail.x+1 {
					rope.tail.x++
					if rope.head.x != rope.tail.x || rope.head.y != rope.tail.y {
						rope.tail = oldPos
					}
				}

			} else if move == "D" {
				rope.head.x--

				if rope.head.x < rope.tail.x-1 {
					rope.tail.x--
					if rope.head.x != rope.tail.x || rope.head.y != rope.tail.y {
						rope.tail = oldPos
					}
				}

			} else if move == "L" {
				rope.head.y--
				if rope.head.y < rope.tail.y-1 {
					rope.tail.y--
					if rope.head.x != rope.tail.x || rope.head.y != rope.tail.y {
						rope.tail = oldPos
					}
				}

			} else if move == "R" {
				rope.head.y++

				if rope.head.y > rope.tail.y+1 {
					rope.tail.y++
					if rope.head.x != rope.tail.x || rope.head.y != rope.tail.y {
						rope.tail = oldPos
					}
				}

			}

			m[rope.tail] = struct{}{}

		}

	}

	//fmt.Println(m)
	fmt.Println(len(m))
}
