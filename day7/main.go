package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
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

func part1(input string) int {
	splitArray := strings.Split(input, "\n")

	var currentDir []string
	for i := range splitArray {
		fmt.Println(splitArray[i])
		step := splitArray[i]
		if string(step[0]) == "$" {
			fmt.Println("step is a command")
			command := step[2:4]
			if command == "cd" {
				destination := step[5:]
				if destination == ".." {
					currentDir = currentDir[0 : len(currentDir)-1]
				}
				currentDir = append(currentDir, destination)

			}
		}
		if string(step[0:2]) == "dir" {
			fmt.Println("step is a dir")
		}
		if unicode.IsNumber(rune(step[0])) {
			fmt.Println("step is a file")
		}
		fmt.Println("pwd is ", currentDir)
	}
	fmt.Println(currentDir)
	return 112
}


func part2(input string) int{
	return 112
}
