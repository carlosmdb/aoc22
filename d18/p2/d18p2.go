package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

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

type Volcano struct {
	size, sizeMin, sizeMax Cube
	cubes                  []Cube
	m                      map[Cube]struct{}
}

type Cube struct {
	x, y, z int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var (
	moves = []Cube{
		{x: 1, y: 0, z: 0},
		{x: -1, y: 0, z: 0},
		{x: 0, y: 1, z: 0},
		{x: 0, y: -1, z: 0},
		{x: 0, y: 0, z: 1},
		{x: 0, y: 0, z: -1},
	}
)

func getNeighbours(c Cube, v Volcano) []Cube {
	var neighbours []Cube

	for _, mov := range moves {
		next := Cube{x: c.x + mov.x, y: c.y + mov.y, z: c.z + mov.z}

		// cube is within bounds
		if next.x >= v.sizeMin.x-v.size.x &&
			next.x <= v.sizeMax.x+v.size.x &&
			next.y >= v.sizeMin.y-v.size.y &&
			next.y <= v.sizeMax.y+v.size.y &&
			next.z >= v.sizeMin.z-v.size.z &&
			next.z <= v.sizeMax.z+v.size.z {
			neighbours = append(neighbours, next)
		}
	}

	return neighbours
}

func bfs(cube Cube, v Volcano) int {
	total := 0
	visited := make(map[Cube]struct{}, 0)
	queue := []Cube{cube}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// get neighbours
		n := getNeighbours(curr, v)
		for _, next := range n {

			// not visited yet
			if _, ok := visited[next]; !ok {
				// cube exists
				if _, ok := v.m[next]; ok {
					// not an air pocket
					total++
				} else {
					// add to visited
					visited[next] = struct{}{}
					queue = append(queue, next)
				}
			}
		}
	}

	return total
}

func main() {
	puzzle, _ := readFile("d18/puzzle_input.txt")

	size := Cube{x: 1, y: 1, z: 1}
	sizeMax := Cube{x: 0, y: 0, z: 0}
	sizeMin := Cube{x: MaxInt, y: MaxInt, z: MaxInt}
	cubes := make([]Cube, 0)
	m := make(map[Cube]struct{}, 0)
	for _, v := range puzzle {
		parts := strings.Split(v, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])

		if x < sizeMin.x {
			sizeMin.x = x
		}
		if y < sizeMin.y {
			sizeMin.y = y
		}
		if z < sizeMin.z {
			sizeMin.z = z
		}

		if x > sizeMax.x {
			sizeMax.x = x
		}
		if y > sizeMax.y {
			sizeMax.y = y
		}
		if z > sizeMax.z {
			sizeMax.z = z
		}

		c := Cube{x: x, y: y, z: z}
		cubes = append(cubes, c)
		m[c] = struct{}{}
	}

	v := Volcano{
		size: size, sizeMin: sizeMin, sizeMax: sizeMax,
		cubes: cubes,
		m:     m,
	}

	total := bfs(Cube{x: 0, y: 0, z: 0}, v)
	println(total)

}
