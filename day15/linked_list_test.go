package day15

import (
	"fmt"
	"strings"
	"testing"
)

const TEST_SIZE int = 100000

func TestLinkedList(t *testing.T) {
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

func TestLinkedList2(t *testing.T) {
	l := CreateLinkedList()
	for i := 0; i < 13; i++ {
		lensName := strings.Repeat(string(byte('a')), i%12+1)
		testLens := lens{
			name:        lensName,
			focalLength: i + 1,
		}
		l.Add(testLens)
	}
	//l.printReverse()

	for i := 0; i < 13; i++ {
		lensName := strings.Repeat(string(byte('a')), i%12+1)
		testLens := lens{
			name:        lensName,
			focalLength: 1,
		}
		l.Add(testLens)
	}

	for i := 0; i < 13; i++ {
		lensName := strings.Repeat(string(byte('a')), i%12+1)
		testLens := lens{
			name: lensName,
		}
		l.Remove(testLens)
	}

	for i := 0; i < 13; i++ {
		lensName := strings.Repeat(string(byte('a')), i%12+1)
		testLens := lens{
			name: lensName,
		}
		l.Remove(testLens)
	}

	for i := 0; i < 13; i++ {
		lensName := strings.Repeat(string(byte('a')), i%12+1)
		testLens := lens{
			name:        lensName,
			focalLength: i + 1,
		}
		l.Add(testLens)
	}
	l.printReverse()

}

func TestLinkedList3(t *testing.T) {
	l := CreateLinkedList()
	for i := 0; i < 4; i++ {
		lensName := strings.Repeat("p", i+1)
		testLens := lens{
			name:        lensName,
			focalLength: i,
		}
		l.Add(testLens)
	}
	fmt.Printf("\n------------------------\n")
	l.printReverse()
	fmt.Printf("\n------------------------\n")
	toRemove := lens{
		name: "pppp",
	}
	l.Remove(toRemove)
	l.printReverse()
	fmt.Printf("\n------------------------\n")
	if l.Get(0).name == "pppp" {
		t.Error("Remove failed!")
	}
}

func TestLinkedList4(t *testing.T) {
	l := CreateLinkedList()
	for i := 0; i < 4; i++ {
		lensName := strings.Repeat("p", i+1)
		testLens := lens{
			name:        lensName,
			focalLength: i,
		}
		l.Add(testLens)
	}
	fmt.Printf("\n------------------------\n")
	l.printReverse()
	fmt.Printf("\n------------------------\n")
	toRemove := lens{
		name: "p",
	}
	l.Remove(toRemove)
	l.printReverse()
	fmt.Printf("\n------------------------\n")
	if l.reverseAsArray()[0].name == "p" {
		t.Error("Remove failed!")
	}
}

func TestLinkedList5(t *testing.T) {
	l := CreateLinkedList()
	for i := 0; i < 4; i++ {
		lensName := strings.Repeat("p", i+1)
		testLens := lens{
			name:        lensName,
			focalLength: i,
		}
		l.Add(testLens)
	}
	fmt.Printf("\n------------------------\n")
	l.printReverse()
	fmt.Printf("\n------------------------\n")
	toRemove := lens{
		name: "ppp",
	}
	l.Remove(toRemove)
	l.printReverse()
	fmt.Printf("\n------------------------\n")
	if l.reverseAsArray()[2].name == "p" {
		t.Error("Remove failed!")
	}
}
