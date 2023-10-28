package pqueue

import (
	"errors"
)

type Heap[T any] struct {
	root *HeapNode[T]
}

type HeapNode[T any] struct {
	size     int
	priority int
	Value    T
	left     *HeapNode[T]
	right    *HeapNode[T]
}

func newHeapNode[T any](val T, priority int) *HeapNode[T] {
	return &HeapNode[T]{
		size:     0,
		Value:    val,
		priority: priority,
		left:     nil,
		right:    nil,
	}
}

func (h *Heap[T]) Size() int {
	if h.Empty() {
		return 0
	}

	return h.root.Size() + 1
}

func (h *HeapNode[T]) Size() int {
	return h.size
}

func (h *HeapNode[T]) Priority() int {
	return h.priority
}

func (h *Heap[T]) Empty() bool {
	return h.root == nil
}

func NewHeap[T any]() *Heap[T] {
	return &Heap[T]{root: nil}
}

func (h *Heap[T]) Add(val T, priority int) {
	if h.root == nil {
		h.root = newHeapNode[T](val, priority)
		return
	}

	h.root.Add(val, priority)
}

func (h *HeapNode[T]) Add(val T, priority int) {
	h.size++

	if priority < h.Priority() {
		h.Value, val = val, h.Value
		h.priority, priority = priority, h.priority
	}

	if h.left == nil {
		h.left = newHeapNode[T](val, priority)
		return
	}

	if h.right == nil {
		h.right = newHeapNode[T](val, priority)
		return
	}

	if h.left.Size() <= h.right.Size() {
		h.left.Add(val, priority)
		return
	}

	h.right.Add(val, priority)
}

func (h *Heap[T]) Remove() (T, error) {
	if h.Empty() {
		var zero T
		return zero, errors.New("heap is empty")
	}

	val := h.root.Value

	if h.root.size == 0 {
		h.root = nil
	} else {
		h.root.Remove()
	}

	return val, nil
}

func (h *HeapNode[T]) Remove() {
	h.size--

	var promoted *HeapNode[T]

	if h.left == nil {
		promoted = h.right
	} else if h.right == nil {
		promoted = h.left
	} else if h.right.Priority() < h.left.Priority() {
		promoted = h.right
	} else {
		promoted = h.left
	}

	h.Value = promoted.Value
	h.priority = promoted.priority

	if promoted.size == 0 {
		if promoted == h.left {
			h.left = nil
		} else {
			h.right = nil
		}
	} else {
		promoted.Remove()
	}
}

func (h *Heap[T]) Push(incr int) (int, error) {
	if h.Empty() {
		return 0, errors.New("heap is empty")
	}

	depth := h.root.push(incr)
	return depth, nil
}

func (h *HeapNode[T]) push(incr int) int {
	h.priority += incr
	depth := 0

	curNode := h
	var nextNode *HeapNode[T] = nil

	for curNode != nil {
		if curNode.left == nil {
			nextNode = curNode.right
		} else if curNode.right == nil {
			nextNode = curNode.left
		} else if curNode.left.Priority() < curNode.right.Priority() {
			nextNode = curNode.left
		} else {
			nextNode = curNode.right
		}

		if nextNode == nil || nextNode.Priority() > curNode.Priority() {
			break
		}

		curNode.Value, nextNode.Value = nextNode.Value, curNode.Value
		curNode.priority, nextNode.priority = nextNode.priority, curNode.priority
		depth++

		curNode = nextNode
	}

	return depth
}
