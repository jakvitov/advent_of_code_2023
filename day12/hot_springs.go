package day12

import (
	"advent_of_code_2023/day1"
	"advent_of_code_2023/day2"
	"advent_of_code_2023/day5"
	"fmt"
	"strings"
)

const INPUT_TEXT string = "day12/text_input2.txt"

func isFinished(line string) bool {
	return !strings.Contains(line, "?")
}

func isCorrect(line string, brokenSequences []int) bool {
	/*if line == ".#....#...###." {
		print("")
	}*/

	seqs := day2.TokenizeString(line, '.')
	if len(seqs) != len(brokenSequences) {
		return false
	}
	for i, seq := range seqs {
		if len(seq) != brokenSequences[i] {
			return false
		}
		if strings.Contains(seq, "?") {
			return false
		}
	}
	return true
}

func getSumOfSubstring(line string, sum, index int, brokenSeqences []int) int {
	//fmt.Println(line)
	//End of recursion -> we are a valid version of a line
	if isFinished(line) && isCorrect(line, brokenSeqences) {
		//fmt.Printf("Valid line: %s\n", line)
		return 1
	} else if isFinished(line) || index >= len(line) {
		//We are finished, but not valid
		return 0
	}

	byteLine := []byte(line)
	byteLine2 := []byte(line)
	if byteLine[index] != '?' {
		return sum + getSumOfSubstring(line, sum, index+1, brokenSeqences)
	}

	byteLine[index] = '.'
	byteLine2[index] = '#'
	res := getSumOfSubstring(string(byteLine), sum, index+1, brokenSeqences)
	res2 := getSumOfSubstring(string(byteLine2), sum, index+1, brokenSeqences)
	//fmt.Printf("%s -> %d\n", string(byteLine), res)

	//fmt.Printf("%s -> %d\n", string(byteLine2), res2)

	return sum + res + res2
}

func GetDifferentArrangementsNum() int {
	sum := 0
	lines := day1.ReadFileAsLines(INPUT_TEXT)

	for i, line := range lines {
		parsedLines := day2.TokenizeString(line, ' ')
		arrs := getSumOfSubstring(parsedLines[0], 0, 0, day5.TokenizeStringToInt(parsedLines[1], ','))
		fmt.Printf("Line [%d] -> %d ways\n", i+1, arrs)
		sum += arrs
	}
	return sum
}
