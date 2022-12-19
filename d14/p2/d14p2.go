package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
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

func compareLists(leftList any, rightList any) int {

	leftArray, leftIsArray := leftList.([]any)
	rightArray, rightIsArray := rightList.([]any)

	// both are integers
	if !leftIsArray && !rightIsArray {
		return int(leftList.(float64) - rightList.(float64))
	} else if !leftIsArray {
		// left is integer and right is list
		leftArray = []any{leftList}
	} else if !rightIsArray {
		// left is list and right is integer
		rightArray = []any{rightList}
	}

	// both values are lists
	for i := 0; i < len(leftArray) && i < len(rightArray); i++ {
		// compare each element
		// if they are the same keep going, otherwise return
		if result := compareLists(leftArray[i], rightArray[i]); result != 0 {
			return result
		}
	}

	return len(leftArray) - len(rightArray)

}

func main() {
	puzzle, _ := readFile("d13/puzzle_input.txt")

	var div1 any
	var div2 any

	json.Unmarshal([]byte("[[2]]"), &div1)
	json.Unmarshal([]byte("[[6]]"), &div2)

	div1Index := 1
	div2Index := 2

	for i := 0; i < len(puzzle)-1; i += 3 {
		left := puzzle[i]
		right := puzzle[i+1]

		var leftList any
		var rightList any

		if err := json.Unmarshal([]byte(left), &leftList); err != nil {
			log.Fatal(err)
		}
		if err := json.Unmarshal([]byte(right), &rightList); err != nil {
			log.Fatal(err)
		}

		if compareLists(leftList, div1) <= 0 {
			div1Index++
		}
		if compareLists(leftList, div2) <= 0 {
			div2Index++
		}
		if compareLists(rightList, div1) <= 0 {
			div1Index++
		}
		if compareLists(rightList, div2) <= 0 {
			div2Index++
		}

	}

	fmt.Println(div1Index * div2Index)
}
