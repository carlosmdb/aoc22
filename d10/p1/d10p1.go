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

func main() {
	puzzle, _ := readFile("d10/puzzle_input.txt")

	cycles := make([]int, 0)
	x := 1
	cycles = append(cycles, x)
	for _, line := range puzzle {
		//fmt.Println("value is", line)

		if strings.Contains(line, "noop") {
			cycles = append(cycles, x)
		} else if strings.Contains(line, "addx") {
			cycles = append(cycles, x)
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

	fmt.Println(strength)

}
