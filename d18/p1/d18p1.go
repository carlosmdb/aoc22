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

type Cube struct {
	x, y, z int
}

func main() {
	puzzle, _ := readFile("d18/puzzle_input.txt")

	cubes := make([]Cube, 0)
	for _, v := range puzzle {
		fmt.Println(v)
		parts := strings.Split(v, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])

		c := Cube{x: x, y: y, z: z}
		cubes = append(cubes, c)
	}

	fmt.Println(cubes)

}
