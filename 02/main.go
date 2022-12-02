package main

import (
	"fmt"
	"os"
	"strings"
)

type Move struct {
	Letter    string
	AltLetter string
	Point     int
}

var (
	Rock     = Move{Letter: "A", AltLetter: "X", Point: 1}
	Paper    = Move{Letter: "B", AltLetter: "Y", Point: 2}
	Scissors = Move{Letter: "C", AltLetter: "Z", Point: 3}
	Lose     = Move{Letter: "X", Point: 0}
	Tie      = Move{Letter: "Y", Point: 3}
	Win      = Move{Letter: "Z", Point: 6}
)

func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(TotalScore(string(input)))
	fmt.Println(TotalScoreV2(string(input)))
}

func TotalScore(data string) int {
	currentScore := 0

	lines := strings.Split(data, "\n")

	// A, X rock
	// B, Y paper
	// C, Z scissors
	for _, v := range lines {
		characters := strings.Split(v, " ")

		switch characters[0] {
		case Rock.Letter:
			switch characters[1] {
			case Rock.AltLetter:
				currentScore += Rock.Point + Tie.Point
			case Paper.AltLetter:
				currentScore += Paper.Point + Win.Point
			case Scissors.AltLetter:
				currentScore += Scissors.Point + Lose.Point
			}
		case Paper.Letter:
			switch characters[1] {
			case Rock.AltLetter:
				currentScore += Rock.Point + Lose.Point
			case Paper.AltLetter:
				currentScore += Paper.Point + Tie.Point
			case Scissors.AltLetter:
				currentScore += Scissors.Point + Win.Point
			}
		case Scissors.Letter:
			switch characters[1] {
			case Rock.AltLetter:
				currentScore += Rock.Point + Win.Point
			case Paper.AltLetter:
				currentScore += Paper.Point + Lose.Point
			case Scissors.AltLetter:
				currentScore += Scissors.Point + Tie.Point
			}
		}
	}
	return currentScore

}

func TotalScoreV2(data string) int {
	currentScore := 0

	lines := strings.Split(data, "\n")

	// A, X rock
	// B, Y paper
	// C, Z scissors
	// X Lose
	// Y Draw
	// Z Win
	for _, v := range lines {
		characters := strings.Split(v, " ")

		switch characters[0] {
		case Rock.Letter:
			switch characters[1] {
			case Win.Letter:
				currentScore += Paper.Point + Win.Point
			case Lose.Letter:
				currentScore += Scissors.Point
			case Tie.Letter:
				currentScore += Rock.Point + Tie.Point
			}
		case Paper.Letter:
			switch characters[1] {
			case Win.Letter:
				currentScore += Scissors.Point + Win.Point
			case Lose.Letter:
				currentScore += Rock.Point
			case Tie.Letter:
				currentScore += Paper.Point + Tie.Point
			}
		case Scissors.Letter:
			switch characters[1] {
			case Win.Letter:
				currentScore += Rock.Point + Win.Point
			case Lose.Letter:
				currentScore += Paper.Point
			case Tie.Letter:
				currentScore += Scissors.Point + Tie.Point
			}
		}
	}
	return currentScore

}
