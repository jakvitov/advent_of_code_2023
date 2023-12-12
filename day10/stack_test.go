package day10

import "testing"

const TEST_PUSH_SIZE int = 10000000

func TestStackPushPop(t *testing.T) {
	stack := CreateStack[string]()
	val := "jdlfjadkjf"
	for i := 0; i < TEST_PUSH_SIZE; i++ {
		stack.Push(&val)
	}
	if stack.len != TEST_PUSH_SIZE {
		t.Error("Sizes don't match!")
	}
	for i := 0; i < TEST_PUSH_SIZE; i++ {
		val = *stack.Pop()
	}
}
