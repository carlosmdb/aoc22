package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
2
1
0
-
=

20 = 10
2= = 8

2=-01

*/

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

func encode(num int) string {

	code := "=-012"
	result := ""
	for num > 0 {
		rem := (num + 2) % 5
		result = string(code[rem]) + result

		num = (num + 2) / 5
	}

	return result
}

func decode(v string) int {
	mul := 1
	total := 0
	for i := len(v) - 1; i >= 0; i-- {

		val := 0
		switch string(v[i]) {
		case "0":
			val = 0
		case "1":
			val, _ = strconv.Atoi(string(v[i]))
		case "2":
			val, _ = strconv.Atoi(string(v[i]))
		case "-":
			val = -1
		case "=":
			val = -2
		}

		total += mul * val

		if mul == 1 {
			mul = 5
		} else {
			mul *= 5
		}
	}

	return total
}

func main() {
	puzzle, _ := readFile("d25/puzzle_input.txt")

	total := 0
	for _, v := range puzzle {
		total += decode(v)
	}
	fmt.Println("total is", total)

	fmt.Println(encode(total))
}
