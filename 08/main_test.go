package main

import (
	"testing"
)

func TestCountVisibleTrees(t *testing.T) {
	type visibleTest struct {
		Name   string
		Input  string
		Output int
	}

	tests := []visibleTest{
		{Name: "Single row", Input: "30373", Output: 5},
		{Name: "Two rows", Input: "30373\n25512", Output: 10},
		{Name: "Three rows", Input: "30373\n25512\n65332", Output: 14},
		{Name: "Full input", Input: "30373\n25512\n65332\n33549\n35390", Output: 21},
		{Name: "Visible from bottom", Input: "222\n222\n212", Output: 9},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got := CountVisibleTrees(tt.Input)

			if got != tt.Output {
				t.Errorf("got %v, expected %v", got, tt.Output)
			}
		})
	}
}

func TestLargestScenicScore(t *testing.T) {
	input := "30373\n25512\n65332\n33549\n35390"
	got := LargestScenicScore(input)

	if got != 8 {
		t.Errorf("got %v, expected 8", got)
	}
}

// func TestM(t *testing.T) {
// 	data, _ := os.ReadFile("input.txt")
// 	input := strings.TrimSuffix(string(data), "\n")
// 	LargestScenicScore(input)
//
// }
