package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Value    int
	Progress int
	Noop     bool
}

func main() {
	data, _ := os.ReadFile("input.txt")
	input := strings.TrimSuffix(string(data), "\n")
	fmt.Println(PartOne(input))
	fmt.Println(PartTwo(input))
}

func PartOne(input string) int {
	lines := strings.Split(input, "\n")
	instructions := []Instruction{}

	for _, v := range lines {
		var instruction Instruction
		if v == "noop" {
			instruction = Instruction{Value: 1, Noop: true}
		} else {
			stringValue := strings.Split(v, " ")[1]
			amount, _ := strconv.Atoi(stringValue)

			instruction = Instruction{Value: amount, Noop: false}
		}

		instructions = append(instructions, instruction)
	}

	checkPoints := map[int]int{20: 0, 60: 0, 100: 0, 1400: 0, 180: 0, 220: 0}
	clockCycle := 1
	xReg := 1

	for _, v := range instructions {
		clockCycle++
		v.Progress++

		if _, ok := checkPoints[clockCycle]; ok {
			checkPoints[clockCycle] = clockCycle * xReg
		}

		for !v.IsDone() {
			clockCycle++
			v.Progress++
			// This isn't needed, but I'm not sure why.
			if _, ok := checkPoints[clockCycle]; ok {
				checkPoints[clockCycle] = clockCycle * xReg
			}
		}
		if !v.Noop {
			xReg += v.Value
		}

		if _, ok := checkPoints[clockCycle]; ok {
			checkPoints[clockCycle] = clockCycle * xReg
		}
	}

	fmt.Println(checkPoints)

	totalxReg := 0

	for _, v := range checkPoints {
		totalxReg += v
	}
	return totalxReg
}

func PartTwo(input string) []string {
	lines := strings.Split(input, "\n")
	instructions := []Instruction{}

	for _, v := range lines {
		var instruction Instruction
		if v == "noop" {
			instruction = Instruction{Value: 1, Noop: true}
		} else {
			stringValue := strings.Split(v, " ")[1]
			amount, _ := strconv.Atoi(stringValue)

			instruction = Instruction{Value: amount, Noop: false}
		}

		instructions = append(instructions, instruction)
	}

	crt := [][]string{
		{"⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️"},
		{"⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️"},
		{"⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️"},
		{"⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️"},
		{"⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️"},
		{"⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️", "⬛️"},
	}

	clockCycle := 1
	xReg := 1
	for _, v := range instructions {
		for !v.IsDone() {
			fmt.Printf("Clock: %v Index: %v CompareClock: %v XReg: %v, Instruction: %v\n", clockCycle, realClock(clockCycle-1), realClock(clockCycle), xReg, v)
			if realClock(clockCycle)-1 == xReg || realClock(clockCycle)-1 == xReg-1 || realClock(clockCycle)-1 == xReg+1 {
				row := clockCycle / 40
				col := realClock(clockCycle)
				crt[row][col] = "🟩"
			}
			clockCycle++
			v.Progress++
		}
		if !v.Noop {
			xReg += v.Value
		}
	}

	result := []string{}
	for _, v := range crt {
		result = append(result, strings.Join(v, ""))
	}

	for _, v := range result {
		fmt.Println(v)
	}

	return result
}

func realClock(c int) int {
	v := int(math.Mod(float64(c), float64(40)))
	return v
}

func (i Instruction) IsDone() bool {
	if i.Noop {
		return i.Progress == 1
	}
	return i.Progress == 2
}
