package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)



func main(){

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


func part1(input string ) int{
	splitLines := strings.Split(input, "\n" )

	var dupes []string
	for i := range splitLines{
		bag := splitLines[i]
		splitPoint := len(bag)/2
		length := len(bag)
		compartment1 := bag[0:splitPoint]
		compartment2 := bag[splitPoint:length]

		Out:
		for a := range compartment1 {
			letterToCheck := string(compartment1[a])
			for b := range compartment2{
				letterToBeChecked := string(compartment2[b])
				if letterToBeChecked == letterToCheck {
					dupes = append(dupes, letterToBeChecked)
					break Out
				}
			}
		}
	}


	priorities := 0
	for c := range dupes {
		runeArray := []rune(dupes[c])
		asciiValue := runeArray[0]
		if asciiValue > 96 {
			asciiValue -= 96
			priorities += int(asciiValue)
		}else{
			asciiValue -= 38
			priorities += int(asciiValue)
		}
	}
	return priorities
}

func part2(input string) int{
	return 112
}
