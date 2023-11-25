package structures

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) All() []T {
	return s.items
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return item
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: []T{},
	}
}
