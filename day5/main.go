package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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

func reverse(numbers []string) []string {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func part1(input string) string {
	splitInput := strings.Split(input, "\n\n")
	commands := splitInput[1]
	rawCrates := splitInput[0]

	rawCratesLines := strings.Split(rawCrates, "\n")
	indexesString := rawCratesLines[len(rawCratesLines)-1]
	indexes := strings.Split(indexesString, "   ")

	var trimmedIndex []int
	for i := range indexes {
		spacesRemoved := strings.ReplaceAll(indexes[i], " ", "")
		intValue, _ := strconv.Atoi(spacesRemoved)
		trimmedIndex = append(trimmedIndex, intValue)
	}

	rawCratesLines = rawCratesLines[:len(rawCratesLines)-1]
	stacks := make(map[int][]string)

	for t := range trimmedIndex {
		for a := range rawCratesLines {
			startPoint := 0
			endPoint := 3
			if rawCratesLines[a][startPoint:endPoint] != "   " {
				item := rawCratesLines[a][startPoint:endPoint]
				item = string(item[1])
				stacks[t+1] = append(stacks[t+1], item)
			}
			lineWithRemovedStack := rawCratesLines[a][3:]
			if len(lineWithRemovedStack) > 0 {
				nextValue := string(lineWithRemovedStack[0])
				if nextValue == " " {
					lineWithRemovedStack = string(lineWithRemovedStack[1:])
				}
			}
			rawCratesLines[a] = lineWithRemovedStack
		}
	}

	splitCommands := strings.Split(commands, "\n")
	splitCommands = splitCommands[:len(splitCommands)-1]

	for c := range splitCommands {
		fromArray := strings.Split(splitCommands[c], " ")
		from, _ := strconv.Atoi(fromArray[3])
		howmany, _ := strconv.Atoi(fromArray[1])
		to, _ := strconv.Atoi(fromArray[5])

		var items []string
		for r := 1; r <= howmany; r++ {
			itemToMove := stacks[from][0]
			// fmt.Println("moving ", itemToMove)
			items = append(items, itemToMove)
			stacks[from] = stacks[from][1:len(stacks[from])]
		}
		var newStack []string
		reversedItems := reverse(items)
		newStack = append(newStack, reversedItems...)
		newStack = append(newStack, stacks[to]...)

		stacks[to] = newStack
	}

	var arr []string
	for f := range trimmedIndex {
		if len(stacks[f]) > 0 {

			arr = append(arr, stacks[f][0])
		}
	}
	arr = append(arr, stacks[len(stacks)][0])
	joinedletters := strings.Join(arr, "")
	return joinedletters
}

func part2(input string) string {
	splitInput := strings.Split(input, "\n\n")
	commands := splitInput[1]
	rawCrates := splitInput[0]

	rawCratesLines := strings.Split(rawCrates, "\n")
	indexesString := rawCratesLines[len(rawCratesLines)-1]
	indexes := strings.Split(indexesString, "   ")

	var trimmedIndex []int
	for i := range indexes {
		spacesRemoved := strings.ReplaceAll(indexes[i], " ", "")
		intValue, _ := strconv.Atoi(spacesRemoved)
		trimmedIndex = append(trimmedIndex, intValue)
	}

	rawCratesLines = rawCratesLines[:len(rawCratesLines)-1]
	stacks := make(map[int][]string)

	for t := range trimmedIndex {
		for a := range rawCratesLines {
			startPoint := 0
			endPoint := 3
			if rawCratesLines[a][startPoint:endPoint] != "   " {
				item := rawCratesLines[a][startPoint:endPoint]
				item = string(item[1])
				stacks[t+1] = append(stacks[t+1], item)
			}
			lineWithRemovedStack := rawCratesLines[a][3:]
			if len(lineWithRemovedStack) > 0 {
				nextValue := string(lineWithRemovedStack[0])
				if nextValue == " " {
					lineWithRemovedStack = string(lineWithRemovedStack[1:])
				}
			}
			rawCratesLines[a] = lineWithRemovedStack
		}
	}

	splitCommands := strings.Split(commands, "\n")
	splitCommands = splitCommands[:len(splitCommands)-1]

	for c := range splitCommands {
		fromArray := strings.Split(splitCommands[c], " ")
		from, _ := strconv.Atoi(fromArray[3])
		howmany, _ := strconv.Atoi(fromArray[1])
		to, _ := strconv.Atoi(fromArray[5])

		var items []string
		for r := 1; r <= howmany; r++ {
			itemToMove := stacks[from][0]
			// fmt.Println("moving ", itemToMove)
			items = append(items, itemToMove)
			stacks[from] = stacks[from][1:len(stacks[from])]
		}
		var newStack []string
		newStack = append(newStack, items...)
		newStack = append(newStack, stacks[to]...)
		stacks[to] = newStack
	}

	var arr []string
	for f := range trimmedIndex {
		if len(stacks[f]) > 0 {
			arr = append(arr, stacks[f][0])
		}
	}
	arr = append(arr, stacks[len(stacks)][0])
	joinedletters := strings.Join(arr, "")
	return joinedletters
}
