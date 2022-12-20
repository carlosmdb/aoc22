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

type CubeSize struct {
	x, y, z int
}

type CubeFaces struct {
	xy1, xy2, yz1, yz2, xz1, xz2 bool
}

type Cube struct {
	x, y, z int
	conn    CubeFaces
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	puzzle, _ := readFile("d18/puzzle_input.txt")

	size := CubeSize{1, 1, 1}
	cubes := make([]Cube, 0)
	for _, v := range puzzle {
		parts := strings.Split(v, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])

		c := Cube{x: x, y: y, z: z}
		cubes = append(cubes, c)
	}

	for i := 0; i < len(cubes); i++ {
		for j := 0; j < len(cubes); j++ {
			if cubes[i] != cubes[j] {
				//compare x and y - only z moves
				if cubes[i].x == cubes[j].x && cubes[i].y == cubes[j].y && cubes[i].z == cubes[j].z+size.z {
					cubes[i].conn.xy1 = true
				}
				if cubes[i].x == cubes[j].x && cubes[i].y == cubes[j].y && cubes[i].z+size.z == cubes[j].z {
					cubes[i].conn.xy2 = true
				}

				//compare x and z - only y moves
				if cubes[i].x == cubes[j].x && cubes[i].z == cubes[j].z && cubes[i].y == cubes[j].y+size.y {
					cubes[i].conn.xz1 = true
				}
				if cubes[i].x == cubes[j].x && cubes[i].z == cubes[j].z && cubes[i].y+size.y == cubes[j].y {
					cubes[i].conn.xz2 = true
				}

				//compare y and z - only x moves
				if cubes[i].y == cubes[j].y && cubes[i].z == cubes[j].z && cubes[i].x == cubes[j].x+size.x {
					cubes[i].conn.yz1 = true
				}
				if cubes[i].y == cubes[j].y && cubes[i].z == cubes[j].z && cubes[i].x+size.x == cubes[j].x {
					cubes[i].conn.yz2 = true
				}

			}
		}
	}

	//fmt.Println(cubes)
	notConnected := 0
	for _, v := range cubes {
		if !v.conn.xy1 {
			notConnected++
		}
		if !v.conn.xy2 {
			notConnected++
		}
		if !v.conn.yz1 {
			notConnected++
		}
		if !v.conn.yz2 {
			notConnected++
		}
		if !v.conn.xz1 {
			notConnected++
		}
		if !v.conn.xz2 {
			notConnected++
		}
	}

	fmt.Println(notConnected)

}
