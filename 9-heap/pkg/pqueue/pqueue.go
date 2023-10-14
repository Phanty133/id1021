package pqueue

type PriorityQueue[T comparable] interface {
	Add(val T, priority int)
	Remove() (T, error)
	Empty() bool
}

type PriorityQueueItem[T comparable] struct {
	Value    T
	Priority int
}

func IterPQueue[T comparable](h PriorityQueue[T]) chan T {
	ch := make(chan T)

	go func() {
		for !h.Empty() {
			val, err := h.Remove()

			if err != nil {
				panic("weirdo error in heap iteration")
			}

			ch <- val
		}

		close(ch)
	}()

	return ch
}
