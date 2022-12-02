package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("d1/puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	elf := 1
	max_elf := 1
	calories := 0
	max_calories := 0
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
			if calories > max_calories {
				max_calories = calories
				max_elf = elf
			}
			elf++
			calories = 0
		}
	}

	fmt.Println("Max calories is", max_calories)
	fmt.Println("This was elf number", max_elf)
}
