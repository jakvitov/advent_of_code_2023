package day16

import (
	"advent_of_code_2023/day1"
	"advent_of_code_2023/day10"
	"advent_of_code_2023/day12"
)

func Max(vals []int) int {
	max := vals[0]
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

// This can be optimised by caching each path calculated from a single point and directions
func GetBestConfigurationEnergizedCount() int {
	m := ParseMatrix(day1.ReadFileAsLines(INPUT_TEXT))
	sums := make([]int, 0)

	//
	for i := 0; i < len(m); i++ {
		energizedTiles := day12.CreateSet[coordDirPair]()
		lines := day1.ReadFileAsLines(INPUT_TEXT)
		m := ParseMatrix(lines)
		m.moveTile(day10.CoordOf(-1, i), day10.CoordOf(1, 0), *energizedTiles)
		sums = append(sums, m.getUniqueEnergizedTilesCount(*energizedTiles))
	}

	for i := 0; i < len(m); i++ {
		energizedTiles := day12.CreateSet[coordDirPair]()
		lines := day1.ReadFileAsLines(INPUT_TEXT)
		m := ParseMatrix(lines)
		m.moveTile(day10.CoordOf(len(m[0]), i), day10.CoordOf(-1, 0), *energizedTiles)
		sums = append(sums, m.getUniqueEnergizedTilesCount(*energizedTiles))
	}

	for i := 0; i < len(m[0]); i++ {
		energizedTiles := day12.CreateSet[coordDirPair]()
		lines := day1.ReadFileAsLines(INPUT_TEXT)
		m := ParseMatrix(lines)
		m.moveTile(day10.CoordOf(i, -1), day10.CoordOf(0, 1), *energizedTiles)
		sums = append(sums, m.getUniqueEnergizedTilesCount(*energizedTiles))
	}

	for i := 0; i < len(m[0]); i++ {
		energizedTiles := day12.CreateSet[coordDirPair]()
		lines := day1.ReadFileAsLines(INPUT_TEXT)
		m := ParseMatrix(lines)
		m.moveTile(day10.CoordOf(i, len(m)), day10.CoordOf(0, -1), *energizedTiles)
		sums = append(sums, m.getUniqueEnergizedTilesCount(*energizedTiles))
	}

	return Max(sums)
}
