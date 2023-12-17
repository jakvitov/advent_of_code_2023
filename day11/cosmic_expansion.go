package day11

import (
	"advent_of_code_2023/day1"
	"advent_of_code_2023/day10"
)

const INPUT_TEXT string = "/home/jakub/Development/advent_of_code/advent_of_code_2023/day11/text_input2.txt"

// Struct representing the space with galaxies
type space struct {
	data       [][]int32
	galaxies   []day10.Coord
	galaxyRows map[int]bool
	galaxyCols map[int]bool
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func parseGalaxiesAsMaze(input string) *space {
	lines := day1.ReadFileAsLines(INPUT_TEXT)
	result := &space{
		data:     make([][]int32, len(lines)),
		galaxies: make([]day10.Coord, 0),
	}
	for y, line := range lines {
		result.data[y] = make([]int32, len(line))
		for x, char := range line {
			if char == '#' {
				result.galaxies = append(result.galaxies, day10.CoordOf(x, y))
			}
			result.data[y][x] = char
		}
	}
	//Two sets representing rows with galaxies and columns with present galaxies
	result.galaxyRows = make(map[int]bool)
	result.galaxyCols = make(map[int]bool)
	for _, key := range result.galaxies {
		result.galaxyRows[key.Y] = true
		result.galaxyCols[key.X] = true
	}

	return result
}

func (s *space) getRowAdditions(first, second day10.Coord) int {
	sum := 0
	for i := min(first.Y, second.Y); i < max(first.Y, second.Y); i++ {
		if !s.galaxyRows[i] {
			sum += 1
		}
	}
	return sum
}

func (s *space) getColumnAdditions(first, second day10.Coord) int {
	sum := 0
	for i := min(first.X, second.X); i < max(first.X, second.X); i++ {
		if !s.galaxyCols[i] {
			sum += 1
		}
	}
	return sum
}

func (s *space) getGalaxyPairDistances() int {
	sum := 0

	//We iterate over all unique pairs
	for i := 0; i < len(s.galaxies); i++ {
		for k := i + 1; k < len(s.galaxies); k++ {
			first := &s.galaxies[i]
			second := &s.galaxies[k]
			xDiff := (max(first.X, second.X) - min(first.X, second.X)) + s.getColumnAdditions(*first, *second)
			yDiff := (max(first.Y, second.Y) - min(first.Y, second.Y)) + s.getRowAdditions(*first, *second)
			//dist := int(math.Sqrt(math.Pow(float64(xDiff), 2) + math.Pow(float64(yDiff), 2)))
			//Over which axis should we increment more due to space expansion
			dist := xDiff + yDiff
			//fmt.Printf("{%d.G}[%d:%d] -> {%d.G}[%d:%d] == %d\n", i+1, first.X, first.Y, k+1, second.X, second.Y, dist)
			sum += dist
		}
	}
	return sum
}

func GetSumOfShortestPaths() int {
	spc := parseGalaxiesAsMaze(INPUT_TEXT)

	return spc.getGalaxyPairDistances()
}
