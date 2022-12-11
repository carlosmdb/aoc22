package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

type Tree struct {
	v bool
	h int
	p Pos
}

func countVisibleTrees(m map[Pos]Tree) int {
	totaVisibile := 0
	for _, v := range m {
		if v.v {
			totaVisibile++
		}
	}

	return totaVisibile
}

func updateVisibility(m map[Pos]Tree, x int, y int) map[Pos]Tree {
	for row := 1; row < x-1; row++ {
		for column := 1; column < y-1; column++ {
			p := Pos{x: row, y: column}
			tree := m[p]
			t, b, l, r := true, true, true, true
			//fmt.Println(m[p])

			// check columns
			for a := 1; a < x; a++ {
				// top
				top := Pos{x: p.x - a, y: p.y}

				if m[top].h >= m[p].h {
					t = false
				}
			}

			for a := 1; a < x; a++ {
				// bottom
				bottom := Pos{x: p.x + a, y: p.y}
				if m[bottom].h >= m[p].h {
					b = false
				}
			}

			// check rows
			for a := 1; a < y; a++ {
				// left
				left := Pos{x: p.x, y: p.y - a}
				if m[left].h >= m[p].h {
					l = false
				}
			}

			for a := 1; a < y; a++ {
				// right
				right := Pos{x: p.x, y: p.y + a}
				if m[right].h >= m[p].h {
					r = false
				}
			}

			if !t && !b && !l && !r {
				tree.v = false
				m[p] = tree
			}
		}
	}

	return m
}

func main() {
	puzzle, _ := readFile("d8/puzzle_input.txt")

	m := make(map[Pos]Tree, 0)
	x, y := 0, 0

	for row, line := range puzzle {
		//fmt.Println("value is", line)
		x++
		y = 0
		for column, v := range line {
			p := Pos{x: row, y: column}
			v, _ := strconv.Atoi(string(v))
			m[p] = Tree{
				p: p,
				h: v,
				v: true,
			}
			y++
		}

	}

	m = updateVisibility(m, x, y)
	//fmt.Println(m)
	fmt.Println("Total visible", countVisibleTrees(m))
}
