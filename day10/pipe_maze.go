package day10

import (
	"advent_of_code_2023/day1"
	"fmt"
)

const INPUT_FILE string = "/home/jakub/Development/advent_of_code/advent_of_code_2023/day10/text_input.txt"

//Recursive search to find path back to S trough the cycle, with steps counting
//Return the size of the current branch back to S
//We reconstruct the original path with backtracking
//The map is passed as a pointer, so we have void returns with just editing the backtracking map.txt"

type maze struct {
	data  [][]int32
	start coord
}

type coord struct {
	x, y int
}

func coordOf(x, y int) coord {
	return coord{
		x: x,
		y: y,
	}
}

func coordEquals(c, c2 coord) bool {
	return c.x == c2.x && c.y == c2.y
}

func parseMaze(inputFile string) *maze {
	lines := day1.ReadFileAsLines(inputFile)
	result := &maze{
		data: make([][]int32, len(lines)),
	}

	for y, line := range lines {
		result.data[y] = make([]int32, len(line))
		for x, char := range line {
			//We keep the information about the start
			if char == 'S' {
				result.start = coordOf(x, y)
			}
			result.data[y][x] = char
		}
	}
	return result
}

func (m *maze) getChar(crd coord) int32 {
	//We are outside of the maze
	if crd.y > len(m.data) || crd.y < 0 || crd.x > len(m.data[crd.y]) || crd.x < 0 {
		return -1
	}
	return m.data[crd.y][crd.x]
}

// We use the Depth first serach to scan the array
func (m *maze) dfsSearch() map[coord]coord {

	backtrackingMap := make(map[coord]coord)
	visited := make(map[coord]bool)
	startVisits := 0
	stack := CreateStack[coord]()
	start := m.start

	bottom := coordOf(start.x, start.y+1)
	top := coordOf(start.x, start.y-1)
	right := coordOf(start.x+1, start.y)
	left := coordOf(start.x-1, start.y)

	stack.Push(&bottom)
	stack.Push(&top)
	stack.Push(&right)
	stack.Push(&left)
	visited[start] = true

	prev := start

	for {
		current := *stack.Pop()

		char := m.getChar(current)

		if char == 'S' && startVisits == 0 {
			startVisits += 1
			continue
		} else if char == 'S' && startVisits == 1 {
			return backtrackingMap
		} else if char == -1 || char == '.' {
			continue
		}

		//Mark current coord as visited
		if visited[current] {
			continue
		} else {
			visited[current] = true
		}

		backtrackingMap[current] = prev
		prev = current

		switch char {
		//We construct the up coord and down coord because of |, accordingly to the others
		case '|':
			top := coordOf(current.x, current.y-1)
			bottom := coordOf(current.x, current.y+1)
			stack.Push(&top)
			stack.Push(&bottom)

		case '-':
			right := coordOf(current.x+1, current.y)
			left := coordOf(current.x-1, current.y)
			stack.Push(&right)
			stack.Push(&left)

		case 'L':
			top := coordOf(current.x, current.y-1)
			right := coordOf(current.x+1, current.y)
			stack.Push(&top)
			stack.Push(&right)

		case 'J':
			top := coordOf(current.x, current.y-1)
			left := coordOf(current.x-1, current.y)
			stack.Push(&top)
			stack.Push(&left)

		case '7':
			bottom := coordOf(current.x, current.y+1)
			left := coordOf(current.x-1, current.y)
			stack.Push(&bottom)
			stack.Push(&left)

		case 'F':
			bottom := coordOf(current.x, current.y+1)
			right := coordOf(current.x+1, current.y)
			stack.Push(&bottom)
			stack.Push(&right)

		//Probably . or other not important letter -> we continue to next iteration and pop
		default:
			continue
		}

	}

}

// Print the backtracking reconstruction of the cycle
func (m *maze) reconstructCycle(backtrackingMap map[coord]coord) {

	for key, val := range backtrackingMap {
		fmt.Printf("[%c] -> [%c]\n", m.getChar(val), m.getChar(key))
	}
}

func GetFurthestNodeDistance() int {
	//How we got to each point [dest] -> origin
	mz := parseMaze(INPUT_FILE)

	backtrackingMap := mz.dfsSearch()

	mz.reconstructCycle(backtrackingMap)
	return 0
}
