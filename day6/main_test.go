package main

import (
	"testing"
)

type testInput struct {
	want  int
	input string
}

func TestPart1test1(t *testing.T) {
	test1 := testInput{
		want:  7,
		input: `mjqjpqmgbljsphdztnvjfqwrcgsmlb`,
	}

	ans := part1(test1.input)

	if ans != test1.want {
		t.Error("wanted: ", test1.want, "got: ", ans)
	}
}

func TestPart1test2(t *testing.T) {
	test1 := testInput{
		want:  5,
		input: `bvwbjplbgvbhsrlpgdmjqwftvncz`,
	}

	ans := part1(test1.input)

	if ans != test1.want {
		t.Error("wanted: ", test1.want, "got: ", ans)
	}
}

func TestPart1test3(t *testing.T) {
	test1 := testInput{
		want:  6,
		input: `nppdvjthqldpwncqszvftbrmjlhg`,
	}

	ans := part1(test1.input)

	if ans != test1.want {
		t.Error("wanted: ", test1.want, "got: ", ans)
	}
}

func TestPart1test3test4(t *testing.T) {
	test1 := testInput{
		want:  10,
		input: `nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`,
	}

	ans := part1(test1.input)

	if ans != test1.want {
		t.Error("wanted: ", test1.want, "got: ", ans)
	}
}

func TestPart1test3test5(t *testing.T) {
	test1 := testInput{
		want:  11,
		input: `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`,
	}

	ans := part1(test1.input)

	if ans != test1.want {
		t.Error("wanted: ", test1.want, "got: ", ans)
	}
}

func TestPart2test1(t *testing.T) {
	test2 := testInput{
		want:  19,
		input: `mjqjpqmgbljsphdztnvjfqwrcgsmlb`,
	}

	ans := part2(test2.input)

	if ans != test2.want {
		t.Error("wanted: ", test2.want, "got: ", ans)
	}
}

func TestPart2test2(t *testing.T) {
	test2 := testInput{
		want:  23,
		input: `bvwbjplbgvbhsrlpgdmjqwftvncz`,
	}

	ans := part2(test2.input)

	if ans != test2.want {
		t.Error("wanted: ", test2.want, "got: ", ans)
	}
}

func TestPart2test3(t *testing.T) {
	test2 := testInput{
		want:  23,
		input: `nppdvjthqldpwncqszvftbrmjlhg`,
	}

	ans := part2(test2.input)

	if ans != test2.want {
		t.Error("wanted: ", test2.want, "got: ", ans)
	}
}

func TestPart2test4(t *testing.T) {
	test2 := testInput{
		want:  29,
		input: `nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`,
	}

	ans := part2(test2.input)

	if ans != test2.want {
		t.Error("wanted: ", test2.want, "got: ", ans)
	}
}

func TestPart2test5(t *testing.T) {
	test2 := testInput{
		want:  26,
		input: `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`,
	}

	ans := part2(test2.input)

	if ans != test2.want {
		t.Error("wanted: ", test2.want, "got: ", ans)
	}
}
