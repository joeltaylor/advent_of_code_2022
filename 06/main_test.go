package main

import "testing"

func TestDetectMarker(t *testing.T) {
	type test struct {
		Name   string
		Input  string
		Length int
		Output int
	}

	markerTests := []test{
		{Name: "first", Input: "bvwbjplbgvbhsrlpgdmjqwftvncz", Length: 4, Output: 5},
		{Name: "second", Input: "nppdvjthqldpwncqszvftbrmjlhg", Length: 4, Output: 6},
		{Name: "third", Input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", Length: 4, Output: 10},
		{Name: "fourth", Input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", Length: 4, Output: 11},
		{Name: "first 14", Input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", Length: 14, Output: 19},
		{Name: "second 14", Input: "bvwbjplbgvbhsrlpgdmjqwftvncz", Length: 14, Output: 23},
		{Name: "third 14", Input: "nppdvjthqldpwncqszvftbrmjlhg", Length: 14, Output: 23},
		{Name: "fourth 14", Input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", Length: 14, Output: 29},
		{Name: "fifth 14", Input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", Length: 14, Output: 26},
	}

	for _, tt := range markerTests {
		t.Run(tt.Name, func(t *testing.T) {
			got := DetectMarker(tt.Input, tt.Length)

			if got != tt.Output {
				t.Errorf("got %v, expected %v", got, tt.Output)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	type test struct {
		Name   string
		Input  string
		Output bool
	}
	uniqueTests := []test{
		{Name: "is unique", Input: "abcd", Output: true},
		{Name: "not unique", Input: "abbd", Output: false},
	}

	for _, tt := range uniqueTests {
		t.Run(tt.Name, func(t *testing.T) {
			got := checkUnique(tt.Input)

			if got != tt.Output {
				t.Errorf("got %v, expected %v", got, tt.Output)
			}
		})
	}

}
