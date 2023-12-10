package day9

import (
	"advent_of_code_2023/day1"
	"advent_of_code_2023/day5"
)

const INPUT_FILE string = "day9/text_input.txt"

// Given input to extrapolate return its extrapolation value as int, that would be
// on the len(input + 1)th position
func extrapolateSequence(input []int) int {
	//Guard to end the recursion -> we end at constant array
	isConstant := true
	//List with the current differences
	diffs := make([]int, len(input)-1)
	for i := 1; i < len(input); i++ {
		//We add to diffs the current difference
		diffs[i-1] = input[i] - input[i-1]
		//If our array is not constant, we break
		if i > 2 && (diffs[i-1] != diffs[i-2]) {
			isConstant = false
		}
	}
	if isConstant {
		//We end the recursion with extrapolating the last item as itself
		return input[len(input)-1] + diffs[len(diffs)-1]
	}
	return input[len(input)-1] + extrapolateSequence(diffs)
}

func SumAllExtrapolatedValues() int {
	lines := day1.ReadFileAsLines(INPUT_FILE)
	result := 0
	for _, line := range lines {
		numbers := day5.TokenizeStringToInt(line, ' ')
		result += extrapolateSequence(numbers)
	}
	return result
}
