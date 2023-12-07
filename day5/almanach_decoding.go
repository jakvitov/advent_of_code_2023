package day5

import (
	"advent_of_code_2023/day1"
	"advent_of_code_2023/day2"
	"strconv"
	"strings"
)

const INPUT_FILE string = "advent_of_code_2023/day5/test_input.txt"

type layer struct {
	startsPrev    []int
	startsCurrent []int
	steps         []int
}

type seed struct {
	seeds []int
}

// todo optimalise this, so we do not iterate over whole mapping each pass trough
func (l *layer) mapNumber(input int) int {
	for index, strtPrev := range l.startsPrev {
		//We fall under the mapping range -> we map in it
		if input >= strtPrev && input < (strtPrev+l.steps[index]) {
			return l.startsCurrent[index] + (input - strtPrev)
		}
	}
	//We did not map the number in the layer -> default mapping n:n
	return input
}

// Apply layer to input data and map the input numbers to output
// Single pass trough layer
func (l *layer) passTrough(input []int) []int {
	result := make([]int, len(input))
	for index, item := range input {
		result[index] = l.mapNumber(item)
	}
	return result
}

// 10;11;88;93 to int[10,11,88,93]
func TokenizeStringToInt(line string, token int32) []int {
	strs := day2.TokenizeString(line, token)
	result := make([]int, len(strs))
	for i := 0; i < len(strs); i++ {
		result[i], _ = strconv.Atoi(strs[i])
	}
	return result
}

func parseSeed(line string) *seed {
	line = strings.TrimLeft(line, "seeds: ")
	seeds := TokenizeStringToInt(line, ' ')
	return &seed{seeds: seeds}
}

func parseLayers(lines []string) []layer {
	result := make([]layer, 0)
	currentLayer := layer{}
	for _, line := range lines {
		//We read empty line -> skip
		if len(line) == 0 {
			continue
		} else if !day1.IsNum(int32(line[0])) {
			if currentLayer.startsCurrent == nil {
				continue
			}
			result = append(result, currentLayer)
			//We read a new layer
			currentLayer = layer{}
			continue
		} else {
			//We read specifications of a layer
			rng := TokenizeStringToInt(line, ' ')
			currentLayer.startsCurrent = append(currentLayer.startsCurrent, rng[0])
			currentLayer.startsPrev = append(currentLayer.startsPrev, rng[1])
			currentLayer.steps = append(currentLayer.steps, rng[2])
		}
	}
	//If current layer is being read -> finish it
	if currentLayer.startsCurrent != nil {
		result = append(result, currentLayer)
	}
	return result
}

func findMin(i []int) int {
	if len(i) == 0 {
		panic("Min on empty array")
	}
	currentMin := i[0]
	for _, k := range i {
		if k < currentMin {
			currentMin = k
		}
	}
	return currentMin
}

func FindLowestLocation() int {
	lines := day1.ReadFileAsLines(INPUT_FILE)
	if len(lines) < 3 {
		panic("Invalid input length < 3")
	}
	s := parseSeed(lines[0])
	layers := parseLayers(lines[1:])
	for _, l := range layers {
		copy(s.seeds, l.passTrough(s.seeds))
	}
	return findMin(s.seeds)
}
