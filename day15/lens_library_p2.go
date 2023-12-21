package day15

import (
	"advent_of_code_2023/day2"
	"fmt"
	"strconv"
	"strings"
)

const (
	HASH_MAP_SIZE int = 256
)

type lens struct {
	name        string
	focalLength int
}

func parseLens(input string) lens {
	val, _ := strconv.Atoi(string(input[len(input)-1]))
	return lens{
		name:        input[:len(input)-2],
		focalLength: val,
	}
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
		line = strings.ReplaceAll(line, "\n", "")
		//We remove from the map
		if line[len(line)-1] == '-' {
			line = strings.TrimSuffix(line, "-")
			index := calculateHash(line)
			toRemove := lens{
				name: line,
			}
			l.data[index].Remove(toRemove)
		} else {
			//We add to the map
			singleLens := parseLens(line)
			index := calculateHash(singleLens.name)
			l.data[index].Add(singleLens)
		}
	}
}

func calculatePower(hashMap *lensHashMap) int {
	sum := 0
	println("-------------------")

	for i, boxes := range hashMap.data {
		boxesOrder := boxes.reverseAsArray()
		if boxes.len >= 1 {
			fmt.Printf("[Box: %d]\n", i+1)

			for k, singleLens := range boxesOrder {
				fmt.Printf("\t - Lens %d [%s:%d] += [%d]\n", k+1, singleLens.name, singleLens.focalLength, (i+1)*(k+1)*(singleLens.focalLength))
				sum += (i + 1) * (k + 1) * (singleLens.focalLength)
			}
			fmt.Println("")
		}
	}
	println("-------------------")

	return sum
}

func GetFocusingPower() int {
	chunks := day2.TokenizeString(ReadFileAsString(INPUT_TEXT), ',')
	lhm := createLensHashMap()
	lhm.fillLensHashMap(chunks)

	//Debug print of all the boxes in order
	for i, list := range lhm.data {
		if list.len != 0 {
			fmt.Printf("%d. ", i+1)
			list.printReverse()
			println("")
		}
	}

	return calculatePower(lhm)
}
