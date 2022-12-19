package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

type Edges struct {
	minX, maxX, minY, maxY int
}

type Point struct {
	y, x int
}

type Line struct {
	points []Point
}

type Lines []Line

type Cavern struct {
	m      map[Point]string
	edges  Edges
	source Point
	sand   []Point
}

func (l *Lines) GetEdges() Edges {
	var minX, maxX, minY, maxY int
	minY = 0
	minX = MaxInt
	for _, v := range *l {
		for _, v := range v.points {

			if v.x < minX {
				minX = v.x
			}
			if v.x > maxX {
				maxX = v.x
			}
			if v.y > maxY {
				maxY = v.y
			}
		}
	}

	return Edges{minX: minX, minY: minY, maxX: maxX, maxY: maxY}
}

func (l *Lines) GetPointType(p Point) string {

	for _, v := range *l {

		for i := 0; i < len(v.points)-1; i++ {
			if p.x >= v.points[i].x && p.x <= v.points[i+1].x && p.y >= v.points[i].y && p.y <= v.points[i+1].y {
				return "#"
			}

			if p.x >= v.points[i+1].x && p.x <= v.points[i].x && p.y >= v.points[i+1].y && p.y <= v.points[i].y {
				return "#"
			}
		}
	}

	return "."
}

func (l *Lines) Print() {
	edges := l.GetEdges()

	for y := edges.minY; y <= edges.maxY; y++ {
		for x := edges.minX; x <= edges.maxX; x++ {

			pointType := l.GetPointType(Point{x: x, y: y})
			fmt.Printf(pointType)
		}
		fmt.Printf("\n")
	}
}

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

func initMap(m map[Point]string, l *Lines, edges Edges) {

	for y := edges.minY; y <= edges.maxY; y++ {
		for x := edges.minX; x <= edges.maxX; x++ {
			point := Point{x: x, y: y}
			pointType := l.GetPointType(point)

			if y == edges.maxY {
				m[point] = "#"
			} else {
				m[point] = pointType
			}

		}

	}
}

func (c *Cavern) PrintMap() {
	fmt.Printf("    ")
	for x := c.edges.minX; x <= c.edges.maxX; x++ {
		fmt.Printf("%v ", x)
	}
	fmt.Printf("\n")

	for y := c.edges.minY; y <= c.edges.maxY; y++ {
		fmt.Printf(" %v ", y)
		for x := c.edges.minX; x <= c.edges.maxX; x++ {
			point := Point{x: x, y: y}
			fmt.Printf("  %v ", c.m[point])
		}
		fmt.Printf("\n")
	}
}

func (c *Cavern) MoveSand(sand Point) Point {
	moved := true
	restPos := sand

	for y := sand.y; y <= c.edges.maxY && moved; y++ {
		moved = false
		originalPos := sand
		// down
		sand.y++

		switch c.m[sand] {
		case "#":
		case "o":
		case ".":
			c.m[sand] = "o"
			restPos = sand
			c.m[originalPos] = "."
			moved = true
		}

		if !moved {
			// down and left
			sand.x--

			switch c.m[sand] {
			case "#":
			case "o":
			case ".":
				c.m[sand] = "o"
				restPos = sand
				c.m[originalPos] = "."
				moved = true
			}

			if !moved {
				// down and right
				sand.x += 2
				switch c.m[sand] {
				case "#":
				case "o":
				case ".":
					c.m[sand] = "o"
					restPos = sand
					c.m[originalPos] = "."
					moved = true
				}
			}
		}

	}
	return restPos
}

func (c *Cavern) AddSand() bool {
	sand := c.source
	c.m[sand] = "o"

	restPos := c.MoveSand(sand)
	if restPos == c.source {
		fmt.Println("reached the source", len(c.sand)+1)
		return false
	}
	c.sand = append(c.sand, restPos)
	//c.PrintMap()

	return true
}

func main() {
	puzzle, _ := readFile("d14/puzzle_input.txt")

	lines := Lines{}
	source := Point{
		x: 500,
		y: 0,
	}

	m := make(map[Point]string, 0)

	for _, v := range puzzle {
		points := strings.Split(v, " -> ")
		line := Line{}
		for _, point := range points {
			data := strings.Split(point, ",")
			x, _ := strconv.Atoi(data[0])
			y, _ := strconv.Atoi(data[1])

			line.points = append(line.points, Point{x: x, y: y})
		}
		lines = append(lines, line)
	}

	edges := lines.GetEdges()
	edges.maxY += 2
	edges.minX -= 1000
	edges.maxX += 1000
	initMap(m, &lines, edges)
	m[source] = "+"

	cavern := Cavern{
		m:      m,
		edges:  edges,
		source: source,
	}

	cont := true
	for ok := true; ok; ok = cont {
		cont = cavern.AddSand()
	}
}
