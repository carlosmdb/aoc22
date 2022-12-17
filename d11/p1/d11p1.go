package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	worryLevels                               []int
	operand                                   int
	operation                                 string
	testTrue, testFalse, test, itemsInspected int
}

func (monkey *Monkey) Dequeue() int {
	worryLevel := monkey.worryLevels[0]
	monkey.worryLevels = monkey.worryLevels[1:]

	return worryLevel
}

func (monkey *Monkey) Enqueue(worryLevel int) {
	monkey.worryLevels = append(monkey.worryLevels, worryLevel)
}

func (monkey *Monkey) Print(i int) {
	fmt.Printf("Monkey %v:", i)
	for _, v := range monkey.worryLevels {
		fmt.Printf("%v,", v)
	}
	fmt.Println()
}

func PrintItemsInspected(monkeys []*Monkey) {

	for index, v := range monkeys {
		fmt.Printf("Monkey %v inspected items %v times.\n", index, v.itemsInspected)
	}

}

func parseWorryLevels(s string) []int {
	parts := strings.Split(s, "  Starting items: ")
	parts = strings.Split(parts[1], ", ")

	worryLevels := make([]int, 0)
	for _, v := range parts {
		worryLevel, _ := strconv.Atoi(v)
		worryLevels = append(worryLevels, worryLevel)
	}
	return worryLevels
}

func parseOperation(s string) (string, int) {
	parts := strings.Split(s, "  Operation: new = old ")
	parts = strings.Split(parts[1], " ")
	operand, _ := strconv.Atoi(parts[1])

	return parts[0], operand
}

func runOperation(op string, operand int, worryLevel int) int {
	if operand == 0 {
		operand = worryLevel
	}

	switch op {
	case "+":
		return (worryLevel + operand) / 3
	case "-":
		return (worryLevel - operand) / 3
	case "*":
		return (worryLevel * operand) / 3
	case "/":
		return (worryLevel / operand) / 3
	}

	return -1
}

func main() {
	puzzle, _ := readFile("d11/puzzle_input.txt")

	monkeys := make([]*Monkey, 0)
	for i := 1; i < len(puzzle)-4; i = i + 7 {
		worryLevels := parseWorryLevels(puzzle[i])
		operation, operand := parseOperation(puzzle[i+1])
		test, _ := strconv.Atoi(strings.Split(puzzle[i+2], "  Test: divisible by ")[1])
		testTrue, _ := strconv.Atoi(strings.Split(puzzle[i+3], "   If true: throw to monkey ")[1])
		testFalse, _ := strconv.Atoi(strings.Split(puzzle[i+4], "   If false: throw to monkey ")[1])

		monkey := Monkey{
			worryLevels: worryLevels,
			operation:   operation,
			operand:     operand,
			test:        test,
			testTrue:    testTrue,
			testFalse:   testFalse,
		}

		monkeys = append(monkeys, &monkey)
	}

	rounds := 20
	for i := 0; i < rounds; i++ {

		for _, monkey := range monkeys {

			for len(monkey.worryLevels) > 0 {
				oldWorryLevel := monkey.Dequeue()
				worryLevel := runOperation(monkey.operation, monkey.operand, oldWorryLevel)

				if worryLevel%monkey.test == 0 {
					//is divisible
					monkeys[monkey.testTrue].Enqueue(worryLevel)

				} else {
					//is not divisible
					monkeys[monkey.testFalse].Enqueue(worryLevel)
				}

				monkey.itemsInspected++
			}
		}

		fmt.Printf("After round %v, the monkeys are holding items with these worry levels:\n", i+1)
		for i, monkey := range monkeys {
			monkey.Print(i)
		}
	}

	PrintItemsInspected(monkeys)
	itemsInpected := make([]int, 0)
	for _, v := range monkeys {
		itemsInpected = append(itemsInpected, v.itemsInspected)
	}
	sort.Ints(itemsInpected)
	fmt.Println("Monkey business ", itemsInpected[len(itemsInpected)-1]*itemsInpected[len(itemsInpected)-2])
}
