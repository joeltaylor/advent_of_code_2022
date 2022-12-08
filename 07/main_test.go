package main

import "testing"

func TestParseCommands(t *testing.T) {
	input := `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
dir d
$ cd d
$ ls
1 i
$ cd ..
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

	got := ParseCommands(input)

	if got["/"] != 48381166 {
		t.Errorf("got %v, expected 48381166", got)
	}
	if got["/-a-e"] != 585 {
		t.Errorf("got %v, expected 585", got)
	}
	if got["/-a"] != 94854 {
		t.Errorf("got %v, expected 94853", got)
	}
	if got["/-d"] != 24933642 {
		t.Errorf("got %v, expected 24933642", got)
	}
	if got["/-a-e-d"] != 1 {
		t.Errorf("got %v, expected 1", got)
	}
}

func TestPartOne(t *testing.T) {
	input := `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

	got := PartOne(input)

	if got != 95437 {
		t.Errorf("got %v, expected 95437 ", got)
	}
}

func TestPartTwo(t *testing.T) {
	input := `$ cd /
	$ ls
	dir a
	14848514 b.txt
	8504156 c.dat
	dir d
	$ cd a
	$ ls
	dir e
	29116 f
	2557 g
	62596 h.lst
	$ cd e
	$ ls
	584 i
	$ cd ..
	$ cd ..
	$ cd d
	$ ls
	4060174 j
	8033020 d.log
	5626152 d.ext
	7214296 k`

	got := PartTwo(input)

	if got != 24933642 {
		t.Errorf("got %v, expected 24933642", got)
	}
}
