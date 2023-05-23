package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	unprocessedLines := strings.Split(string(content), "\n")
	processLines := make([]string, len(unprocessedLines))
	for i, line := range unprocessedLines {
		processLines[i] = strings.ReplaceAll(strings.TrimSpace(line), "\r", "")
	}
	return processLines
}

func partA() {
	//read the input file
	inputs := readInput("day1-input.txt")
	currentDepth, err := strconv.Atoi(inputs[0])
	if err != nil {
		fmt.Println(err)
		fmt.Printf("error converting this value %s", inputs[0])
		return
	}
	count := 0

	for _, inputLine := range inputs {
		prevDepth := currentDepth
		currentDepth, err = strconv.Atoi(inputLine)
		if err != nil {
			println("Error converting %s input to int", inputLine)
		}
		if currentDepth > prevDepth {
			count++
		}
	}
	println("Count: %d", count)
}

func partB() {
	inputs := readInput("day1-input.txt")
	//inputs := readInput("day1-sampleInput.txt")
	currentCombinations := make([]int, 3)
	for i, inputLine := range inputs[0:3] {
		currentCombinations[i], _ = strconv.Atoi(inputLine)
	}
	currentDepth := currentCombinations[0] + currentCombinations[1] + currentCombinations[2]
	println(currentDepth)
	count := 0
	newInputs := inputs[1:]
	for i, inputLine := range newInputs[:len(newInputs)-3] {
		firstNumber, _ := strconv.Atoi(inputLine)
		secondNumber, _ := strconv.Atoi(newInputs[i+1])
		thirdNumber, _ := strconv.Atoi(newInputs[i+2])
		fmt.Printf("First: %d, Second: %d, Third: %d\n", firstNumber, secondNumber, thirdNumber)
		previousDepth := currentDepth
		currentDepth = firstNumber + secondNumber + thirdNumber
		println(currentDepth)
		println(previousDepth)
		println(count)
		println("----")
		if currentDepth > previousDepth {
			count++
		}
	}
	fmt.Printf("Count: %d\n", count)
}

func main() {

	//partA()
	partB()
}
