package day10

import (
	"advent_of_code_2023/day1"
)

const INPUT_FILE string = "/home/jakub/Development/advent_of_code/advent_of_code_2023/day10/text_input5.txt"

//Recursive search to find path back to S trough the cycle, with steps counting
//Return the size of the current branch back to S
//We reconstruct the original path with backtracking
//The map is passed as a pointer, so we have void returns with just editing the backtracking map.txt"

type maze struct {
	data  [][]int32
	start Coord
}

type Coord struct {
	X, Y int
}

func (c *Coord) PlusCoord(a Coord) Coord {
	return CoordOf(c.X+a.X, c.Y+a.Y)
}

func (c *Coord) PlusCoordNum(x, y int) Coord {
	return CoordOf(c.X+x, c.Y+y)
}

func CoordOf(x, y int) Coord {
	return Coord{
		X: x,
		Y: y,
	}
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
				result.start = CoordOf(x, y)
			}
			result.data[y][x] = char
		}
	}
	return result
}

func (m *maze) getChar(crd Coord) int32 {
	//We are outside of the maze
	if crd.Y > len(m.data) || crd.Y < 0 || crd.X > len(m.data[crd.Y]) || crd.X < 0 {
		return -1
	}
	return m.data[crd.Y][crd.X]
}

// We use the Depth first serach to scan the array
func (m *maze) dfsSearch() map[Coord]Coord {

	backtrackingMap := make(map[Coord]Coord)
	visited := make(map[Coord]bool)
	startVisits := 0
	stack := CreateStack[Coord]()
	start := m.start

	bottom := CoordOf(start.X, start.Y+1)
	top := CoordOf(start.X, start.Y-1)
	right := CoordOf(start.X+1, start.Y)
	left := CoordOf(start.X-1, start.Y)

	if m.getChar(bottom) == '|' || m.getChar(bottom) == 'L' || m.getChar(bottom) == 'J' {
		stack.Push(&bottom)
	}
	if m.getChar(right) == '-' || m.getChar(right) == '7' || m.getChar(right) == 'J' {
		stack.Push(&right)
	}
	if m.getChar(left) == 'F' || m.getChar(left) == '-' || m.getChar(left) == 'L' {
		stack.Push(&left)
	}
	if m.getChar(top) == '|' || m.getChar(top) == '7' || m.getChar(top) == 'F' {
		stack.Push(&top)
	}

	visited[start] = true

	prev := start

	for {
		current := *stack.Pop()

		char := m.getChar(current)

		if char == 'S' && startVisits == 0 {
			startVisits += 1
			continue
		} else if char == 'S' && startVisits == 1 {
			backtrackingMap[current] = prev
			return backtrackingMap
		} else if char == -1 || char == '.' {
			continue
		}

		//Mark current Coord as visited
		if visited[current] {
			continue
		} else {
			visited[current] = true
		}

		//How we got to this point
		backtrackingMap[current] = prev
		prev = current

		switch char {
		//We construct the up Coord and down Coord because of |, accordingly to the others
		case '|':
			top := CoordOf(current.X, current.Y-1)
			bottom := CoordOf(current.X, current.Y+1)
			stack.Push(&top)
			stack.Push(&bottom)

		case '-':
			right := CoordOf(current.X+1, current.Y)
			left := CoordOf(current.X-1, current.Y)
			stack.Push(&right)
			stack.Push(&left)

		case 'L':
			top := CoordOf(current.X, current.Y-1)
			right := CoordOf(current.X+1, current.Y)
			stack.Push(&top)
			stack.Push(&right)

		case 'J':
			top := CoordOf(current.X, current.Y-1)
			left := CoordOf(current.X-1, current.Y)
			stack.Push(&top)
			stack.Push(&left)

		case '7':
			bottom := CoordOf(current.X, current.Y+1)
			left := CoordOf(current.X-1, current.Y)
			stack.Push(&bottom)
			stack.Push(&left)

		case 'F':
			bottom := CoordOf(current.X, current.Y+1)
			right := CoordOf(current.X+1, current.Y)
			stack.Push(&bottom)
			stack.Push(&right)

		//Probably . or other not important letter -> we continue to next iteration and pop
		default:
			continue
		}

	}

}

// Print the backtracking reconstruction of the cycle
func (m *maze) reconstructCycle(backtrackingMap map[Coord]Coord) int {

	len := 0
	end := backtrackingMap[m.start]
	for {
		println(string(m.getChar(end)))
		len += 1
		if m.getChar(end) == 'S' {
			break
		}
		end = backtrackingMap[end]
	}
	//For backtracking print
	/**for key, val := range backtrackingMap {
		fmt.Printf("[%c] -> [%c]\n", m.getChar(val), m.getChar(key))
	}*/
	return len / 2
}

func GetFurthestNodeDistance() int {
	//How we got to each point [dest] -> origin
	mz := parseMaze(INPUT_FILE)

	backtrackingMap := mz.dfsSearch()

	return mz.reconstructCycle(backtrackingMap)
}
