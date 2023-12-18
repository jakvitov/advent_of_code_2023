package day12

import "testing"

const TEST_SIZE int = 100

func TestLinkedList(t *testing.T) {
	ll := CreateLinkedList[int]()
	for i := 0; i < TEST_SIZE; i++ {
		ll.Add(i)
		k := i
		if !ll.IsPresent(k) {
			t.Error("Tested int is not present!")
			return
		}
	}

	for i := 0; i < TEST_SIZE; i++ {
		ll.Remove(i)
		k := i
		if ll.IsPresent(k) {
			t.Errorf("Tested already deleted is present {%d}\n", k)
			return
		}
	}
}
