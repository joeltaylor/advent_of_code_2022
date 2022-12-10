package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Coordinates struct {
	X int
	Y int
}

type Knot struct {
	Coordinates
	Visited map[string]bool
	Head    *Knot
	Tail    *Knot
	Key     int
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	fmt.Println(Simulate(scanner, 1))
	fmt.Println(Simulate(scanner, 9))
}

func Simulate(s *bufio.Scanner, knots int) int {
	head := Knot{Key: 0}

	currHead := &head
	for i := 0; i < knots; i++ {
		tail := Knot{Key: i, Head: currHead, Visited: map[string]bool{"0|0": true}}
		currHead.Tail = &tail
		currHead = &tail
	}

	for s.Scan() {
		line := strings.Split(s.Text(), " ")
		direction := line[0]
		distance, _ := strconv.Atoi(line[1])

		for i := distance; i > 0; i-- {
			head.Move(direction)
			head.RelayMove()
		}
	}

	n := head

	for !n.last() {
		n = *n.Tail
	}

	return len(n.Visited)
}

func (c *Knot) RelayMove() {
	curr := c.Tail
	for curr != nil {
		if !curr.HasContact() {
			curr.Visited[fmt.Sprintf("%d|%d", curr.X, curr.Y)] = true
			curr.MoveToContact()
			curr.Visited[fmt.Sprintf("%d|%d", curr.X, curr.Y)] = true
		}
		curr = curr.Tail
	}
}

func (c Knot) HasContact() bool {
	tailCoordinates := c.Coordinates
	switch h := c.Head; {
	case reflect.DeepEqual(h.left(), tailCoordinates):
		return true
	case reflect.DeepEqual(h.right(), tailCoordinates):
		return true
	case reflect.DeepEqual(h.top(), tailCoordinates):
		return true
	case reflect.DeepEqual(h.bottom(), tailCoordinates):
		return true
	case reflect.DeepEqual(h.top_left(), tailCoordinates):
		return true
	case reflect.DeepEqual(h.bottom_left(), tailCoordinates):
		return true
	case reflect.DeepEqual(h.top_right(), tailCoordinates):
		return true
	case reflect.DeepEqual(h.bottom_right(), tailCoordinates):
		return true
	case reflect.DeepEqual(h.Coordinates, tailCoordinates):
		return true
	default:
		return false

	}
}

func (c *Knot) MoveToContact() {
	head := c.Head

	if c.Y < head.Y {
		c.Y++
	}
	if c.Y > head.Y {
		c.Y--
	}
	if c.X < head.X {
		c.X++
	}
	if c.X > head.X {
		c.X--
	}
}

func (c *Knot) Move(direction string) {
	switch direction {
	case "R":
		c.X++
	case "U":
		c.Y++
	case "L":
		c.X--
	case "D":
		c.Y--
	}
}
func (c Knot) last() bool {
	return c.Tail == nil
}

func (c Knot) left() Coordinates {
	return Coordinates{X: c.X - 1, Y: c.Y}
}
func (c Knot) right() Coordinates {
	return Coordinates{X: c.X + 1, Y: c.Y}
}
func (c Knot) top() Coordinates {
	return Coordinates{X: c.X, Y: c.Y + 1}
}
func (c Knot) bottom() Coordinates {
	return Coordinates{X: c.X, Y: c.Y - 1}
}
func (c Knot) top_left() Coordinates {
	return Coordinates{X: c.X - 1, Y: c.Y + 1}
}
func (c Knot) top_right() Coordinates {
	return Coordinates{X: c.X + 1, Y: c.Y + 1}
}
func (c Knot) bottom_left() Coordinates {
	return Coordinates{X: c.X - 1, Y: c.Y - 1}
}
func (c Knot) bottom_right() Coordinates {
	return Coordinates{X: c.X + 1, Y: c.Y - 1}
}
