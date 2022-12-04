package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
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

func convStringArrayToInt(stringArray []string) []int {
	var intArray []int
	for a := range stringArray {
		intValue, _ := strconv.Atoi(stringArray[a])
		intArray = append(intArray, intValue)
	}
	return intArray
}

func part1(input string) int {
	elfPairs := strings.Split(input, "\n")
	elfPairs = elfPairs[:len(elfPairs)-1]
	numberOfDuplicates := 0

	for i := range elfPairs {
		splitElves := strings.Split(elfPairs[i], ",")
		firstElf := splitElves[0]
		secondElf := splitElves[1]

		firstElfStringSections := strings.Split(firstElf, "-")
		secondElfStringSections := strings.Split(secondElf, "-")

		firstElfIntSections := convStringArrayToInt(firstElfStringSections)
		secondElfIntSections := convStringArrayToInt(secondElfStringSections)

		var firstElfFullSections []int
		var secondElfFullSections []int

		for b := firstElfIntSections[0]; b < firstElfIntSections[1]+1; b++ {
			firstElfFullSections = append(firstElfFullSections, b)
		}

		for b := secondElfIntSections[0]; b <= secondElfIntSections[1]; b++ {
			secondElfFullSections = append(secondElfFullSections, b)
		}

		var joined []int
		joined = append(firstElfFullSections, secondElfFullSections...)

		sort.Slice(joined, func(i, j int) bool {
			return joined[i] < joined[j]
		})

		var last int
		firstItem := true
		var duplicates []int

		for c := range joined {
			if firstItem {
				firstItem = false
			} else {
				if joined[c] == last {
					duplicates = append(duplicates, joined[c])
				}
			}
			last = joined[c]
		}

		dupeLength := len(duplicates)
		firstLength := len(firstElfFullSections)
		secondLength := len(secondElfFullSections)

		if dupeLength == firstLength || dupeLength == secondLength {
			numberOfDuplicates++
		}
	}
	return numberOfDuplicates
}

func part2(input string) int {
	elfPairs := strings.Split(input, "\n")
	elfPairs = elfPairs[:len(elfPairs)-1]
	dupesInPairs := 0

	for i := range elfPairs {
		splitElves := strings.Split(elfPairs[i], ",")
		firstElf := splitElves[0]
		secondElf := splitElves[1]

		firstElfStringSections := strings.Split(firstElf, "-")
		secondElfStringSections := strings.Split(secondElf, "-")

		firstElfIntSections := convStringArrayToInt(firstElfStringSections)
		secondElfIntSections := convStringArrayToInt(secondElfStringSections)

		var firstElfFullSections []int
		var secondElfFullSections []int

		for b := firstElfIntSections[0]; b < firstElfIntSections[1]+1; b++ {
			firstElfFullSections = append(firstElfFullSections, b)
		}

		for b := secondElfIntSections[0]; b <= secondElfIntSections[1]; b++ {
			secondElfFullSections = append(secondElfFullSections, b)
		}

		var joined []int
		joined = append(firstElfFullSections, secondElfFullSections...)

		sort.Slice(joined, func(i, j int) bool {
			return joined[i] < joined[j]
		})

		var last int
		firstItem := true
		var duplicates []int

		for c := range joined {
			if firstItem {
				firstItem = false
			} else {
				if joined[c] == last {
					duplicates = append(duplicates, joined[c])
				}
			}
			last = joined[c]
		}

		dupeLength := len(duplicates)

		if dupeLength > 0 {
			dupesInPairs++
		}
	}
	return dupesInPairs
}
