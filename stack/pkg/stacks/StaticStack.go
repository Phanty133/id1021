package stacks

import "fmt"

type StaticStack[T any] struct {
	size int
	ip   int
	data []T
}

func NewStaticStack[T any](dataArr []T) *StaticStack[T] {
	return &StaticStack[T]{
		size: cap(dataArr),
		ip:   -1,
		data: dataArr,
	}
}

func (stack *StaticStack[T]) Push(value T) error {
	if stack.ip == stack.size-1 {
		return fmt.Errorf("stack is full")
	}

	stack.ip++
	stack.data[stack.ip] = value
	return nil
}

func (stack *StaticStack[T]) Pop() (T, error) {
	if stack.ip == 0 {
		var zero T
		return zero, fmt.Errorf("stack is empty")
	}

	val := stack.data[stack.ip]
	stack.ip--
	return val, nil
}

func (stack *StaticStack[T]) Empty() bool {
	return stack.ip == -1
}
