package day15

type Comparator interface {
	Equals(a Comparator) bool
}

type Node struct {
	val  Comparator
	next *Node
}

func CreateNode(val Comparator) *Node {
	return &Node{
		val: val,
	}
}

type LinkedList struct {
	start *Node
	end   *Node
	len   int
}

func CreateLinkedList() *LinkedList {
	return &LinkedList{
		len: 0,
	}
}

func (l *LinkedList) Add(item Comparator) {
	if l.len == 0 {
		l.start = CreateNode(item)
		l.end = l.start
		l.len = 1
		return
	}

	node := CreateNode(item)
	node.next = l.start
	l.start = node
	l.len += 1
}

func (l *LinkedList) Get(i int) Comparator {
	for ptr := l.start; ptr != l.end; ptr = ptr.next {
		if i == 0 {
			return ptr.val
		}
		i--
	}
	return nil
}

func (l *LinkedList) Remove(input Comparator) bool {
	if l.len == 0 {
		return true
	}
	prev := l.start
	current := prev.next
	if current == nil && prev.val.Equals(input) {
		l.start = nil
		l.end = l.start
		l.len = 0
		return true
	} else if current == nil {
		return false
	}

	for current != nil {
		prev = current
		current = current.next

		if current.val.Equals(input) {
			if current.next == nil {
				l.end = prev
			}
			prev.next = current.next
			l.len -= 1
			return true
		}
	}
	return false
}
