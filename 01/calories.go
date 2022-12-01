package main

import (
	"io"
	"io/fs"
	"strconv"
	"strings"
)

type Calories struct {
	Values []int
}

func NewCaloriesFromFS(fileSystem fs.FS, filename string) ([]Calories, error) {
	caloriesFile, err := fileSystem.Open(filename)
	if err != nil {
		return nil, err
	}
	defer caloriesFile.Close()

	data, err := io.ReadAll(caloriesFile)
	lines := strings.Split(string(data), "\n\n")
	nums := make([]Calories, 0, len(lines))

	for _, v := range lines {
		o := strings.Split(v, "\n")
		if len(o) == 0 {
			continue
		}
		temp := make([]int, 0, len(o))
		for _, v := range o {
			n, _ := strconv.Atoi(v)
			temp = append(temp, n)
		}
		nums = append(nums, Calories{Values: temp})
	}

	return nums, nil
}
