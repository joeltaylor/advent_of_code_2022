package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read input from input.txt
	input, _ := os.ReadFile("input.txt")
	fmt.Println(Part2(string(input)))
	fmt.Println(Part1(string(input)))
}

func Part1(input string) int {
	sum := 0
	grouped := GroupArray(strings.Split(input, "\n"), 3)

	for _, trio := range grouped {
		sum += SumInteresection(trio)
	}

	return sum
}

func Part2(input string) int {
	sum := 0

	for _, line := range strings.Split(input, "\n") {
		l := len(line) / 2

		sum += SumInteresection([]string{line[:l], line[l:]})
	}
	return sum
}

func SumInteresection(items []string) int {
	sum := 0
	characterCount := make(map[rune]int)

	for _, line := range items {
		characterCountByLine := make(map[rune]int)

		for _, c := range line {
			if characterCountByLine[c] == 0 {
				characterCount[c]++
				characterCountByLine[c]++
			}
		}
	}

	for k, v := range characterCount {
		if v == len(items) {
			sum += charValue(k)
		}
	}
	return sum
}

// Function to return value of character. a=1, b=2, c=3, etc.
// A=27, B=28, C=29, etc.
func charValue(c rune) int {
	if c >= 'a' && c <= 'z' {
		return int(c - 'a' + 1)
	}
	if c >= 'A' && c <= 'Z' {
		return int(c - 'A' + 27)
	}
	return 0
}

func GroupArray(input []string, size int) [][]string {
	var result [][]string
	for i := 0; i < len(input); i += size {
		end := i + size
		if end > len(input) {
			end = len(input)
		}
		result = append(result, input[i:end])
	}
	return result
}
