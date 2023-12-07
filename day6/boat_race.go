package day6

import (
	"advent_of_code_2023/day1"
	"advent_of_code_2023/day5"
	"strings"
)

const INPUT_FILE string = "day6/test_input.txt"
const TIME_BEGINNING string = "Time:"
const DISTANCE_BEGINNING string = "Distance:"

// Structure representing sheet with the info about times and records
type sheet struct {
	times   []int
	records []int
}

// todo - optimalise this, so we get rid of quadratic time
func getMoreThanRecForEntry(time, record int) int {
	result := 0
	//We change time, so we need a constant here
	stopFlag := time / 2
	for charge := 0; charge <= stopFlag; charge++ {
		//We can beat the record with this combination
		if time*charge > record {
			result += 1
		}
		time -= 1
	}
	//The problem is symmetric charge * time = time * charge, so we calculate only the first half
	//If we have even number of nums, we count the half twice (causing result += 1 once more, so we remove it
	result = result * 2
	if (time % 2) == 0 {
		result -= 1
	}
	return result
}

// Get how many of the times, we can get better than the record
func (s *sheet) getMoreThanRecForSheet() int {
	result := 1
	for i, time := range s.times {
		println(getMoreThanRecForEntry(time, s.records[i]))
		result *= getMoreThanRecForEntry(time, s.records[i])
	}
	return result
}

func parseInput(inputFile string) *sheet {
	lines := day1.ReadFileAsLines(inputFile)
	if len(lines) < 2 {
		panic("Invalid file read linenum < 2.")
	}
	lines[0] = strings.TrimLeft(lines[0], TIME_BEGINNING)
	lines[1] = strings.TrimLeft(lines[1], DISTANCE_BEGINNING)

	return &sheet{
		times:   day5.TokenizeStringToInt(lines[0], ' '),
		records: day5.TokenizeStringToInt(lines[1], ' '),
	}
}

func GetAllCombinations() int {
	sht := parseInput(INPUT_FILE)
	return sht.getMoreThanRecForSheet()
}
