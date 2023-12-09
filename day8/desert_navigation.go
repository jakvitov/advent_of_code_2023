package day8

import (
	"advent_of_code_2023/day1"
	"advent_of_code_2023/day2"
	"strings"
)

const INPUT_FILE string = "day8/test_input3.txt"

type node struct {
	left  string
	right string
	val   string
}

// Mathematical graph representation using a hash map insted of standard pointer one
// that allows us easier construction of the whole graph with nodes randomly pointing at each other
// no need to care about constructed node pointing to not yet created one
type graph struct {
	start *node
	//We keep this map for quickly finding node in the given tree
	valMap map[string]*node
}

func parseNode(input string) *node {
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, "(", "")
	input = strings.ReplaceAll(input, ")", "")
	halfs := day2.TokenizeString(input, '=')
	lrParse := day2.TokenizeString(halfs[1], ',')
	return &node{
		val:   halfs[0],
		left:  lrParse[0],
		right: lrParse[1],
	}
}

// We do not return pointer, since graph is just one map pointer and start node pointer
func parseGraph(input []string) graph {
	result := graph{
		valMap: make(map[string]*node, len(input)),
	}
	for _, line := range input {
		currentNode := parseNode(line)
		//We read the start value
		if currentNode.val == "AAA" {
			result.start = currentNode
		}
		result.valMap[currentNode.val] = currentNode
	}
	return result
}

func GetStepsCount() int {
	lines := day1.ReadFileAsLines(INPUT_FILE)
	instructions := lines[0]
	//Suppose the second line is always empty
	grph := parseGraph(lines[2:])
	currentVal := grph.start
	steps := 0
	//We suppose, that we cannot cycle
	for currentVal.val != "ZZZ" {
		//We cycle trough the instructions
		instruction := instructions[steps%len(instructions)]
		if instruction == 'R' {
			currentVal = grph.valMap[currentVal.right]
		} else {
			//Go left
			currentVal = grph.valMap[currentVal.left]
		}
		steps += 1
	}
	return steps
}
