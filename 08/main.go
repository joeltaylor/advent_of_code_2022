package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	input := strings.TrimSuffix(string(data), "\n")
	fmt.Println(CountVisibleTrees(input))
	fmt.Println(LargestScenicScore(input))
}

func CountVisibleTrees(input string) int {
	forestGrid := buildForest(input)
	totalVisible := 0

	for row, rowValue := range forestGrid {
		for col, colValue := range rowValue {
			if (row == 0) || (row == len(forestGrid)-1) || col == 0 || col == len(rowValue)-1 {
				totalVisible++
				continue
			}

			maxLeft := true
			for _, v := range rowValue[:col] {
				if v >= colValue {
					maxLeft = false
				}
			}

			maxRight := true
			for _, v := range rowValue[col+1:] {
				if v >= colValue {
					maxRight = false
				}
			}

			maxTop := true
			for _, tempRow := range forestGrid[:row] {
				if tempRow[col] >= colValue {
					maxTop = false
				}
			}

			maxBottom := true
			for _, tempRow := range forestGrid[row+1:] {
				if tempRow[col] >= colValue {
					maxBottom = false
				}
			}

			if maxLeft || maxRight || maxTop || maxBottom {
				totalVisible++
			}
		}
	}

	return totalVisible
}

func LargestScenicScore(input string) int {
	forestGrid := buildForest(input)

	maxScore := 0

	for row, rowValue := range forestGrid {
		for col, colValue := range rowValue {
			topCount, bottomCount, leftCount, rightCount := 0, 0, 0, 0

			for i := len(rowValue[:col]) - 1; i >= 0; i-- {
				v := rowValue[i]
				if v >= colValue {
					leftCount++
					break
				}
				leftCount++
			}

			for _, v := range rowValue[col+1:] {
				if v >= colValue {
					rightCount++
					break
				}
				rightCount++
			}

			for i := row - 1; i >= 0; i-- {
				value := forestGrid[i][col]
				if value >= colValue {
					topCount++
					break
				}
				topCount++
			}

			for i := row + 1; i < len(forestGrid); i++ {
				value := forestGrid[i][col]
				if value >= colValue {
					bottomCount++
					break
				}
				bottomCount++
			}

			score := topCount * bottomCount * leftCount * rightCount
			if score > maxScore {
				maxScore = score
			}
		}
	}

	return maxScore
}

func buildForest(input string) [][]int {
	lines := strings.Split(input, "\n")
	forestGrid := make([][]int, len(lines))

	for i, line := range lines {
		for _, height := range line {
			heightInt, _ := strconv.Atoi(string(height))
			forestGrid[i] = append(forestGrid[i], heightInt)
		}
	}
	return forestGrid
}
