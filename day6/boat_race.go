package day6

import (
	"advent_of_code_2023/day1"
	"advent_of_code_2023/day5"
)

const INPUT_FILE string = "/home/jakub/go/advent_of_code_2023/day6/test_input.txt"
const TIME_BEGINNING string = "Time:"
const DISTANCE_BEGINNING string = "Distance:"

//Structure representing sheet with the info about times and records
type sheet struct {
	times   []int
	records []int
}

//todo - optimalise this, so we get rid of quadratic time
func getMoreThanRecForEntry(time, record int) int {
	result := 0
	for charge := 0; charge < time/2; charge++ {
		//We can beat the record with this combination
		if time*(time-charge) > record {
			result += 1
		}
	}
	//The problem is symmetric charge * time = time * charge, so we calculate only the first half
	return result * 2
}

//Get how many of the times, we can get better than the record
func (s *sheet) getMoreThanRecForSheet() int {
	result := 0
	for i, time := range s.times {
		result += getMoreThanRecForEntry(time, s.records[i])
	}
	return result
}

func parseInput(inputFile string) *sheet {
	lines := day1.ReadFileAsLines(inputFile)
	if len(lines) < 2 {
		panic("Invalid file read linenum < 2.")
	}

	return &sheet{
		times:   day5.TokenizeStringToInt(lines[0], ' '),
		records: day5.TokenizeStringToInt(lines[1], ' '),
	}
}

func GetAllCombinations() int {
	sht := parseInput(INPUT_FILE)
	return sht.getMoreThanRecForSheet()
}
