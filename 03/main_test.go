package main

import (
	"reflect"
	"testing"
)

func TestDetectDupe(t *testing.T) {
	input := "vJrwpWtwJgWrhcsFMMfFFhFp"
	got := DetectDupe(input)

	if got != 'p' {
		t.Errorf("Got %v, wanted 'p'", got)

	}
}

func TestCharValue(t *testing.T) {
	got := charValue('a')
	if got != 1 {
		t.Errorf("Got %d, wanted 1", got)
	}

	got = charValue('A')
	if got != 27 {
		t.Errorf("Got %d, wanted 27", got)
	}
}

func TestPart2(t *testing.T) {
	input := "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw"

	got := Part2(input)
	if got != 157 {
		t.Errorf("Got %d, wanted 157", got)
	}
}
func TestPart1(t *testing.T) {
	input := "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw"

	got := Part1(input)
	if got != 70 {
		t.Errorf("Got %d, wanted 70", got)
	}

}

func TestGroupArray(t *testing.T) {
	input := []string{"a", "b", "c", "d", "e", "f"}
	got := GroupArray(input, 3)

	if !reflect.DeepEqual(got, [][]string{{"a", "b", "c"}, {"d", "e", "f"}}) {
		t.Errorf("Got %v, wanted %v", got, [][]string{{"a", "b", "c"}, {"d", "e", "f"}})
	}
}
