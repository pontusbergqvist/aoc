package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// Parse input from filesystem
	// =============================================================================

	inputFile, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("Could not read from input file: %v", err)
	}
	input := strings.Split(string(inputFile), "\n")

	// Part 1
	// =============================================================================

	ns := inputToNums(input)

	var sum1 int
	for _, v := range ns {
		sum1 += v
	}

	fmt.Println("Solution - part 1:", sum1)

	// Part 2
	// =============================================================================

	var sum2 int
	for _, v := range input {
		sum2 += findNumber(v)
	}

	fmt.Println("Solution - part 2:", sum2)
}

// ===============================================================================

func getFirstDigit(string string) rune {
	var digit rune
	for _, v := range string {
		if unicode.IsDigit(v) {
			digit = v - '0' // rune - '0': Transform the rune value from Unicode code point to UTF-8 character (still of type rune)
			return digit
		}
	}
	return digit
}

func getLastDigit(string string) rune {
	var digit rune
	r := []rune(string)
	for i := len(r) - 1; i >= 0; i-- {
		if unicode.IsDigit(r[i]) {
			digit = r[i] - '0'
			return digit
		}
	}
	return digit
}

func inputToNums(input []string) []int {
	var ns []int

	for _, s := range input {
		n1 := getFirstDigit(s)
		n2 := getLastDigit(s)

		n := fmt.Sprintf("%v%v", n1, n2)
		i, _ := strconv.Atoi(n)
		ns = append(ns, i)
	}
	return ns
}

// Part two

func findNumber(inputString string) int {
	var matches []string

	re := regexp.MustCompile(`^(\d|one|two|three|four|five|six|seven|eight|nine)`)

	for i := range inputString {
		match := re.FindString(inputString[i:])
		if match != "" {
			matches = append(matches, match)
		}
	}

	for i := range matches {
		matches[i] = wordToNum(matches[i])
	}

	if len(matches) == 0 {
		return 0
	}

	num, _ := strconv.Atoi(fmt.Sprintf("%v%v", matches[0], matches[len(matches)-1]))
	return num
}

func wordToNum(word string) string {
	switch word {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"

	default:
		return word
	}
}
