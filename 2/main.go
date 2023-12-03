package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Which games are possible if max 12 red, 13 green, 14 blue

func main() {
	// Parse input from filesystem
	// =============================================================================
	f, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("Could not read from file: %v", err)
	}
	input := string(f)
	sum, pows := sumOfIdsFromValidGames(input)

	fmt.Println("Solution - part 1:", sum)
	fmt.Println("Solution - part 1:", pows)
}

func sumOfIdsFromValidGames(inputString string) (int, int) {
	var sum int
	var powSum int

	m := make(map[string][]string) // [Game number]: Slice of games played with result of each game
	for _, l := range strings.Split(string(inputString), "\n") {
		if l == "" {
			continue
		}
		parts := strings.Split(l, ":")
		m[parts[0]] = strings.Split(parts[1], ";")
	}

	for game, v := range m {
		gameIsValid := true

		rs := make(map[string]int) // [color]: highest number of cubes needed for color in game (part 2)

		// New Game
		for _, rounds := range v {
			round := strings.Split(rounds, ",")
			r := make(map[string]int) // [color]: amount

			// Each round in game
			for _, v := range round {
				parts := strings.Split(strings.TrimSpace(v), " ")
				amount, _ := strconv.Atoi(parts[0])
				color := parts[1]

				r[color] += int(amount)
			}

			if red, green, blue := r["red"], r["green"], r["blue"]; red > 12 || green > 13 || blue > 14 {
				gameIsValid = false
			}

			// If there is a new highest number of cubes, assign it to rs
			for k, v := range r {
				if v > rs[k] {
					rs[k] = v
				}
			}

		}

		// Get power of each round and sum it
		pows := 1
		for _, v := range rs {
			pows *= v
		}
		powSum += pows

		if gameIsValid {
			gameNo, _ := strconv.Atoi(strings.Split(game, " ")[1])
			sum += gameNo
		}
	}

	return sum, powSum
}
