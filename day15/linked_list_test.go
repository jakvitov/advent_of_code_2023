package day15

import (
	"testing"
)

const TEST_SIZE int = 100000

func TestLinkedLIst(t *testing.T) {
	l := CreateLinkedList()

	for i := 0; i < TEST_SIZE; i++ {
		asciiChar := byte(i % 255)

		testLens := lens{
			name:        string(asciiChar),
			focalLength: i,
		}
		l.Add(testLens)
	}

	for i := 0; i < 255; i++ {
		asciiChar := byte(i % 255)
		testLens := lens{
			name:        string(asciiChar),
			focalLength: 2,
		}
		l.Add(testLens)
	}

	l.printReverse()

	list := l.reverseAsArray()

	for i := 0; i < 255; i++ {
		if list[i].focalLength != 2 {
			t.Errorf("Linked list at %d is not 2, current val %d\n", i, list[i].focalLength)
		}
	}

	for i := 0; i < 255; i++ {
		asciiChar := byte(i % 255)
		testLens := lens{
			name:        string(asciiChar),
			focalLength: 2,
		}
		l.Remove(testLens)
	}

	if l.len != 0 {
		l.printReverse()
		t.Errorf("List not empty. List len: %d\n", l.len)
	}
}
