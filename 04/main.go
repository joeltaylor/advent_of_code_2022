package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	realMain()
}

func realMain() []int {
	input, _ := os.ReadFile("input.txt")
	return []int{PartOne(string(input)), PartTwo(string(input))}
}

func PartOne(input string) int {
	overlapCount := 0
	allAssignmentPairs := strings.Split(input, "\n")

	for _, assignmentPair := range allAssignmentPairs {
		if len(assignmentPair) == 0 {
			continue
		}
		min1, max1, min2, max2 := parseAssignment(assignmentPair)
		if (min1 <= min2 && max1 >= max2) || (min2 <= min1 && max2 >= max1) {
			overlapCount++
		}
	}

	return overlapCount
}

func PartTwo(input string) int {
	overlapCount := 0
	allAssignmentPairs := strings.Split(input, "\n")

	for _, assignmentPair := range allAssignmentPairs {
		if len(assignmentPair) == 0 {
			continue
		}
		min1, max1, min2, max2 := parseAssignment(assignmentPair)
		if (min1 < min2 && max1 < min2) || (min2 < min1 && max2 < min1) {
			continue
		}
		overlapCount++
	}

	return overlapCount
}

func parseAssignment(input string) (int, int, int, int) {
	pattern := regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)
	matches := pattern.FindStringSubmatch(input)

	min1, _ := strconv.Atoi(matches[1])
	max1, _ := strconv.Atoi(matches[2])
	min2, _ := strconv.Atoi(matches[3])
	max2, _ := strconv.Atoi(matches[4])

	return min1, max1, min2, max2
}
