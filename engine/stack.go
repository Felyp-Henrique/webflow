package engine

type Stack[T any] struct {
	_items []T
}

func (s *Stack[T]) Peek() T {
	return s._items[len(s._items)-1]
}

func (s *Stack[T]) Push(item T) {
	s._items = append(s._items, item)
}

func (s *Stack[T]) Pop() {
	if len(s._items) > 0 {
		s._items = s._items[:len(s._items)-1]
	}
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		_items: []T{},
	}
}
