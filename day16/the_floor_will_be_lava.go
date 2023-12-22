package day16

import (
	"advent_of_code_2023/day1"
	"advent_of_code_2023/day10"
	"advent_of_code_2023/day12"
	"fmt"
)

const INPUT_TEXT string = "/home/jakub/Development/advent_of_code/advent_of_code_2023/day16/text_input2.txt"

// Structure representing one coordinates and direction pair
// This combination describes deterministically all future steps
type coordDirPair struct {
	crd day10.Coord
	dir day10.Coord
}

func coordDirPairOf(a, b day10.Coord) coordDirPair {
	return coordDirPair{
		crd: a,
		dir: b,
	}
}

// After sooo many previous one time solutions making universal matrix to save the data in :)
type Matrix [][]byte

func ParseMatrix(input []string) Matrix {
	data := make([][]byte, len(input))
	for i, line := range input {
		data[i] = make([]byte, len(line))
		for k, char := range line {
			data[i][k] = byte(char)
		}
	}
	return data
}

func (m Matrix) GetAt(input day10.Coord) (byte, bool) {
	if input.Y >= len(m) || input.X >= len(m[0]) || input.X < 0 || input.Y < 0 {
		//Error flag ->
		return 0, false
	}
	return m[input.Y][input.X], true
}

// Recursive function to get all the energized tiles
func (m Matrix) moveTile(current, direction day10.Coord, energizedTiles day12.Set[coordDirPair]) day12.Set[coordDirPair] {
	current = day10.CoordOf(current.X+direction.X, current.Y+direction.Y)

	char, found := m.GetAt(current)
	if !found || energizedTiles.IsPresent(coordDirPairOf(current, direction)) {
		//fmt.Printf("Stopping [%d:%d]\n", current.X, current.Y)
		return energizedTiles
	}
	energizedTiles.Add(coordDirPairOf(current, direction))

	switch char {
	case '.':
		energizedTiles = m.moveTile(current, direction, energizedTiles)
	case '|':
		if direction.X == 0 {
			energizedTiles = m.moveTile(current, direction, energizedTiles)
		} else {
			energizedTiles = m.moveTile(current, day10.CoordOf(0, -1), energizedTiles)
			energizedTiles = m.moveTile(current, day10.CoordOf(0, 1), energizedTiles)
		}
	case '-':
		if direction.Y == 0 {
			energizedTiles = m.moveTile(current, direction, energizedTiles)
		} else {
			energizedTiles = m.moveTile(current, day10.CoordOf(-1, 0), energizedTiles)
			energizedTiles = m.moveTile(current, day10.CoordOf(1, 0), energizedTiles)
		}
	case '/':
		energizedTiles = m.moveTile(current, day10.CoordOf(direction.Y*(-1), direction.X*(-1)), energizedTiles)
	case '\\':
		energizedTiles = m.moveTile(current, day10.CoordOf(direction.Y, direction.X), energizedTiles)
	default:
		panic("Unknown tile")
	}

	return energizedTiles
}

func (m Matrix) PrintMatrix() {
	for _, line := range m {
		fmt.Printf("| ")
		for _, char := range line {
			fmt.Printf(" %s ", string(char))
		}
		fmt.Printf(" |\n")
	}
}

func (m Matrix) getUniqueEnergizedTilesCount(energizedTiles day12.Set[coordDirPair]) int {
	data := energizedTiles.GetData()
	uniqueCoord := day12.CreateSet[day10.Coord]()
	for key := range data {
		uniqueCoord.Add(key.crd)
	}

	/*for y, line := range m {
		fmt.Printf("| ")
		for x, char := range line {
			if uniqueCoord.IsPresent(day10.CoordOf(x, y)) {
				fmt.Printf(" # ")
				continue
			}
			fmt.Printf(" %s ", string(char))
		}
		fmt.Printf(" |\n")
	}*/

	return uniqueCoord.GetLen()
}

func GetEnergizedTilesCount() int {
	energizedTiles := day12.CreateSet[coordDirPair]()
	lines := day1.ReadFileAsLines(INPUT_TEXT)
	m := ParseMatrix(lines)
	m.moveTile(day10.CoordOf(-1, 0), day10.CoordOf(1, 0), *energizedTiles)
	/*println("---------------------------------------------------")
	m.PrintMatrix()
	println("---------------------------------------------------")*/
	return m.getUniqueEnergizedTilesCount(*energizedTiles)
}
