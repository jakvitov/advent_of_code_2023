package day12

type Set[K comparable] struct {
	data map[K]bool
}

func CreateSet[K comparable]() *Set[K] {
	return &Set[K]{
		data: make(map[K]bool),
	}
}

func (s *Set[K]) Add(item K) {
	s.data[item] = true
}

func (s *Set[K]) Remove(item K) {
	delete(s.data, item)
}

func (s *Set[K]) IsPresent(item K) bool {
	return s.data[item]
}

func (s *Set[K]) GetLen() int {
	return len(s.data)
}

func (s *Set[K]) GetData() map[K]bool {
	return s.data
}
