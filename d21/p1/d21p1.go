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

type Monkey struct {
	name   string
	n      int
	op     string
	m1, m2 string
}

func getMonkeyYell(name string, monkeys map[string]Monkey) int {
	m := monkeys[name]
	if m.n != 0 {
		return m.n
	} else {
		m1, m2 := 0, 0
		// need to get number
		if len(m.m1) > 0 {
			m1 = getMonkeyYell(m.m1, monkeys)
		}
		if len(m.m2) > 0 {
			m2 = getMonkeyYell(m.m2, monkeys)
		}

		if m1 > 0 && m2 > 0 {
			switch m.op {
			case "+":
				return m1 + m2
			case "-":
				return m1 - m2
			case "*":
				return m1 * m2
			case "/":
				return m1 / m2
			}
		}
	}

	return -1
}

func getN(monkeys map[string]Monkey) {
	for _, m := range monkeys {
		if m.n == 0 {
			n := getMonkeyYell(m.name, monkeys)
			m.n = n
			monkeys[m.name] = m
		}
	}
}

func main() {
	puzzle, _ := readFile("d21/puzzle_input.txt")

	monkeys := make(map[string]Monkey, 0)
	for _, v := range puzzle {
		parts := strings.Split(v, ": ")
		name := parts[0]

		parts = strings.Split(parts[1], " ")
		var n int
		if len(parts) > 2 {
			monkey1 := parts[0]
			monkey2 := parts[2]
			op := parts[1]

			if _, ok := monkeys[name]; ok {
				m := monkeys[name]
				m.op = op
				m.m1 = monkey1
				m.m2 = monkey2
				monkeys[name] = m
			} else {
				m := Monkey{
					name: name,
					op:   op,
					m1:   monkey1,
					m2:   monkey2,
				}
				monkeys[name] = m
			}

		} else {
			n, _ = strconv.Atoi(parts[len(parts)-1])

			if _, ok := monkeys[name]; ok {
				m := monkeys[name]
				m.n = n

				monkeys[name] = m
			} else {

				m := Monkey{
					name: name,
					n:    n,
				}
				monkeys[name] = m
			}
		}

	}

	getN(monkeys)
	fmt.Println(monkeys["root"].n)

}
