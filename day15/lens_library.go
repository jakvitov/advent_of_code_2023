package day15

import (
	"advent_of_code_2023/day2"
	"fmt"
	"os"
)

const INPUT_TEXT string = "/home/jakub/go/advent_of_code_2023/day15/text_input2.txt"

func ReadFileAsString(input string) string {
	data, err := os.ReadFile(input)

	if err != nil {
		panic("Cannot open file. " + err.Error())
	}
	return string(data)
}

func calculateHash(input string) int {
	hash := 0
	for _, char := range input {
		if char == '\n' {
			continue
		}
		hash += int(char)
		hash *= 17
		hash %= 256
	}
	return hash
}

func GetHashSum() int {
	res := 0
	chunks := day2.TokenizeString(ReadFileAsString(INPUT_TEXT), ',')
	for _, chunk := range chunks {
		hash := calculateHash(chunk)
		fmt.Printf("[%s] -> H[%d]\n", chunk, hash)
		res += hash
	}
	return res
}
