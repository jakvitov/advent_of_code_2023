package day12

import "testing"

func TestSet(t *testing.T) {
	s := CreateSet[int]()
	for i := 0; i < 10000000; i++ {
		s.Add(i + 265)
		if !s.isPresent(i + 265) {
			t.Error("I is not present!")
		}
		s.Remove(i + 265)
		if s.isPresent(i + 265) {
			t.Error("I is present after remove")
		}
	}
}
