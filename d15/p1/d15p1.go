package main

import (
	"bufio"
	"fmt"
	"os"
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distance(a, b pos) int {
	return Abs(a.x-b.x) + Abs(a.y-b.y)
}

type pos struct {
	y, x int
}

type min pos
type max pos

type size struct {
	min min
	max max
}

type signal struct {
	sensor, beacon pos
	distance       int
}

func processSignal(s string, m map[pos]string, area *size) signal {

	x1, x2, y1, y2 := 0, 0, 0, 0

	fmt.Sscanf(s, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &x1, &y1, &x2, &y2)

	sensor := pos{
		x: x1,
		y: y1,
	}
	beacon := pos{
		x: x2,
		y: y2,
	}
	m[sensor] = "S"
	m[beacon] = "B"

	dist := distance(sensor, beacon)
	signal := signal{
		sensor:   sensor,
		beacon:   beacon,
		distance: dist,
	}

	if (x1 - dist) < area.min.x {
		area.min.x = (x1 - dist)
	} else if (x1 + dist) > area.max.x {
		area.max.x = (x2 + dist)
	}

	if (x2 - dist) < area.min.x {
		area.min.x = (x2 - dist)
	} else if (x2 + dist) > area.max.x {
		area.max.x = (x2 + dist)
	}

	if (y1 - dist) < area.min.y {
		area.min.y = (y1 - dist)
	} else if (y1 + dist) > area.max.y {
		area.max.y = (y1 + dist)
	}

	if (y2 - dist) < area.min.y {
		area.min.y = (y2 - dist)
	} else if (y2 + dist) > area.max.y {
		area.max.y = (y2 + dist)
	}

	return signal

}

/* func buildMap(m map[pos]string, area *size) {
	for i := area.min.y; i < area.max.y; i++ {
		for j := area.min.x; j < area.max.x; j++ {

			if _, ok := m[pos{i, j}]; !ok {
				m[pos{i, j}] = "."
			}
		}
	}

}

func printMap(m map[pos]string, area *size) {

	fmt.Println()
	for i := area.min.y; i < area.max.y; i++ {
		for j := area.min.x; j < area.max.x; j++ {

			if _, ok := m[pos{i, j}]; !ok {
				m[pos{i, j}] = "."
			}
			fmt.Printf("%v", m[pos{i, j}])

		}
		fmt.Println()
	}

} */

func main() {
	puzzle, _ := readFile("d15/puzzle_input.txt")

	signals := make([]signal, 0)
	m := make(map[pos]string, 0)
	area := size{}

	for _, v := range puzzle {
		signal := processSignal(v, m, &area)
		signals = append(signals, signal)
	}

	//buildMap(m, &area)
	//printMap(m, &area)

	count := 0
	i := 2000000
	for j := area.min.x; j < area.max.x; j++ {
		p := pos{
			y: i,
			x: j,
		}

		for _, v := range signals {
			if distance(v.sensor, p) <= v.distance {
				if m[p] != "B" && m[p] != "S" {
					count++
					break
				}
			}
		}

	}
	fmt.Println(count)
}
