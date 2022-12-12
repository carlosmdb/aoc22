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

type CRT struct {
	h, v   int
	screen [6][40]string
}

func (c *CRT) DrawPixel(x int) {
	// if current pixel is contained in the sprite
	if c.h >= x-1 && c.h <= x+1 {
		c.screen[c.v][c.h] = "#"
	} else {
		c.screen[c.v][c.h] = "."
	}
	c.h++

	if c.h == 40 {
		c.h = 0
		c.v++
	}
}

func (c *CRT) Print() {
	for _, v := range c.screen {
		for _, h := range v {
			fmt.Printf("%v", h)
		}
		fmt.Printf("\n")
	}

}

func main() {
	puzzle, _ := readFile("d10/puzzle_input.txt")

	crt := CRT{}

	cycles := make([]int, 0)
	x := 1
	cycles = append(cycles, x)

	for _, line := range puzzle {
		//fmt.Println("value is", line)

		if strings.Contains(line, "noop") {
			crt.DrawPixel(x)
			cycles = append(cycles, x)

		} else if strings.Contains(line, "addx") {

			crt.DrawPixel(x)
			cycles = append(cycles, x)
			crt.DrawPixel(x)

			v, _ := strconv.Atoi(string(strings.Split(line, " ")[1]))
			x += v
			cycles = append(cycles, x)

		}
	}

	//fmt.Println(cycles)

	strength := 0
	for i := 19; i < len(cycles); i += 40 {
		strength += cycles[i] * (i + 1)
	}

	//fmt.Println(strength)

	crt.Print()

}
