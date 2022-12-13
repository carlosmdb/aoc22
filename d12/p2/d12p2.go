package main

import (
	"bufio"
	"os"
	"sort"
)

type position struct {
	y, x int
}

var (
	moves = []position{
		{x: 0, y: 1},  // down
		{x: 1, y: 0},  // right
		{x: -1, y: 0}, // left
		{x: 0, y: -1}, // up
	}
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

func main() {
	puzzle, _ := readFile("d12/puzzle_input.txt")
	area := make([][]rune, 0)

	start := position{}
	end := position{}

	starts := make([]position, 0)

	for y, line := range puzzle {
		row := make([]rune, 0)
		for x, v := range line {
			if v == 'a' {
				starts = append(starts, position{y, x})
			} else if v == 'S' {
				start = position{y, x}
				v = 'a'
				starts = append(starts, start)
			} else if v == 'E' {
				end = position{y, x}
				v = 'z'
			}
			row = append(row, v)
		}
		area = append(area, row)
	}

	result := bfs(area, start, end)
	println("part 1:", result)

	results := make([]int, 0)
	for _, start := range starts {
		result = bfs(area, start, end)
		if result != -1 {
			results = append(results, result)
		}
	}

	sort.Ints(results)
	println("part 2:", results[0])
}

func getNeighbours(p position, area [][]rune) []position {
	var pos []position

	for _, m := range moves {
		newPos := position{y: p.y + m.y, x: p.x + m.x}

		if newPos.y >= 0 && newPos.y < len(area) && newPos.x >= 0 && newPos.x < len(area[0]) {
			// found neighbour

			// the destination square can be at most one higher than the elevation of your current square
			if area[newPos.y][newPos.x] <= area[p.y][p.x]+1 {
				pos = append(pos, newPos)
			}

		}
	}

	return pos
}

func bfs(area [][]rune, start position, end position) int {
	queue := []position{start}
	visited := make(map[position]int, 0)

	visited[start] = 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr == end {
			return visited[end]
		}

		n := getNeighbours(curr, area)
		for _, next := range n {

			// if haven't visited yet
			if _, ok := visited[next]; !ok {
				// store distance to the previous position
				visited[next] = visited[curr] + 1
				queue = append(queue, next)
			}

		}
	}

	return -1
}
