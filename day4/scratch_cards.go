package day4

import (
	"advent_of_code_2023/day1"
	"advent_of_code_2023/day2"
	"math"
	"strconv"
	"strings"
)

const INPUT_FILE string = "day4/test_input.txt"

type card struct {
	winning map[int]bool
	present []int
}

func (c *card) getWinningPoints() int {
	sum := 0
	for _, i := range c.present {
		found := c.winning[i]
		if found {
			sum += 1
		}
	}
	if sum == 0 {
		return 0
	}
	return int(math.Pow(2, float64(sum-1)))
}

func parseWinning(input string) map[int]bool {
	result := make(map[int]bool)
	currentNum := ""

	for _, char := range input {
		//In case we end a number
		if char == ' ' && len(currentNum) != 0 {
			num, _ := strconv.Atoi(currentNum)
			currentNum = ""
			result[num] = true
		} else if char != ' ' {
			//We read a number, else we ignore
			currentNum += string(char)
		}
	}
	if len(currentNum) != 0 {
		num, _ := strconv.Atoi(currentNum)
		result[num] = true
	}
	return result
}

func parseMine(input string) []int {
	result := make([]int, 0)
	currentNum := ""

	for _, char := range input {
		//In case we end a number
		if char == ' ' && len(currentNum) != 0 {
			num, _ := strconv.Atoi(currentNum)
			currentNum = ""
			result = append(result, num)
		} else if char != ' ' {
			//We read a number, else we ignore
			currentNum += string(char)
		}
	}
	if len(currentNum) != 0 {
		num, _ := strconv.Atoi(currentNum)
		result = append(result, num)
	}
	return result
}

func parseCard(line string) *card {
	line = strings.TrimLeft(line, ":")
	halves := day2.TokenizeString(line, '|')
	winning := parseWinning(halves[0])
	mine := parseMine(halves[1])
	return &card{
		winning: winning,
		present: mine,
	}
}

func GetTotalCardsValue() int {
	winning := 0
	lines := day1.ReadFileAsLines(INPUT_FILE)
	for _, line := range lines {
		crd := parseCard(line)
		winning += crd.getWinningPoints()
	}
	return winning
}
