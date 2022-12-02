package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type play struct {
	a string
	b string
}

func main() {
	file, err := os.Open("d2/puzzle_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	points := make(map[string]int)
	points["A"] = 1
	points["B"] = 2
	points["C"] = 3

	points["X"] = 1
	points["Y"] = 2
	points["Z"] = 3

	winner := make(map[play]int)
	winner[play{a: "A", b: "X"}] = 0
	winner[play{a: "A", b: "Y"}] = 2
	winner[play{a: "A", b: "Z"}] = 1

	winner[play{a: "B", b: "X"}] = 1
	winner[play{a: "B", b: "Y"}] = 0
	winner[play{a: "B", b: "Z"}] = 2

	winner[play{a: "C", b: "X"}] = 2
	winner[play{a: "C", b: "Y"}] = 1
	winner[play{a: "C", b: "Z"}] = 0

	player1 := 0
	player2 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		value := strings.Split(line, " ")
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println("value is ", value)
		hand := ""

		for key, element := range winner {
			//fmt.Println("Key:", key, "=>", "value:", element)

			if value[1] == "X" {
				//lose
				if key.a == value[0] && element == 1 {
					hand = key.b
				}
			} else if value[1] == "Y" {
				//draw
				if key.a == value[0] && element == 0 {
					hand = key.b
				}

			} else if value[1] == "Z" {
				//win
				if key.a == value[0] && element == 2 {
					hand = key.b
				}

			}

		}

		//fmt.Println("points for opponent ", points[value[0]])
		//fmt.Println("points for me ", points[hand])
		//fmt.Println("winner is ", winner[play{a: value[0], b: hand}])

		if (winner[play{a: value[0], b: hand}] == 1) {
			//player 1 wins
			player1 += 6 + points[value[0]]
			player2 += points[hand]
		} else if (winner[play{a: value[0], b: hand}] == 2) {
			//player 2 wins
			player1 += points[value[0]]
			player2 += 6 + points[hand]
		} else if (winner[play{a: value[0], b: hand}] == 0) {
			//draw
			player1 += 3 + points[value[0]]
			player2 += 3 + points[hand]
		}

		fmt.Printf("player 1 score %v, player 2 score %v\n", player1, player2)
	}
}
