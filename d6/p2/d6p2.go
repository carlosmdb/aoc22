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

func removeFromMap(m map[string]int, pos int) map[string]int {
	for k, v := range m {
		if v <= pos {
			delete(m, k)
		}
	}
	return m
}

func main() {
	puzzle, _ := readFile("d6/puzzle_input.txt")

	for _, line := range puzzle {
		fmt.Println("value is", line)
		m := make(map[string]int)
		for i := 0; i < len(line); i++ {
			l := string(line[i])
			if _, ok := m[l]; !ok {
				m[l] = i
			} else {
				m = removeFromMap(m, m[l])
				m[l] = i
			}

			if len(m) == 14 {
				fmt.Println(m)
				fmt.Println("Found 14th letter at position", i+1)
				break
			}
		}
	}
}
