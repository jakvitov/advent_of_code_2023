package day10

//Implementation of stack for DFS algorithmn
//We use a singly linked list to get constant insertion and remove-top operations
type Stack[K any] struct {
	top *Node[K]
	len int
}

type Node[K any] struct {
	val  *K
	next *Node[K]
}

//Create an empty node to be pushed to stack
func createNode[K any](val *K) *Node[K] {
	return &Node[K]{
		val: val,
	}
}

//Create empty stack
func CreateStack[K any]() *Stack[K] {
	return &Stack[K]{
		len: 0,
	}
}

//Push item to stack
func (s *Stack[K]) Push(val *K) {
	//We init the stack
	if s.len == 0 {
		s.top = createNode[K](val)
		s.len = 1
		return
	}
	//We push value to already non empty stack
	n := createNode[K](val)
	n.next = s.top
	s.top = n
	s.len += 1
}

//Remove the top item from the stack
func (s *Stack[K]) Pop() *K {
	//Stack is empty -> we return nil
	if s.len == 0 {
		return nil
	}
	//Stack has only one item -> we need to set it to empty
	if s.len == 1 {
		result := s.top
		s.top = nil
		s.len = 0
		return result.val
	}
	//Top node of the stack to be returned
	result := s.top
	//
	s.top = s.top.next
	s.len -= 1
	return result.val
}
