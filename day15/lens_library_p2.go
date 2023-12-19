package day15

import (
	"advent_of_code_2023/day2"
	"strings"
)

const HASH_MAP_SIZE int = 256

type lens struct {
	name        string
	focalLength int
}

func (ls lens) Equals(a lens) bool {
	return ls.name == a.name
}

type lensHashMap struct {
	data []LinkedList
}

func createLensHashMap() *lensHashMap {
	return &lensHashMap{
		data: make([]LinkedList, HASH_MAP_SIZE),
	}
}

func (l *lensHashMap) fillLensHashMap(lines []string) {
	for _, line := range lines {
		index := calculateHash(line)
		if line[len(line)-1] == '-' {
			line = strings.TrimSuffix(line, "-")
			toRemove := lens{
				name: line,
			}
			l.data[index].Remove(toRemove)
		} else {

		}
	}
}

func GetFocusingPower() int {
	chunks := day2.TokenizeString(ReadFileAsString(INPUT_TEXT), ',')
	fillLensHashMap(chunks)
	return 0
}
