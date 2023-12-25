package day17

import (
	"advent_of_code_2023/day1"
	"advent_of_code_2023/day10"
	"advent_of_code_2023/day12"
	"advent_of_code_2023/day16"
	"strconv"
)

const INPUT_FILE string = "/home/jakub/Development/advent_of_code/advent_of_code_2023/day17/text_input.txt"

var UP day10.Coord = day10.CoordOf(0, -1)
var DOWN day10.Coord = day10.CoordOf(0, 1)
var LEFT day10.Coord = day10.CoordOf(-1, 0)
var RIGHT day10.Coord = day10.CoordOf(1, 0)

type node struct {
	coord            day10.Coord
	direction        day10.Coord
	edgeCost         int
	tenativeDistance int
	distRow          int
}

func Equals(a, b node) bool {
	return a.coord == b.coord
}

func sortNodes(a, b node) bool {
	return a.tenativeDistance < b.tenativeDistance
}

func isFinal(a node, m day16.Matrix) bool {
	return a.coord.X == len(m[0])-1 && a.coord.Y == len(m)-1
}

// Find the shortest path len from [0,0] to [n, m]
// We use edited Dijkstra algorithm to confirm to the rules
func getMinimalDistance(matrix day16.Matrix) int {
	queue := PriorityQueueInit[node](sortNodes)
	visited := day12.CreateSet[day10.Coord]()
	lens := make(map[day10.Coord]int)

	start := node{
		coord: day10.CoordOf(0, 0),
	}
	queue.Enqueue(start)

	for visited.GetLen() != len(matrix)*len(matrix[0]) {
		current, _ := queue.Dequeue()

		left := current.coord.PlusCoord(LEFT)
		right := current.coord.PlusCoord(RIGHT)
		up := current.coord.PlusCoord(UP)
		down := current.coord.PlusCoord(DOWN)
		if !visited.IsPresent(left) && (current.direction != LEFT || current.distRow < 3) {
			distRow := 1
			if current.direction == LEFT {
				distRow = current.distRow + 1
			}
			leftByte, foundLeft := matrix.GetAt(left)
			if foundLeft {
				edgeCostParsed, _ := strconv.Atoi(string(leftByte))
				leftNode := node{
					distRow:          distRow,
					coord:            left,
					direction:        LEFT,
					edgeCost:         edgeCostParsed,
					tenativeDistance: current.tenativeDistance + edgeCostParsed,
				}
				nd, found := queue.GetItem(leftNode, Equals)
				if !found {
					queue.Enqueue(leftNode)
				} else if nd.tenativeDistance > leftNode.tenativeDistance {
					queue.UpdatePriority(leftNode, Equals)
				}
			}
		}
		if !visited.IsPresent(right) && (current.direction != RIGHT || current.distRow < 3) {
			distRow := 1
			if current.direction == RIGHT {
				distRow = current.distRow + 1
			}
			rightByte, foundRight := matrix.GetAt(right)
			if foundRight {
				edgeCostParsed, _ := strconv.Atoi(string(rightByte))
				rightNode := node{
					distRow:          distRow,
					coord:            right,
					direction:        RIGHT,
					edgeCost:         edgeCostParsed,
					tenativeDistance: current.tenativeDistance + edgeCostParsed,
				}
				nd, found := queue.GetItem(rightNode, Equals)
				if !found {
					queue.Enqueue(rightNode)
				} else if nd.tenativeDistance > rightNode.tenativeDistance {
					queue.UpdatePriority(rightNode, Equals)
				}
			}
		}
		if !visited.IsPresent(up) && (current.direction != UP || current.distRow < 3) {
			distRow := 1
			if current.direction == UP {
				distRow = current.distRow + 1
			}
			upByte, foundUp := matrix.GetAt(up)
			if foundUp {
				edgeCostParsed, _ := strconv.Atoi(string(upByte))
				upNode := node{
					distRow:          distRow,
					coord:            up,
					direction:        UP,
					edgeCost:         edgeCostParsed,
					tenativeDistance: current.tenativeDistance + edgeCostParsed,
				}
				nd, found := queue.GetItem(upNode, Equals)
				if !found {
					queue.Enqueue(upNode)
				} else if nd.tenativeDistance > upNode.tenativeDistance {
					queue.UpdatePriority(upNode, Equals)
				}
			}
		}
		if !visited.IsPresent(down) && (current.direction != DOWN || current.distRow < 3) {
			distRow := 1
			if current.direction == DOWN {
				distRow = current.distRow + 1
			}
			downByte, foundDown := matrix.GetAt(down)
			if foundDown {
				edgeCostParsed, _ := strconv.Atoi(string(downByte))
				downNode := node{
					distRow:          distRow,
					coord:            down,
					direction:        DOWN,
					edgeCost:         edgeCostParsed,
					tenativeDistance: current.tenativeDistance + edgeCostParsed,
				}
				nd, found := queue.GetItem(downNode, Equals)
				if !found {
					queue.Enqueue(downNode)
				} else if nd.tenativeDistance > downNode.tenativeDistance {
					queue.UpdatePriority(downNode, Equals)
				}
			}
		}
		visited.Add(current.coord)
		lens[current.coord] = current.tenativeDistance
	}
	return lens[day10.CoordOf(len(matrix[0])-1, len(matrix)-1)]
}

func GetOptimalPathHeatLoss() int {
	matrix := day16.ParseMatrix(day1.ReadFileAsLines(INPUT_FILE))
	return getMinimalDistance(matrix)
}
