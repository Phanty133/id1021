package queue

import (
	"errors"
)

type QueueStatic[T comparable] struct {
	data  []T
	front int
	back  int
}

func NewQueueStatic[T comparable](size int) *QueueStatic[T] {
	return &QueueStatic[T]{
		data:  make([]T, size),
		back:  0,
		front: 0,
	}
}

func (q *QueueStatic[T]) nextIndexWrapped(idx int) int {
	return mod((idx + 1), cap(q.data))
}

func (q *QueueStatic[T]) Enqueue(val T) error {
	if q.back == q.front && !q.Empty() {
		return errors.New("queue is full")
	}

	q.data[q.back] = val
	q.back = q.nextIndexWrapped(q.back)

	return nil
}

func (q *QueueStatic[T]) Dequeue() (T, error) {
	var zero T

	if q.Empty() {
		return zero, errors.New("queue is empty")
	}

	val := q.data[q.front]
	q.data[q.front] = zero
	q.front = q.nextIndexWrapped(q.front)

	return val, nil
}

func (q *QueueStatic[T]) Empty() bool {
	var zero T
	return q.back == q.front && q.data[q.front] == zero
}
