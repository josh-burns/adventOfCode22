package main

import (
	"testing"
)

type testInput struct{
	want int
	input string
}

func TestPart1(t *testing.T){
	test1 := testInput{
		want: 95437,
		input: `$ cd /
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
7214296 k`,
	}

	ans := part1(test1.input)

	if ans != test1.want {
		t.Error("wanted: ", test1.want, "got: ", ans)
	}
}

func TestPart2(t *testing.T){
	test2 := testInput{
		want: 112,
		input: `
                      112
                       `,
	}

	ans := part2(test2.input)

	if ans != test2.want {
		t.Error("wanted: ", test2.want, "got: ", ans)
	}
}
