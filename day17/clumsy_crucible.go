package day17

import (
	"advent_of_code_2023/day1"
	"advent_of_code_2023/day10"
	"advent_of_code_2023/day12"
	"advent_of_code_2023/day16"
	"fmt"
)

const INPUT_TEXT string = "/home/jakub/Development/advent_of_code/advent_of_code_2023/day17/text_input2.txt"

var LEFT day10.Coord = day10.CoordOf(-1, 0)
var RIGHT day10.Coord = day10.CoordOf(1, 0)
var DOWN day10.Coord = day10.CoordOf(0, 1)
var UP day10.Coord = day10.CoordOf(0, -1)

type node struct {
	//Max 3 -> then change direction
	rowDirSize    int
	currentCost   int
	heurysticCost int
	prevCosts     int
	previous      *node
	direction     day10.Coord
	coord         day10.Coord
}

func sortNodes(a, b node) bool {
	return getTotalNodeValue(&a) < getTotalNodeValue(&b)
}

func Equals(a, b node) bool {
	return a.coord == b.coord && a.direction == b.direction && a.rowDirSize == b.rowDirSize
}

// Return if the current coord is the final one
func isFinal(m day16.Matrix, coord day10.Coord) bool {
	return coord.X == len(m[0])-1 && coord.Y == len(m)-1
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(input int) int {
	if input < 0 {
		return -input
	}
	return input
}

func Pow(a, b int) float64 {
	for ; b >= 0; b-- {
		a *= a
	}
	return float64(a)
}

func getTotalNodeValue(n *node) int {
	return n.currentCost + n.prevCosts + n.heurysticCost
}

// Use Eucleidian heurystics for the A* algorithm
func getHeuristic(m day16.Matrix, c day10.Coord) int {
	/*resX := len(m[0]) - 1
	resY := len(m) - 1
	distX := resX - c.X
	distY := resY - c.Y
	return Max(distX, distY)*/

	//return int(math.Sqrt(Pow(distX, 2)+Pow(distY, 2))) / 2

	//Heurystic turned off for normal Dijkstra
	return 0
}

func printPathInMatrix(m day16.Matrix, lastNode node) {
	backtrackMap := day12.CreateSet[day10.Coord]()
	for curr := lastNode; ; curr = *curr.previous {
		backtrackMap.Add(curr.coord)
		if curr.coord == day10.CoordOf(0, 0) {
			break
		}
	}

	for y := 0; y < len(m); y++ {
		fmt.Printf("| ")
		for x := 0; x < len(m[0]); x++ {
			char, _ := m.GetAt(day10.CoordOf(x, y))
			if backtrackMap.IsPresent(day10.CoordOf(x, y)) {
				fmt.Printf("{%s}", string(char))
				continue
			}
			fmt.Printf(" %s ", string(char))
		}
		fmt.Printf(" |\n")
	}
	println("")
}

// We use A* to find the optimal path
func findOptimalPath(input day16.Matrix) int {
	//Map coord to distance to the node
	assignedWeights := make(map[day10.Coord]int)
	toVisitQueue := PriorityQueueInit[node](sortNodes)

	start := node{
		currentCost:   0,
		heurysticCost: 0,
		prevCosts:     0,
		coord:         day10.CoordOf(0, 0),
	}
	toVisitQueue.Enqueue(start)
	maxQueueSize := 1

	for !toVisitQueue.IsEmpty() {

		if toVisitQueue.Count() > maxQueueSize {
			maxQueueSize = toVisitQueue.Count()
		}

		current, _ := toVisitQueue.Dequeue()
		if isFinal(input, current.coord) {
			printPathInMatrix(input, current)
			fmt.Printf("Priority queue max size: [%d]\n", toVisitQueue.Count())
			return current.prevCosts + current.currentCost
		}

		//printPathInMatrix(input, current)

		cost, found := assignedWeights[current.coord]
		if !found || cost > current.currentCost {
			//fmt.Printf("Assigning [%d:%d] -> {%d}\n", current.coord.X, current.coord.Y, current.currentCost+current.prevCosts)
			assignedWeights[current.coord] = current.currentCost + current.prevCosts
		}

		left, foundLeft := input.GetAt(current.coord.PlusCoord(LEFT))
		right, foundRight := input.GetAt(current.coord.PlusCoord(RIGHT))
		down, foundDown := input.GetAt(current.coord.PlusCoord(DOWN))
		up, foundUp := input.GetAt(current.coord.PlusCoord(UP))

		//All necessities for enqueueing of
		if (current.rowDirSize != 3 || (current.rowDirSize == 3 && current.direction != LEFT)) && foundLeft {
			rds := 1
			if current.direction == LEFT {
				rds = current.rowDirSize + 1
			}
			leftNode := node{
				rowDirSize:    rds,
				currentCost:   int(left) - '0',
				heurysticCost: getHeuristic(input, current.coord.PlusCoord(LEFT)),
				previous:      &current,
				direction:     LEFT,
				prevCosts:     current.currentCost + current.prevCosts,
				coord:         current.coord.PlusCoord(LEFT),
			}

			val, found := assignedWeights[leftNode.coord]
			valQ, foundQ := toVisitQueue.GetItem(leftNode, Equals)
			if (!found || val > getTotalNodeValue(&leftNode)) && (!foundQ || getTotalNodeValue(valQ) > getTotalNodeValue(&leftNode)) {
				toVisitQueue.Enqueue(leftNode)
			}
		}
		if (current.rowDirSize != 3 || (current.rowDirSize == 3 && current.direction != RIGHT)) && foundRight {
			rds := 1
			if current.direction == RIGHT {
				rds = current.rowDirSize + 1
			}
			rightNode := node{
				rowDirSize:    rds,
				currentCost:   int(right) - '0',
				direction:     RIGHT,
				heurysticCost: getHeuristic(input, current.coord.PlusCoord(RIGHT)),
				previous:      &current,
				prevCosts:     current.currentCost + current.prevCosts,
				coord:         current.coord.PlusCoord(RIGHT),
			}

			val, found := assignedWeights[rightNode.coord]
			valQ, foundQ := toVisitQueue.GetItem(rightNode, Equals)
			if (!found || val > getTotalNodeValue(&rightNode)) && (!foundQ || getTotalNodeValue(valQ) > getTotalNodeValue(&rightNode)) {
				toVisitQueue.Enqueue(rightNode)
			}
		}
		if (current.rowDirSize != 3 || (current.rowDirSize == 3 && current.direction != DOWN)) && foundDown {
			rds := 1
			if current.direction == DOWN {
				rds = current.rowDirSize + 1
			}
			downNode := node{
				rowDirSize:    rds,
				currentCost:   int(down) - '0',
				heurysticCost: getHeuristic(input, current.coord.PlusCoord(DOWN)),
				previous:      &current,
				direction:     DOWN,
				prevCosts:     current.currentCost + current.prevCosts,
				coord:         current.coord.PlusCoord(DOWN),
			}
			val, found := assignedWeights[downNode.coord]
			valQ, foundQ := toVisitQueue.GetItem(downNode, Equals)
			if (!found || val > getTotalNodeValue(&downNode)) && (!foundQ || getTotalNodeValue(valQ) > getTotalNodeValue(&downNode)) {
				toVisitQueue.Enqueue(downNode)
			}
		}
		if (current.rowDirSize != 3 || (current.rowDirSize == 3 && current.direction != UP)) && foundUp {
			rds := 1
			if current.direction == UP {
				rds = current.rowDirSize + 1
			}
			upNode := node{
				rowDirSize:    rds,
				currentCost:   int(up) - '0',
				heurysticCost: getHeuristic(input, current.coord.PlusCoord(UP)),
				previous:      &current,
				direction:     UP,
				prevCosts:     current.currentCost + current.prevCosts,
				coord:         current.coord.PlusCoord(UP),
			}
			val, found := assignedWeights[upNode.coord]
			valQ, foundQ := toVisitQueue.GetItem(upNode, Equals)
			if (!found || val > getTotalNodeValue(&upNode)) && (!foundQ || getTotalNodeValue(valQ) > getTotalNodeValue(&upNode)) {
				toVisitQueue.Enqueue(upNode)
			}
		}
	}
	//Assigned weight of the last node
	return assignedWeights[day10.CoordOf(len(input[0])-1, len(input)-1)]
}

func GetOptimalPathHeatLoss() int {

	m := day16.ParseMatrix(day1.ReadFileAsLines(INPUT_TEXT))
	return findOptimalPath(m)
}
