package main

import (
	"fmt"
	"os"
)

func main() {
	calories, _ := NewCaloriesFromFS(os.DirFS("."), "input.txt")
	fmt.Println(Part1(calories))
	fmt.Println(Part2(calories))
}
