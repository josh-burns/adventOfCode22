package main

import (
	"testing"
)

type testInput struct {
	want  string
	input string
}

func TestPart1(t *testing.T) {
	test1 := testInput{
		want: "CMZ",
		input: `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`,
	}

	ans := part1(test1.input)

	if ans != test1.want {
		t.Error("wanted: ", test1.want, "got: ", ans)
	}
}

// func TestPart2(t *testing.T){
// 	test2 := testInput{
// 		want: 112,
// 		input: `
//                       112
//                        `,
// 	}

// 	ans := part2(test2.input)

// 	if ans != test2.want {
// 		t.Error("wanted: ", test2.want, "got: ", ans)
// 	}
// }
