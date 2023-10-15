package pqueue

import (
	"errors"
	"github.com/phanty133/id1021/9-heap/pkg/llist"
)

type PQueueLLFastRemove[T comparable] struct {
	list *llist.LinkedList[PriorityQueueItem[T]]
}

func NewPQueueLLFastRemove[T comparable]() *PQueueLLFastRemove[T] {
	return &PQueueLLFastRemove[T]{
		list: llist.New[PriorityQueueItem[T]](),
	}
}

func (h *PQueueLLFastRemove[T]) Add(value T, priority int) {
	newItem := PriorityQueueItem[T]{value, priority}

	// Inserts the item before the first node that's highest priority
	// such that equal-prio items maintain their relative order

	// Keep track of prev nodes to make removal O(1)
	var prevNode *llist.LinkedListItem[PriorityQueueItem[T]]
	prevNode = nil

	node := h.list.First()

	for node != nil {
		if node.Head.Priority < priority {
			break
		}

		prevNode = node
		node = node.Next()
	}

	newNode := &llist.LinkedListItem[PriorityQueueItem[T]]{Head: newItem}

	if prevNode == nil {
		newNode.SetNext(h.list.First())
		h.list.SetFirst(newNode)
	} else {
		newNode.SetNext(prevNode.Next())
		prevNode.SetNext(newNode)
	}
}

func (h *PQueueLLFastRemove[T]) Remove() (T, error) {
	if h.Empty() {
		var zero T
		return zero, errors.New("heap empty")
	}

	f := h.list.First()
	val := f.Head.Value
	h.list.SetFirst(f.Next())
	return val, nil
}

func (h *PQueueLLFastRemove[T]) Empty() bool {
	return h.list.First() == nil
}
