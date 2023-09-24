package llist

import "errors"

type LinkedListStack[T comparable] struct {
	list *LinkedList[T]
}

func NewStack[T comparable]() *LinkedListStack[T] {
	return &LinkedListStack[T]{list: New[T]()}
}

func (s *LinkedListStack[T]) Push(value T) {
	s.list.Add(value)
}

func (s *LinkedListStack[T]) Pop() (T, error) {
	item := s.list.First()

	if item == nil {
		var result T
		return result, errors.New("stack is empty")
	}

	s.list.first = item.next

	return item.Head, nil
}
