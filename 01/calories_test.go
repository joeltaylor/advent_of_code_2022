package main

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewCalories(t *testing.T) {
	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

	fs := fstest.MapFS{
		"input.txt": {Data: []byte(input)},
	}

	caloriesList, err := NewCaloriesFromFS(fs, "input.txt")

	if err != nil {
		t.Fatal(err)
	}

	assertLength(t, caloriesList, 5)
	assertCalorie(t, caloriesList[0], Calories{Values: []int{1000, 2000, 3000}})
	assertCalorie(t, caloriesList[1], Calories{Values: []int{4000}})
	assertCalorie(t, caloriesList[2], Calories{Values: []int{5000, 6000}})
}

func assertLength(t *testing.T, list []Calories, length int) {
	t.Helper()

	if len(list) != length {
		t.Errorf("got %d, wanted %d", len(list), length)
	}
}

func assertCalorie(t *testing.T, got Calories, want Calories) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
