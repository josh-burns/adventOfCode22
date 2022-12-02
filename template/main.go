package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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
	return 112
}

func part2(input string) int{
	return 112
}
