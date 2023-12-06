package day2

import (
	"advent_of_code_2023/day1"
	"strconv"
	"strings"
)

const (
	GREEN_LIMIT int    = 13
	BLUE_LIMIT  int    = 14
	RED_LIMIT   int    = 12
	INPUT_FILE  string = "day2/test_input.txt"
	COMMA       int32  = 44
	DELIM       int32  = 58
	SEMICOLON   int32  = 59
)

type hand struct {
	greenCount int
	blueCount  int
	redCount   int
}

type game struct {
	hands   []hand
	gameNum int
}

func (g *game) isGameValid() bool {
	for _, h := range g.hands {
		if h.blueCount > BLUE_LIMIT || h.redCount > RED_LIMIT || h.greenCount > GREEN_LIMIT {
			return false
		}
		println("----------")
		println(h.redCount)
		println(h.greenCount)
		println(h.blueCount)
		println("----------")
	}

	return true
}

func createInt(toReverse string) int {
	//casting it to int
	val, err := strconv.Atoi(toReverse)
	if err != nil {
		panic("Error while casting string to int " + err.Error())
	}
	return val
}

// Parse 10blue to 10 and color blue
func (h *hand) colorNumberToInt(input string) {
	numPart := strings.Builder{}
	for _, char := range input {
		//We are parsing a number not a color
		if char > 47 && char < 58 {
			//We do not yet know the length of the number - we create it in a reverse order and later reverse it
			numPart.WriteByte(byte(char))
		} else {
			//We started to read a color
			//Read the first letter and suspect, which is it (limited number of colors)
			switch char {
			//color red
			case 114:
				{
					h.redCount += createInt(numPart.String())
				}
				//color blue
			case 98:
				{
					h.blueCount += createInt(numPart.String())
				}
				//color green
			case 103:
				{
					h.greenCount += createInt(numPart.String())
				}
				//not recognised color
			default:
				{
				}
			}
			break
		}
	}
}

func parseHands(input string) []hand {
	resultHands := make([]hand, 0)
	input = strings.Replace(input, " ", "", -1)
	parsedInputToHands := tokenizeString(input, SEMICOLON)

	for _, hnd := range parsedInputToHands {
		//create a result hand
		resultHand := hand{
			greenCount: 0,
			blueCount:  0,
			redCount:   0,
		}
		colors := tokenizeString(hnd, COMMA)
		for _, color := range colors {
			//parse 10red to 10 and color red and add to current hand
			resultHand.colorNumberToInt(color)
		}
		resultHands = append(resultHands, resultHand)
	}
	return resultHands
}

func parseGame(inputLine string) *game {
	result := &game{
		gameNum: 0,
		hands:   make([]hand, 0),
	}

	//Parse to game n and the hands
	parsedDelim := tokenizeString(inputLine, DELIM)
	result.gameNum = convertGameStringToNum(parsedDelim[0])
	result.hands = parseHands(parsedDelim[1])
	return result
}

func tokenizeString(line string, token int32) []string {
	result := make([]string, 0)
	current := strings.Builder{}
	for _, char := range line {
		if char == token {
			result = append(result, current.String())
			current.Reset()
			continue
		}
		current.WriteByte(byte(char))
	}
	result = append(result, current.String())
	return result
}

// Convert Game n to int n
func convertGameStringToNum(game string) int {
	game = strings.TrimLeft(game, "Game ")
	i, err := strconv.Atoi(game)
	if err != nil {
		panic("Cannot parse game number " + game)
	}
	return i
}

func SumPossibleGames() int {
	read := day1.ReadFileAsLines(INPUT_FILE)
	resultSum := 0
	for _, line := range read {
		parsedGame := parseGame(line)
		if parsedGame.isGameValid() {
			resultSum += parsedGame.gameNum
		}
	}
	return resultSum
}
