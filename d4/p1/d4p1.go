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
	puzzle, _ := readFile("d4/puzzle_input.txt")

	contained := 0
	for _, line := range puzzle {
		//fmt.Println("value is ", line)

		pair := strings.Split(line, ",")
		pair1, pair2 := strings.Split(pair[0], "-"), strings.Split(pair[1], "-")

		a1, _ := strconv.Atoi(pair1[0])
		a2, _ := strconv.Atoi(pair1[1])
		b1, _ := strconv.Atoi(pair2[0])
		b2, _ := strconv.Atoi(pair2[1])

		fmt.Printf("%v-%v,%v-%v\n", a1, a2, b1, b2)

		if a1 >= b1 && a2 <= b2 {
			// pair 1 is contained in pair 2
			contained++
		} else if a1 <= b1 && a2 >= b2 {
			// pair 2 is contained in pair 1
			contained++
		}
	}

	println(contained)
}
