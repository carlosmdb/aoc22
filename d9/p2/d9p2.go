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

const ropeLength = 10

type Pos struct {
	x, y int
}

type Knot struct {
	p Pos
}

type Rope struct {
	knots [ropeLength]Knot
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
		knots: [ropeLength]Knot{},
	}
	m[rope.knots[9].p] = struct{}{}

	for _, line := range puzzle {
		//fmt.Println("value is", line)
		move := strings.Split(line, " ")[0]
		n, _ := strconv.Atoi(string(strings.Split(line, " ")[1]))

		for i := 0; i < n; i++ {
			if move == "U" {
				rope.knots[0].p.x++
			} else if move == "D" {
				rope.knots[0].p.x--

			} else if move == "L" {
				rope.knots[0].p.y--

			} else if move == "R" {
				rope.knots[0].p.y++

			}

			for j := 0; j < ropeLength-1; j++ {

				touching := false
				for x := -1; x < 2 && !touching; x++ {
					for y := -1; y < 2 && !touching; y++ {
						reached := Pos{rope.knots[j+1].p.x + x, rope.knots[j+1].p.y + y}
						if reached == rope.knots[j].p {
							touching = true
						}
					}
				}

				if !touching {

					if rope.knots[j].p.x > rope.knots[j+1].p.x {
						rope.knots[j+1].p.x++
					} else if rope.knots[j].p.x < rope.knots[j+1].p.x {
						rope.knots[j+1].p.x--
					}

					if rope.knots[j].p.y > rope.knots[j+1].p.y {
						rope.knots[j+1].p.y++
					} else if rope.knots[j].p.y < rope.knots[j+1].p.y {
						rope.knots[j+1].p.y--
					}

				}

			}

			m[rope.knots[9].p] = struct{}{}

		}

	}

	fmt.Println(m)
	fmt.Println(len(m))
}
