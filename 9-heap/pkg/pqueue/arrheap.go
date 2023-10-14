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

func NewArrHeap[T any]() *ArrHeap[T] {
	return &ArrHeap[T]{data: make([]ArrHeapNode[T], 0)}
}

func (h *ArrHeap[T]) Add(val T, priority int) {
	newNode := ArrHeapNode[T]{Value: val, priority: priority}
	h.data = append(h.data, newNode)
	h.bubble()
}

func (h *ArrHeap[T]) bubble() {
	i := len(h.data) - 1

	for i > 0 {
		parent := (i - 1) / 2

		if h.data[parent].priority <= h.data[i].priority {
			break
		}

		h.data[parent], h.data[i] = h.data[i], h.data[parent]
		i = parent
	}
}

func (h *ArrHeap[T]) Remove() (T, error) {
	if h.Empty() {
		var zero T
		return zero, errors.New("empty heap")
	}

	val := h.data[0].Value

	lastEl := len(h.data) - 1
	h.data[0] = h.data[lastEl]
	h.data = h.data[:lastEl]
	h.sink()

	return val, nil
}

func (h *ArrHeap[T]) sink() {
	i := 0
	dataLen := len(h.data)

	for i < dataLen {
		left := 2*i + 1
		right := 2*i + 2

		if left >= dataLen {
			break
		}

		if right >= dataLen {
			if h.data[left].priority < h.data[i].priority {
				h.data[left], h.data[i] = h.data[i], h.data[left]
			}

			break
		}

		if h.data[left].priority < h.data[right].priority {
			if h.data[left].priority <= h.data[i].priority {
				break
			}

			h.data[left], h.data[i] = h.data[i], h.data[left]
			i = left
		} else {
			if h.data[right].priority >= h.data[i].priority {
				break
			}

			h.data[right], h.data[i] = h.data[i], h.data[right]
			i = right
		}
	}
}
