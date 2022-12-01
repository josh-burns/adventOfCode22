package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	content, _ := ioutil.ReadFile("input.txt")
	var caloriesPerElf []int
	caloriesArray := strings.Split(string(content), "\n\n")

	for i := range caloriesArray {
		elfString := strings.Split(caloriesArray[i], "\n")
		var elfInt []int

		for a := range elfString {
			intVal, _ := strconv.Atoi(elfString[a])
			elfInt = append(elfInt, intVal)
		}
		runningTotal := 0

		for b := range elfInt {
			runningTotal += elfInt[b]
		}
		caloriesPerElf = append(caloriesPerElf, runningTotal)
	}

	sort.Ints(caloriesPerElf)
	// part 1:
	fmt.Println(caloriesPerElf[len(caloriesPerElf)-1])

	// part 2:
	top3 := caloriesPerElf[len(caloriesPerElf)-1] +
		caloriesPerElf[len(caloriesPerElf)-2] +
		caloriesPerElf[len(caloriesPerElf)-3]

	fmt.Println(top3)

}
