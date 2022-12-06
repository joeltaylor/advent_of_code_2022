package main

import (
	"reflect"
	"testing"
)

type test struct {
	name     string
	input    [][]int
	expected bool
}

func TestPartOne(t *testing.T) {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

	got := PartOne(input)

	if got != 2 {
		t.Errorf("got %v instead of 2", got)
	}

}

func TestPartTwo(t *testing.T) {
	input := `2-4,6-8
	2-3,4-5
	5-7,7-9
	2-8,3-7
	6-6,4-6
	2-6,4-8`

	got := PartTwo(input)

	if got != 4 {
		t.Errorf("got %v instead of 4", got)
	}

}

func TestRealMain(t *testing.T) {
	got := realMain()
	if !reflect.DeepEqual(got, []int{556, 876}) {
		t.Errorf("got %v, expected 556, 876", got)
	}
}
