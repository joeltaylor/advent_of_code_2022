package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestHasContact(t *testing.T) {
	type test struct {
		Name   string
		Input  []Coordinates
		Output bool
	}

	tests := []test{
		{Name: "left contact", Input: []Coordinates{{X: 2, Y: 2}, {X: 1, Y: 2}}, Output: true},
		{Name: "right contact", Input: []Coordinates{{X: 2, Y: 2}, {X: 3, Y: 2}}, Output: true},
		{Name: "top contact", Input: []Coordinates{{X: 2, Y: 2}, {X: 2, Y: 3}}, Output: true},
		{Name: "bottom contact", Input: []Coordinates{{X: 2, Y: 2}, {X: 2, Y: 1}}, Output: true},
		{Name: "bottom left diagonal contact", Input: []Coordinates{{X: 2, Y: 2}, {X: 1, Y: 1}}, Output: true},
		{Name: "bottom right diagonal contact", Input: []Coordinates{{X: 2, Y: 2}, {X: 3, Y: 1}}, Output: true},
		{Name: "top left diagonal contact", Input: []Coordinates{{X: 2, Y: 2}, {X: 1, Y: 3}}, Output: true},
		{Name: "top right diagonal contact", Input: []Coordinates{{X: 2, Y: 2}, {X: 3, Y: 3}}, Output: true},
		{Name: "overlapping", Input: []Coordinates{{X: 2, Y: 2}, {X: 2, Y: 2}}, Output: true},
		{Name: "no contact", Input: []Coordinates{{X: 2, Y: 2}, {X: 0, Y: 2}}, Output: false},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			head := Knot{Coordinates: tt.Input[0]}
			tail := Knot{Coordinates: tt.Input[1], Head: &head}

			got := tail.HasContact()

			if got != tt.Output {
				t.Errorf("got %v, exptected %v", tt.Output, got)

			}
		})
	}

}

// 5 . . . . . .
// 4 . . . . . .
// 3 . . . . . .
// 2 . T . . . .
// 1 . . . H . .
// 0 . . . . . .
//	 0 1 2 3 4 5

func TestMoveToContact(t *testing.T) {
	type test struct {
		Name   string
		Head   Coordinates
		Tail   Coordinates
		Output Coordinates
	}
	tests := []test{
		{Name: "Follow for side contact", Head: Coordinates{X: 4, Y: 2}, Tail: Coordinates{X: 2, Y: 2}, Output: Coordinates{X: 3, Y: 2}},
		{Name: "Follow for bottom contact", Head: Coordinates{X: 2, Y: 4}, Tail: Coordinates{X: 2, Y: 2}, Output: Coordinates{X: 2, Y: 3}},
		{Name: "Follow from diagonal contact", Head: Coordinates{X: 3, Y: 3}, Tail: Coordinates{X: 1, Y: 2}, Output: Coordinates{X: 2, Y: 3}},
		{Name: "Follow from bottom diagonal contact", Head: Coordinates{X: 4, Y: 2}, Tail: Coordinates{X: 3, Y: 0}, Output: Coordinates{X: 4, Y: 1}},
		{Name: "Follow into the negative void", Head: Coordinates{X: -2, Y: 0}, Tail: Coordinates{X: 0, Y: 0}, Output: Coordinates{X: -1, Y: 0}},
		{Name: "Another", Head: Coordinates{X: 3, Y: 1}, Tail: Coordinates{X: 1, Y: 2}, Output: Coordinates{X: 2, Y: 1}},
		{Name: "Anotherrrrr", Head: Coordinates{X: 1, Y: 1}, Tail: Coordinates{X: 3, Y: 2}, Output: Coordinates{X: 2, Y: 1}},
		{Name: "Anotherrrrrrrrr", Head: Coordinates{X: 1, Y: 3}, Tail: Coordinates{X: 3, Y: 2}, Output: Coordinates{X: 2, Y: 3}},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			head := Knot{Coordinates: tt.Head}
			subject := Knot{Coordinates: tt.Tail, Head: &head}
			subject.MoveToContact()

			if subject.X != tt.Output.X || subject.Y != tt.Output.Y {
				t.Errorf("got %v, exptected %v", subject, tt.Output)

			}
		})
	}

}

func TestMove(t *testing.T) {
	type test struct {
		Name   string
		Input  string
		Output Coordinates
	}
	tests := []test{
		{Name: "Move right", Input: "R", Output: Coordinates{X: 3, Y: 2}},
		{Name: "Move left", Input: "L", Output: Coordinates{X: 1, Y: 2}},
		{Name: "Move up", Input: "U", Output: Coordinates{X: 2, Y: 3}},
		{Name: "Move down", Input: "D", Output: Coordinates{X: 2, Y: 1}},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			subject := Knot{Coordinates: Coordinates{X: 2, Y: 2}}
			subject.Move(tt.Input)

			if subject.X != tt.Output.X || subject.Y != tt.Output.Y {
				t.Errorf("got %v, exptected %v", tt.Output, subject)
			}
		})
	}

}

func TestPartOne(t *testing.T) {
	input := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

	buf := strings.NewReader(input)
	scanner := bufio.NewScanner(buf)

	got := Simulate(scanner, 1)

	if got != 13 {
		t.Errorf("got %v, expected 13", got)
	}

}

func TestPartTwo(t *testing.T) {
	input := `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`
	buf := strings.NewReader(input)
	scanner := bufio.NewScanner(buf)

	got := Simulate(scanner, 9)

	if got != 36 {
		t.Errorf("got %v, expected 36", got)
	}

}
