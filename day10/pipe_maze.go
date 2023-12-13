package day10

import (
	"advent_of_code_2023/day1"
	"fmt"
)

const INPUT_FILE string = "/home/jakub/go/advent_of_code_2023/day10/text_input.txt"

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

func (m *maze) dfs(visited map[coord]coord) map[coord]coord {

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

	visited[bottom] = start
	visited[top] = start
	visited[right] = start
	visited[left] = start

	for {
		current := *stack.Pop()

		if m.getChar(current) == 'S' {
			return visited
		} else if m.getChar(current) == -1 {
			continue
		}

		_, found := visited[current]
		if found {
			continue
		}

		switch m.getChar(current) {
		//We construct the up coord and down coord because of |, accordingly to the others
		case '|':
			top := coordOf(current.x, current.y-1)
			bottom := coordOf(current.x, current.y+1)
			visited[top] = current
			visited[bottom] = current
			stack.Push(&top)
			stack.Push(&bottom)

		case '-':
			right := coordOf(current.x+1, current.y)
			left := coordOf(current.x-1, current.y)
			visited[right] = current
			visited[left] = current
			stack.Push(&right)
			stack.Push(&left)

		case 'L':
			top := coordOf(current.x, current.y-1)
			right := coordOf(current.x+1, current.y)
			visited[top] = current
			visited[right] = current
			stack.Push(&right)
			stack.Push(&left)

		case 'J':
			top := coordOf(current.x, current.y-1)
			left := coordOf(current.x-1, current.y)
			visited[top] = current
			visited[left] = current
			stack.Push(&top)
			stack.Push(&left)

		case '7':
			bottom := coordOf(current.x, current.y+1)
			left := coordOf(current.x-1, current.y)
			visited[bottom] = current
			visited[left] = current
			stack.Push(&bottom)
			stack.Push(&left)

		case 'F':
			bottom := coordOf(current.x, current.y+1)
			right := coordOf(current.x+1, current.y)
			visited[bottom] = current
			visited[right] = current
			stack.Push(&bottom)
			stack.Push(&left)

		//Probably . or other not important letter -> we continue to next iteration and pop
		default:
			continue
		}
	}
}

// Print the backtracking reconstruction of the cycle
func (m *maze) reconstructCycle(backtrackingMap map[coord]coord) {
	current := backtrackingMap[m.start]

	for {
		prev := current
		next, found := backtrackingMap[current]
		if !found {
			break
		}

		fmt.Printf("[%d:%d]{%s} -> [%d:%d]{%s}\t", prev.x, prev.y, string(m.getChar(prev)), next.x, next.y, string(m.getChar(next)))
	}
}

func GetFurthestNodeDistance() int {
	//How we got to each point [dest] -> origin
	backtrackingMap := make(map[coord]coord)
	mz := parseMaze(INPUT_FILE)

	backtrackingMap = mz.dfs(backtrackingMap)

	mz.reconstructCycle(backtrackingMap)
	return 0
}
