package day15

import "fmt"

type Node struct {
	val  lens
	next *Node
}

func CreateNode(val lens) *Node {
	return &Node{
		val: val,
	}
}

type LinkedList struct {
	start *Node
	len   int
}

func CreateLinkedList() *LinkedList {
	return &LinkedList{
		len: 0,
	}
}

func (l *LinkedList) Add(item lens) {
	if l.len == 0 {
		l.start = CreateNode(item)
		l.len = 1
		return
	}

	for current := l.start; current != nil; current = current.next {
		if current.val.name == item.name {
			current.val.focalLength = item.focalLength
			return
		}
	}

	node := CreateNode(item)
	node.next = l.start
	l.start = node
	l.len += 1
}

func (l *LinkedList) Get(i int) *lens {
	for ptr := l.start; ptr != nil; ptr = ptr.next {
		if i == 0 {
			return &ptr.val
		}
		i--
	}
	return nil
}

func (l *LinkedList) Remove(input lens) bool {
	if l.len == 0 {
		return true
	}
	prev := l.start
	current := prev.next
	if current == nil && prev.val.Equals(input) {
		l.start = nil
		l.len = 0
		return true
	} else if current == nil {
		return false
	} else if prev.val.Equals(input) {
		l.start = l.start.next
		l.len -= 1
		return true
	}

	for current != nil {
		if current == nil && prev.val.Equals(input) {
			l.start = nil
			l.len = 0
			return true
		} else if current == nil {
			return false
		}

		if current.val.Equals(input) {
			prev.next = current.next
			l.len -= 1
			return true
		}
		prev = current
		current = prev.next
	}
	return false
}

func (l *LinkedList) reverseAsArray() []lens {
	result := make([]lens, l.len)
	current := l.start
	for i := l.len - 1; i >= 0; i-- {
		result[i] = current.val
		current = current.next
	}
	return result
}

func (l *LinkedList) printReverse() {
	fmt.Printf("[S]->")
	res := l.reverseAsArray()
	for _, oneLens := range res {
		fmt.Printf("[%s;%d]-->", oneLens.name, oneLens.focalLength)
	}
	fmt.Printf("[E]")
}
