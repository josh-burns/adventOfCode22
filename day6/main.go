package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	input, _ := ioutil.ReadFile("input.txt")

	if part == 1 {
		ans := part1(string(input))
		fmt.Println(ans)
	}
	if part == 2 {
		ans := part2(string(input))
		fmt.Println(ans)
	}
}

func part1(input string) int {
	inputArray := strings.Split(input, "")

	var marker int
Out:
	for i := range inputArray {
		if i < len(inputArray)-3 {
			potentialMarker := input[i : i+4]
			notUnique := 0
			for a := range potentialMarker {
				for b := a + 1; b < len(potentialMarker); b++ {
					if string(potentialMarker[a]) == string(potentialMarker[b]) {
						notUnique++
					}
				}
			}
			if notUnique == 0 {
				marker = i + 4
				break Out
			}

		}
	}
	return marker
}

func part2(input string) int {
	inputArray := strings.Split(input, "")

	var marker int
Out:
	for i := range inputArray {
		if i < len(inputArray)-14 {
			potentialMarker := input[i : i+14]
			// fmt.Println(len(input))
			// fmt.Println(potentialMarker)
			notUnique := 0
			for a := range potentialMarker {
				for b := a + 1; b < len(potentialMarker); b++ {
					if string(potentialMarker[a]) == string(potentialMarker[b]) {
						notUnique++
					}
				}
			}
			if notUnique == 0 {
				marker = i + 14
				break Out
			}

		}
	}
	return marker
}
