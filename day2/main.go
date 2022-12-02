package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)



func main(){

	input, err := os.ReadFile("input.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	if part == 1 {
		ans := part1(string(input))
		fmt.Println(ans)
	}
	if part == 2 {
		ans := part2(string(input))
		fmt.Println(ans)
	}
}


func part1(input string ) int{
	rounds := strings.Split(input, "\n")
	pointsTally := 0

	pointsLegend := make(map[string]int)
	pointsLegend["X"] = 1
	pointsLegend["Y"] = 2
	pointsLegend["Z"] = 3

	weaknesses := make(map[string]string)
	weaknesses["X"] = "B"
	weaknesses["Y"] = "C"
	weaknesses["Z"] = "A"

	legend := make(map[string]string)
	legend["A"] = "X"
	legend["B"] = "Y"
	legend["C"] = "Z"

	for i := 0; i < len(rounds); i++{
		splitChoices := strings.Split(rounds[i], " ")
		myChoice := splitChoices[1]
		pointsTally += pointsLegend[myChoice]
		opponentChoice := splitChoices[0]
		switch {
		case legend[opponentChoice] == myChoice:
			pointsTally += 3
		case opponentChoice != weaknesses[myChoice]:
			pointsTally += 6
		default:
		}
	}
	return pointsTally
}

func part2(input string) int{
	pointsTally := 0
	legend := make(map[string]int)
	legend["AX"] = 3
	legend["AY"] = 4
	legend["AZ"] = 8
	legend["BX"] = 1
	legend["BY"] = 5
	legend["BZ"] = 9
	legend["CX"] = 2
	legend["CY"] = 6
	legend["CZ"] = 7

	rounds := strings.Split(input, "\n")


	for i := 0; i < len(rounds); i++{
		splitChoices := strings.Split(rounds[i], " ")
		joined := strings.Join(splitChoices, "")

		pointsTally += legend[joined]
	}
	return pointsTally
}
