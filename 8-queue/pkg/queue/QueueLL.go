package queue

import (
	"errors"

	"github.com/phanty133/id1021/8-queue/pkg/llist"
)

type QueueLL[T comparable] struct {
	data *llist.LinkedList[T]
}

func NewQueueLL[T comparable]() *QueueLL[T] {
	return &QueueLL[T]{
		data: llist.New[T](),
	}
}

func (q *QueueLL[T]) Enqueue(val T) error {
	q.data.Append(val)
	return nil
}

func (q *QueueLL[T]) Dequeue() (T, error) {
	if q.Empty() {
		var zero T
		return zero, errors.New("queue is empty")
	}

	val := q.data.First().Head
	q.data.Remove(val)

	return val, nil
}

func (q *QueueLL[T]) Empty() bool {
	return q.data.First() == nil
}
