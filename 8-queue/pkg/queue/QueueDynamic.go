package queue

import (
	"errors"
)

type QueueDynamic[T comparable] struct {
	data  []T
	front int
	back  int
	count int
}

func NewQueueDynamic[T comparable](size int) *QueueDynamic[T] {
	return &QueueDynamic[T]{
		data:  make([]T, size),
		back:  0,
		front: 0,
	}
}

func (q *QueueDynamic[T]) nextIndexWrapped(idx int) int {
	return mod((idx + 1), cap(q.data))
}

func (q *QueueDynamic[T]) Reallocate(newSize int) {
	newData := make([]T, newSize)

	if newSize < cap(q.data) {
		copy(newData, q.data[q.front:q.front+q.count])
	} else {
		endDist := cap(q.data) - q.front
		copy(newData[0:endDist], q.data[q.front:])
		copy(newData[endDist:], q.data[:q.front])
	}

	q.front = 0
	q.back = q.count
	q.data = newData
}

func (q *QueueDynamic[T]) Enqueue(val T) error {
	if q.back == q.front && !q.Empty() {
		q.Reallocate(cap(q.data) * 2)
	}

	q.data[q.back] = val
	q.back = q.nextIndexWrapped(q.back)
	q.count++

	return nil
}

func (q *QueueDynamic[T]) Dequeue() (T, error) {
	var zero T

	if q.Empty() {
		return zero, errors.New("queue is empty")
	}

	val := q.data[q.front]
	q.data[q.front] = zero
	q.front = q.nextIndexWrapped(q.front)
	q.count--

	if q.count < cap(q.data)/4 {
		q.Reallocate(cap(q.data) / 2)
	}

	return val, nil
}

func (q *QueueDynamic[T]) Empty() bool {
	var zero T
	return q.back == q.front && q.data[q.front] == zero
}
