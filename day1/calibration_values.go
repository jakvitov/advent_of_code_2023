package day1

import (
	"bufio"
	"log"
	"os"
)

const INPUT_FILE string = "day1/test_input.txt"

func ReadFileAsLines(input string) []string {
	result := make([]string, 0)
	file, err := os.Open(input)

	if err != nil {
		log.Fatal("Error while opening the file: " + input + " " + err.Error())
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}

func IsNum(char int32) bool {
	return char > 47 && char < 58
}

func composeNum(first, second int32) int32 {
	return (10 * first) + second
}

func decodeRow(input string) int32 {
	//Suppose -1 cannot be encoded into utf-8 value
	first := int32(-1)
	last := int32(-1)
	for _, char := range input {
		if IsNum(char) {
			if first == -1 {
				first = char - 48
				last = first
				continue
			}
			last = char - 48
		}
	}
	return composeNum(first, last)
}

func DecodeFile() int64 {
	result := int64(0)
	lines := ReadFileAsLines(INPUT_FILE)
	for _, line := range lines {
		result += int64(decodeRow(line))
	}
	return result
}
