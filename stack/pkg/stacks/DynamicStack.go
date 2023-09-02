package stacks

import "fmt"

type DynamicStack[T any] struct {
	size int
	ip   int
	data []T
}

func NewDynamicStack[T any](initialSize int) *DynamicStack[T] {
	return &DynamicStack[T]{
		size: initialSize,
		ip:   -1,
		data: make([]T, initialSize),
	}
}

func (stack *DynamicStack[T]) Reallocate(newSize int) {
	newData := make([]T, newSize)
	copy(newData, stack.data)
	stack.data = newData
	stack.size = newSize

	// fmt.Printf("Reallocated stack to size %d\n", newSize)
}

func (stack *DynamicStack[T]) Push(value T) error {
	if stack.ip == stack.size - 1 {
		stack.Reallocate(stack.size * 2)
	}

	stack.ip++
	stack.data[stack.ip] = value
	return nil
}

func (stack *DynamicStack[T]) Pop() (T, error) {
	if stack.ip == 0 {
		var zero T
		return zero, fmt.Errorf("stack is empty")
	}

	val := stack.data[stack.ip]
	stack.ip--

	if stack.ip < stack.size / 4 {
		stack.Reallocate(stack.size / 2)
	}

	return val, nil
}

func (stack *DynamicStack[T]) Empty() bool {
	return stack.ip == -1
}
