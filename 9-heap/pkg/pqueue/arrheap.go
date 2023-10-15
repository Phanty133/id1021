package pqueue

import "errors"

type ArrHeap[T any] struct {
	data []ArrHeapNode[T]
}

type ArrHeapNode[T any] struct {
	priority int
	Value    T
}

func (h *ArrHeap[T]) Size() int {
	return len(h.data)
}

func (h *ArrHeap[T]) Empty() bool {
	return len(h.data) == 0
}

func NewArrHeap[T any](prealloc int) *ArrHeap[T] {
	return &ArrHeap[T]{data: make([]ArrHeapNode[T], 0, prealloc)}
}

func (h *ArrHeap[T]) Add(val T, priority int) {
	newNode := ArrHeapNode[T]{Value: val, priority: priority}
	h.data = append(h.data, newNode)
	h.bubble()
}

func (h *ArrHeap[T]) bubble() {
	n := len(h.data) - 1

	for n > 0 {
		var parent int

		// Calculate the parent node
		if n%2 == 0 {
			parent = (n - 2) / 2
		} else {
			parent = (n - 1) / 2
		}

		if h.data[parent].priority <= h.data[n].priority {
			break
		}

		h.data[parent], h.data[n] = h.data[n], h.data[parent]
		n = parent
	}
}

func (h *ArrHeap[T]) Remove() (T, error) {
	if h.Empty() {
		var zero T
		return zero, errors.New("empty heap")
	}

	val := h.data[0].Value

	lastIdx := len(h.data) - 1
	h.data[0] = h.data[lastIdx] // Replace root value
	h.data = h.data[:lastIdx]   // Remove the last node
	h.sink()

	return val, nil
}

func (h *ArrHeap[T]) sink() {
	n := 0
	dataLen := len(h.data)

	for n < dataLen {
		left := 2*n + 1
		right := 2*n + 2

		if left >= dataLen {
			// No children
			break
		}

		if right >= dataLen {
			// Only left child

			if h.data[left].priority < h.data[n].priority {
				h.data[left], h.data[n] = h.data[n], h.data[left]
			}

			break
		}

		var targetIdx int

		if h.data[left].priority < h.data[right].priority {
			targetIdx = left
		} else {
			targetIdx = right
		}

		if h.data[targetIdx].priority > h.data[n].priority {
			break
		}

		h.data[targetIdx], h.data[n] = h.data[n], h.data[targetIdx]
		n = targetIdx
	}
}
