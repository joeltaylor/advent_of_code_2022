package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Println(PartOne(string(data)))
	fmt.Println(PartTwo(string(data)))
}

var Crate = regexp.MustCompile("[(A-Z)]")
var NewLine = regexp.MustCompile("\n")

type Instruction struct {
	From   int
	To     int
	Amount int
}

func PartOne(input string) string {
	splitInput := strings.Split(input, "\n\n")
	crates := GenerateStackMap(splitInput[0])
	instructions := GenerateInstructions(splitInput[1])

	for _, v := range instructions {
		items := make([]string, v.Amount)
		items, crates[v.From-1] = crates[v.From-1][:v.Amount], crates[v.From-1][v.Amount:]

		for i := 0; i < v.Amount; i++ {
			var x string
			x, items = items[0], items[1:]
			crates[v.To-1] = append([]string{x}, crates[v.To-1]...)
		}
	}

	topStrings := ""

	for _, v := range crates {
		topStrings += v[0]
	}

	return topStrings

}

func PartTwo(input string) string {
	splitInput := strings.Split(input, "\n\n")
	crates := GenerateStackMap(splitInput[0])
	instructions := GenerateInstructions(splitInput[1])

	for _, v := range instructions {
		items := make([]string, v.Amount)
		items, crates[v.From-1] = crates[v.From-1][:v.Amount], crates[v.From-1][v.Amount:]

		for i := 0; i < v.Amount; i++ {
			var x string
			x, items = items[len(items)-1], items[:len(items)-1]
			crates[v.To-1] = append([]string{x}, crates[v.To-1]...)
		}
	}

	topStrings := ""

	for _, v := range crates {
		topStrings += v[0]
	}

	return topStrings

}

func GenerateStackMap(input string) [][]string {
	firstNewLine := NewLine.FindAllStringIndex(input, 1)[0][1]
	allIndicies := Crate.FindAllStringSubmatchIndex(input, -1)
	numberOfColumns := firstNewLine / 4

	crates := make([][]string, numberOfColumns)
	for _, v := range allIndicies {
		c := string(input[v[0]])
		p := v[1] / 4
		column := int(math.Mod(float64(p), float64(numberOfColumns)))

		crates[column] = append(crates[column], c)
	}

	return crates
}

func GenerateInstructions(input string) []Instruction {
	instructionPattern := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	instructions := instructionPattern.FindAllStringSubmatch(input, -1)

	parsed := []Instruction{}
	for _, v := range instructions {
		from, _ := strconv.Atoi(v[2])
		to, _ := strconv.Atoi(v[3])
		amount, _ := strconv.Atoi(v[1])
		parsed = append(parsed, Instruction{From: from, To: to, Amount: amount})
	}
	return parsed
}
