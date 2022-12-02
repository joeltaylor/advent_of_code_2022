package main

import "testing"

func TestExamplePart1(t *testing.T) {
	input := `A Y
B X
C Z`

	got := TotalScore(input)

	if got != 15 {
		t.Errorf("got %d, want 15", got)
	}
}

func TestExamplePart2(t *testing.T) {
	input := `A Y
B X
C Z`

	got := TotalScoreV2(input)

	if got != 12 {
		t.Errorf("got %d, want 12", got)
	}
}
