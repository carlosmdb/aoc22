package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type group struct {
	elves        [3]rucksack
	common_value rune
}

type rucksack struct {
	first  []rune
	second []rune
	all    []rune
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

func print_tribe(tribe []group) {
	for i, v := range tribe {
		fmt.Println("group ", i+1)
		print_rucksacks(v.elves[:])

	}

}

func tribe_total(tribe []group) int {
	total := 0
	for _, v := range tribe {
		if int(v.common_value) <= 90 {
			total += int(v.common_value) - 38
		} else {
			total += int(v.common_value) - 96
		}
	}
	return total
}

func (team *group) find_group_common_value() {

	for _, v := range team.elves {
		sort.Slice(v.all, func(i, j int) bool {
			return v.all[i] < v.all[j]
		})
	}

	i := 0
	j := 0
	k := 0

	for ok := true; ok; ok = (i < len(team.elves[0].all)) && (j < len(team.elves[1].all)) && (k < len(team.elves[2].all)) {
		// If x = y and y = z, print any of them and move
		// ahead in all arrays
		if team.elves[0].all[i] == team.elves[1].all[j] && team.elves[1].all[j] == team.elves[2].all[k] {
			fmt.Printf("common element %v\n", string(team.elves[0].all[i]))
			team.common_value = team.elves[0].all[i]
			i++
			j++
			k++
		} else if team.elves[0].all[i] < team.elves[1].all[j] {
			i++
		} else if team.elves[1].all[j] < team.elves[2].all[k] {
			j++
		} else {
			k++
		}
	}
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
	puzzle, _ := readFile("d3/puzzle_input.txt")

	rucksacks := make([]rucksack, 0)
	tribe := make([]group, 0)

	team := group{}
	count := 0

	for _, line := range puzzle {
		//fmt.Println("value is ", line)
		ruck := rucksack{
			first:  []rune(line[:(len(line) / 2)]),
			second: []rune(line[(len(line) / 2):]),
			all:    []rune(line),
		}

		rucksacks = append(rucksacks,
			ruck,
		)

		team.elves[count] = ruck

		count++
		if count == 3 {
			// new group
			team.find_group_common_value()
			tribe = append(tribe, team)
			count = 0
		}

	}

	print_tribe(tribe)
	total := tribe_total(tribe)
	fmt.Println("total ", total)

}
