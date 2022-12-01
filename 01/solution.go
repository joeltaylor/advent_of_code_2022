package main

import (
	"sort"
)

func Part1(list []Calories) int {
	var maxValue int

	for _, v := range list {
		res := 0
		for _, v := range v.Values {
			res += v
		}
		if res > maxValue {
			maxValue = res
		}
	}
	return maxValue
}

func Part2(list []Calories) int {
	var values []int

	for _, v := range list {
		res := 0
		for _, v := range v.Values {
			res += v
		}
		values = append(values, res)
	}

	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})

	sum := 0

	for _, v := range values[:3] {
		sum += v
	}

	return sum

}
