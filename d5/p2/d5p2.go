package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack []byte

func (s stack) Len() int     { return len(s) }
func (s stack) Empty() bool  { return len(s) == 0 }
func (s stack) Peek() byte   { return s[len(s)-1] }
func (s *stack) Push(b byte) { (*s) = append((*s), b) }
func (s *stack) Append(b []byte) {
	for _, v := range b {
		(*s) = append((*s), v)
	}
}

func (s *stack) Pop() byte {
	d := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return d
}
func reverseStack(s stack) stack {
	var newStack stack
	len := s.Len()

	for i := 0; i < len; i++ {
		newStack.Push(s.Pop())
	}

	return newStack
}

func (s stack) Print() {
	for _, v := range s {
		fmt.Printf("%v ", string(v))
	}
	println()
}

func printStacks(s []stack) {
	for _, v := range s {
		v.Print()
	}
	fmt.Println()
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

func main() {
	puzzle, _ := readFile("d5/puzzle_input.txt")

	nStacks := (len(puzzle[0]) + 1) / 4
	//fmt.Println("number of stacks is ", nStacks)

	stacks := make([]stack, nStacks)

	for _, line := range puzzle {
		//fmt.Println("value is ", line)

		if strings.Contains(line, "[") {
			stackId := 1
			for i := 1; i < len(line); i += 4 {
				if line[i] != ' ' {
					//fmt.Printf("'%v' goes in stack %v\n", string(line[i]), stackId)
					stacks[stackId-1].Push(line[i])
				}
				stackId++
			}
		} else if len(line) < 1 {
			for index, v := range stacks {
				stacks[index] = reverseStack(v)
			}
			//fmt.Println("Original stack:")
			printStacks(stacks)
		} else if strings.Contains(line, "move") {
			// moves
			entries := strings.Split(line, " ")

			nMoves, _ := strconv.Atoi(entries[1])
			source, _ := strconv.Atoi(entries[3])
			target, _ := strconv.Atoi(entries[5])

			//fmt.Printf("move %v from %v to %v\n", nMoves, source, target)

			source--
			target--
			nMoves = len(stacks[source]) - nMoves

			containers := stacks[source][nMoves:]
			stacks[target].Append(containers)
			stacks[source] = stacks[source][:nMoves]

			//printStacks(stacks)

		}
	}
	fmt.Printf("Answer is ")
	for _, v := range stacks {
		fmt.Printf("%v", string(v.Pop()))
	}

}
