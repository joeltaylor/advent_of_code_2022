package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Println(DetectMarker(string(data), 4))
	fmt.Println(DetectMarker(string(data), 14))

}

func DetectMarker(input string, length int) int {
	for i, j := 0, length; j <= len(input); i, j = i+1, j+1 {
		if checkUnique(input[i:j]) {
			fmt.Println(input[i:j])
			return j
		}
	}
	return -1
}

func checkUnique(input string) bool {
	seen := map[rune]bool{}
	for _, v := range input {
		if seen[v] {
			return false
		}
		seen[v] = true
	}

	return true

}
