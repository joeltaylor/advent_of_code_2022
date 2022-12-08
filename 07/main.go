package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	fmt.Println(PartOne(string(data)))
	fmt.Println(PartTwo(string(data)))
}

func PartOne(input string) int {
	fileSizes := ParseCommands(input)
	sum := 0

	for _, v := range fileSizes {
		if v <= 100000 {
			sum += v
		}
	}
	return sum
}

func PartTwo(input string) int {
	fileSizes := ParseCommands(input)
	unusedSpace := 70000000 - fileSizes["/"]
	spaceRequired := 30000000 - unusedSpace
	smallestDirectorySize := math.MaxInt

	for _, v := range fileSizes {
		if v >= spaceRequired {
			if v < smallestDirectorySize {
				smallestDirectorySize = v
			}
		}
	}
	return smallestDirectorySize
}

var cdPattern = regexp.MustCompile(`\$ cd (.*)`)
var filePattern = regexp.MustCompile(`(\d+) .*`)

func ParseCommands(input string) map[string]int {
	lines := strings.Split(input, "\n")

	breadCrumb := []string{}
	directoryTree := map[string]int{}

	for _, v := range lines {
		switch i := v; {
		case cdPattern.MatchString(i):
			match := cdPattern.FindStringSubmatch(i)
			// Update the breadCrumb whenever the directory is changed
			if match[1] == ".." {
				breadCrumb = breadCrumb[:len(breadCrumb)-1]
			} else {
				breadCrumb = append(breadCrumb, match[1])
			}
		case filePattern.MatchString(i):
			match := filePattern.FindStringSubmatch(i)
			fileSize, _ := strconv.Atoi(match[1])
			// Whenever a file is visited, add the file size to every directory listed
			// in the breadcrumb. A prefix is required since directory names are not guaranteed
			// to be unique throughout the tree.
			for i := range breadCrumb {
				prefix := strings.Join(breadCrumb[0:i+1], "-")
				directoryTree[prefix] += fileSize
			}
		}
	}
	return directoryTree

}
