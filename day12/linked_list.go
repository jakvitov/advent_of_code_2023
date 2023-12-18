package day12

import (
	"advent_of_code_2023/day10"
)

// Trivial singly linked list implementation
type LinkedList[K comparable] struct {
	start *day10.Node[K]
	len   int
}

func CreateLinkedList[K comparable]() *LinkedList[K] {
	return &LinkedList[K]{
		len: 0,
	}
}

func (l *LinkedList[K]) Add(item K) {
	if l.len == 0 {
		l.len = 1
		l.start = day10.CreateNode[K](&item)
		return
	}

	nd := day10.CreateNode[K](&item)
	nd.Next = l.start
	l.start = nd
	l.len += 1
}

// Return true if item was present
func (l *LinkedList[K]) Remove(item K) bool {
	if l.len == 0 {
		return false
	}

	var prevNode *day10.Node[K]
	currentNode := l.start

	//We remove the start
	if *currentNode.Val == item {
		l.len -= 1
		l.start = l.start.Next
		return true
	}

	for i := 0; i < l.len; i++ {
		prevNode = currentNode
		currentNode = currentNode.Next
		if *currentNode.Val == item {
			prevNode.Next = currentNode.Next
			l.len -= 1
			return true
		}
	}

	return false
}

func (l *LinkedList[K]) IsPresent(item K) bool {
	current := l.start
	for ; current != nil; current = current.Next {
		if *current.Val == item {
			return true
		}
	}
	return false
}
