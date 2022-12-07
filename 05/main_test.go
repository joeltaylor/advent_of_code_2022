package main

import (
	"reflect"
	"testing"
)

//	[D]
//
// [N] [C]
// [Z] [M] [P]
//
//	1   2   3
//
// move 1 from 2 to 1
// move 3 from 1 to 3
// move 2 from 2 to 1
// move 1 from 1 to 2
func TestGenerateStackMap(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

	got := GenerateStackMap(input)

	if !reflect.DeepEqual(got, [][]string{{"N", "Z"}, {"D", "C", "M"}, {"P"}}) {
		t.Errorf("got %v", got)
	}

	input = `    [D]        
[N] [C]        
[Z] [M] [P] [Q]
1   2   3   4

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

	got = GenerateStackMap(input)

	if !reflect.DeepEqual(got, [][]string{{"N", "Z"}, {"D", "C", "M"}, {"P"}, {"Q"}}) {
		t.Errorf("test got %v", got)
	}
}

func TestGenerateInstructions(t *testing.T) {
	input := `move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

	got := GenerateInstructions(input)

	if reflect.DeepEqual(got, []Instruction{{From: 2, To: 1, Amount: 1}, {From: 1, To: 3, Amount: 3}, {From: 2, To: 1, Amount: 2}, {From: 1, To: 1, Amount: 1}}) {
		t.Errorf("got %v", got)
	}

}

func TestPartOne(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

	got := PartOne(input)

	if got != "CMZ" {
		t.Errorf("got %s, expected %s", got, "CMZ")
	}
}
func TestPartTwo(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

	got := PartTwo(input)

	if got != "MCD" {
		t.Errorf("got %s, expected %s", got, "MCD")
	}
}
