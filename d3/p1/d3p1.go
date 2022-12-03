package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type rucksack struct {
	first  []rune
	second []rune
}

func add_total(values [][]rune) int {
	total := 0
	for _, v := range values {
		if int(v[0]) <= 90 {
			total += int(v[0]) - 38
		} else {
			total += int(v[0]) - 96
		}
	}
	return total
}

func removeDuplicateInt(rucksacks [][]rune) [][]rune {
	complete_list := [][]rune{}
	for _, ruck := range rucksacks {
		allKeys := make(map[rune]bool)
		list := []rune{}
		for _, item := range ruck {
			if _, value := allKeys[item]; !value {
				allKeys[item] = true
				list = append(list, item)
			}
		}
		complete_list = append(complete_list, list)
	}

	return complete_list
}

func find_common_values(val []rucksack) [][]rune {
	common_values := make([][]rune, 0)
	for _, ruck := range val {
		common_value := make([]rune, 0)
		for _, first := range ruck.first {
			for _, second := range ruck.second {
				if first == second {
					common_value = append(common_value, first)
				}
			}
		}
		common_values = append(common_values, common_value)
	}
	fmt.Println("common values ", common_values)
	return common_values
}

func print_rucksacks(val []rucksack) {
	for _, v := range val {
		fmt.Printf("%v - %v\n", string(v.first), string(v.second))
	}
}

func list_rucksacks(val []rucksack) {
	for _, v := range val {
		fmt.Printf("%v - %v\n", v.first, v.second)
	}
}

func main() {
	file, err := os.Open("d3/puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rucksacks := make([]rucksack, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println("value is ", line)
		rucksacks = append(rucksacks,
			rucksack{
				first:  []rune(line[:(len(line) / 2)]),
				second: []rune(line[(len(line) / 2):]),
			},
		)

	}

	//print_rucksacks(rucksacks)
	//list_rucksacks(rucksacks)
	common_values := find_common_values(rucksacks)
	final_list := removeDuplicateInt(common_values)
	fmt.Println("final list ", final_list)

	total := add_total(final_list)
	fmt.Println("total ", total)
}
