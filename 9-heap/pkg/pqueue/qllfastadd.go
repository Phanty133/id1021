package pqueue

import (
	"errors"
	"github.com/phanty133/id1021/9-heap/pkg/llist"
)

type PQueueLLFastAdd[T comparable] struct {
	list *llist.LinkedList[PriorityQueueItem[T]]
}

func NewPQueueLLFastAdd[T comparable]() *PQueueLLFastAdd[T] {
	return &PQueueLLFastAdd[T]{
		list: llist.New[PriorityQueueItem[T]](),
	}
}

func (h *PQueueLLFastAdd[T]) Add(value T, priority int) {
	h.list.Add(PriorityQueueItem[T]{value, priority})
}

// Finds highest prio with lin search
func (h *PQueueLLFastAdd[T]) Remove() (T, error) {
	if h.Empty() {
		var zero T
		return zero, errors.New("heap empty")
	}

	var min *llist.LinkedListItem[PriorityQueueItem[T]]
	min = nil

	// Keep track of prev nodes to make removal O(1)
	var minPrevNode *llist.LinkedListItem[PriorityQueueItem[T]]
	var prevNode *llist.LinkedListItem[PriorityQueueItem[T]]
	minPrevNode = nil
	prevNode = nil

	for node := range h.list.Iter() {
		if min == nil || min.Head.Priority < node.Head.Priority {
			min = node
			minPrevNode = prevNode
		}

		prevNode = node
	}

	if minPrevNode == nil {
		h.list.SetFirst(min.Next())
	} else {
		minPrevNode.SetNext(min.Next())
	}

	return min.Head.Value, nil
}

func (h *PQueueLLFastAdd[T]) Empty() bool {
	return h.list.First() == nil
}
