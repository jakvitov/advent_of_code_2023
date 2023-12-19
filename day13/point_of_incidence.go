package day13

import (
	"advent_of_code_2023/day1"
	"fmt"
)

const INPUT_TEXT string = "/home/jakub/Development/advent_of_code/advent_of_code_2023/day13/text_input2.txt"

// Parse input file to individual scatches
func parseScatches(lines []string) [][]string {
	scatches := make([][]string, 0)
	currentScatch := make([]string, 0)

	for _, line := range lines {
		if line == "" {
			scatches = append(scatches, currentScatch)
			currentScatch = make([]string, 0)
			continue
		}
		currentScatch = append(currentScatch, line)
	}
	if len(currentScatch) != 0 {
		scatches = append(scatches, currentScatch)
	}
	return scatches
}

// Check if given string is symmetric over a given index
func checkHorizontalSymmetry(line string, index int) bool {
	if index == len(line) {
		return false
	}
	k := index + 1
	for index >= 0 && k < len(line) {
		if line[index] != line[k] {
			return false
		}
		index -= 1
		k += 1
	}
	return true
}

// Check if array of lines is symmetric vertically
func checkVerticalSymmetry(lines []string, index int) bool {
	k := index + 1

	for index > 0 && k < len(lines) {
		if lines[index] != lines[k] {
			return false
		}
		index -= 1
		k += 1
	}
	return true
}

// Return the reflection index in one particular scatch
func findReflectionIndex(lines []string) int {
	//Horizontal part
	horizontalPossibilities := make(map[int]bool)
	firstLine := lines[0]
	for i := 1; i < len(firstLine); i++ {
		k := i - 1
		if firstLine[i] == firstLine[k] {
			horizontalPossibilities[k] = true
		}
	}

	for _, line := range lines {
		for index := range horizontalPossibilities {
			if len(horizontalPossibilities) == 0 {
				break
			}
			if !checkHorizontalSymmetry(line, index) {
				delete(horizontalPossibilities, index)
			}
		}
	}

	//We expect 0 or 1 here
	for pos := range horizontalPossibilities {
		//Indexing according to rules starts at 1
		return pos + 1
	}

	//If horizontal symmetry failed, we try vertical one
	verticalPossibilites := make(map[int]bool)
	for i := 1; i < len(lines); i++ {
		k := i - 1
		if lines[i] == lines[k] {
			verticalPossibilites[k] = true
		}
	}

	for index := range verticalPossibilites {
		if len(verticalPossibilites) == 0 {
			break
		}
		if !checkVerticalSymmetry(lines, index) {
			delete(verticalPossibilites, index)
		}
	}

	for i := range verticalPossibilites {
		//According to rules we multiple horizontal indexes by 100
		return 100 * (i + 1)
	}
	panic("No symmetry found for !")
}

func SummariseLinesOfReflections() int {
	sum := 0
	lines := day1.ReadFileAsLines(INPUT_TEXT)
	parsed := parseScatches(lines)
	for i, p := range parsed {
		refInd := findReflectionIndex(p)

		defer func(i int) {
			if r := recover(); r != nil {
				fmt.Printf("No symmetry for %d\n", i)
			}
		}(i)

		fmt.Printf("[%d] -> {%d}\n", i+1, refInd)
		sum += refInd
	}
	return sum
}
