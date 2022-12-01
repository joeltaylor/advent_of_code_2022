package main

import "testing"

func TestPart1(t *testing.T) {
	list := []Calories{
		{[]int{1000, 2000}},
		{[]int{1000}},
	}
	got := Part1(list)

	if got != 3000 {
		t.Errorf("got %d, want %d", got, 3000)
	}

}

func TestPart2(t *testing.T) {
	list := []Calories{
		{[]int{2000}},
		{[]int{1000}},
		{[]int{2000, 2000}},
		{[]int{1000}},
	}
	got := Part2(list)

	if got != 7000 {
		t.Errorf("got %d, want %d", got, 7000)
	}

}
