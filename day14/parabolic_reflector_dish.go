package day14

import (
	"advent_of_code_2023/day1"
	"fmt"
)

const INPUT_TEXT string = "day14/text_input2.txt"

func StringToByteMatrix(input []string) [][]byte {
	result := make([][]byte, len(input))
	i := 0
	for _, line := range input {
		result[i] = []byte(line)
		i += 1
	}
	return result
}

// Rotate matrix by 90 degrees clockwise
func RotateClockwise(input [][]byte) [][]byte {
	result := make([][]byte, len(input[0]))
	for i := range result {
		result[i] = make([]byte, len(input))
	}

	for i, line := range input {
		for k := range line {
			result[k][i] = input[i][k]
		}
	}
	//todo optimise by putting it into the previsous loop
	for i, line := range input {
		for k := 0; k < len(line)/2; k++ {
			temp := result[i][k]
			target := (len(line) - 1) - k
			result[i][k] = result[i][target]
			result[i][target] = temp
		}
	}
	return result
}

func moveRock(line []byte, index int) []byte {
	if len(line) == index {
		return line
	}
	if line[index] == '.' || line[index] == '#' {
		return moveRock(line, index+1)
	}
	//We have "O" ->  a round stone
	line = moveRock(line, index+1)
	for i := index + 1; i < len(line) && line[i] == '.'; i++ {
		line[i-1] = '.'
		line[i] = '0'
	}
	return line
}

func printMatrix(line [][]byte) {
	for i := 0; i < len(line); i++ {
		for k := 0; k < len(line[i]); k++ {
			fmt.Printf("%c", line[i][k])
		}
		fmt.Printf("\n")
	}
}

// Calculate the overall sum according to rules of rotated result
func calculateResults(platform [][]byte) int {
	sum := 0
	for i, line := range platform {
		for _, char := range line {
			if char == '0' || char == 'O' {
				sum += i + 1
			}
		}
	}
	return sum
}

func GetLoadAfterTilt() int {
	lines := day1.ReadFileAsLines(INPUT_TEXT)
	platform := StringToByteMatrix(lines)
	platform = RotateClockwise(platform)

	for i, line := range platform {
		platform[i] = moveRock(line, 0)
	}
	platform = RotateClockwise(platform)
	printMatrix(platform)
	return calculateResults(platform)
}
