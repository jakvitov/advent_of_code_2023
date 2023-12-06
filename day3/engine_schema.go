package day3

import (
	"advent_of_code_2023/day1"
	"os"
	"strconv"
)

const (
	INPUT_FILE string = "day3/test_input.txt"
	NEWLINE    byte   = '\n'
)

type schema struct {
	matrix [][]byte
}

func (s *schema) getByteAt(x, y int) byte {
	//In case of outside of the array, we return 0, so as number, it remains ignored
	if x < 0 || y < 0 || y > (len(s.matrix)-1) || x > (len(s.matrix[0])-1) {
		return byte(46)
	}
	//We rotate x and y, so it matches the matix dimensions
	return s.matrix[y][x]
}

func isSpecialSign(char byte) bool {
	return !((char > 47 && char < 58) || char == 46) && (char > 32 && char < 127)
}

func (s *schema) isSurroudingByteSpecialSymbol(x, y int) bool {
	surroundingTiles := make([]byte, 8)
	surroundingTiles[0] = s.getByteAt(x+1, y)
	surroundingTiles[1] = s.getByteAt(x+1, y+1)
	surroundingTiles[2] = s.getByteAt(x+1, y-1)
	surroundingTiles[3] = s.getByteAt(x-1, y)
	surroundingTiles[4] = s.getByteAt(x-1, y+1)
	surroundingTiles[5] = s.getByteAt(x-1, y-1)
	surroundingTiles[6] = s.getByteAt(x, y-1)
	surroundingTiles[7] = s.getByteAt(x, y+1)

	for _, char := range surroundingTiles {
		if isSpecialSign(char) {
			println(string(char))
			return true
		}
	}
	return false
}

// Read file and compose a matrix of bytes from it
func parseSchema(inputFile string) *schema {
	result := &schema{
		matrix: make([][]byte, 0),
	}
	fileBytes, err := os.ReadFile(inputFile)
	if err != nil {
		panic("Cannot read file " + err.Error())
	}

	row := make([]byte, 0)
	for _, char := range fileBytes {
		//We split at newline
		if char == NEWLINE {
			newRow := make([]byte, len(row))
			copy(newRow, row)
			result.matrix = append(result.matrix, newRow)
			row = make([]byte, 0)
		} else {
			//Or we just add a byte to the array
			row = append(row, char)
		}
	}
	if len(row) > 0 {
		result.matrix = append(result.matrix, row)
	}
	return result
}

func numberFromCurrentNum(currentNum string) int {
	if len(currentNum) == 0 {
		return 0
	}
	val, err := strconv.Atoi(currentNum)
	if err != nil {
		panic("Error while parsing number " + err.Error())
	}
	return val
}

func GetSumOfNumbers() int {
	sum := 0
	schem := parseSchema(INPUT_FILE)
	currentNum := ""
	hasSpecialNeightbour := false
	for y, row := range schem.matrix {
		for x, char := range row {
			//We read number, so we add it to the current number composition
			if day1.IsNum(int32(char)) {
				currentNum += string(char)
				if hasSpecialNeightbour == false {
					hasSpecialNeightbour = schem.isSurroudingByteSpecialSymbol(x, y)
				}
			} else if hasSpecialNeightbour {
				sum += numberFromCurrentNum(currentNum)
				currentNum = ""
				hasSpecialNeightbour = false
			} else {
				if len(currentNum) > 0 {
				}
				currentNum = ""
				hasSpecialNeightbour = false
			}
		}
	}
	return sum
}
