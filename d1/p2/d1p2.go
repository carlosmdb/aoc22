package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func compare(a [3]int, b int) [3]int {
	sort.Ints(a[:3])

	if b > a[0] {
		a[0] = b
	}
	return a
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func main() {
	file, err := os.Open("d1/puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	elf := 1
	max_calories := [3]int{0, 0, 0}
	calories := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			value, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			//fmt.Println("value is ", value)
			calories += value

		} else {
			fmt.Printf("total for elf %v is %v\n", elf, calories)
			max_calories = compare(max_calories, calories)
			calories = 0
			elf++
		}
	}

	fmt.Println("Max calories is", max_calories)
	fmt.Println("Total calories is", sum(max_calories[:]))
}
