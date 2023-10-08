package queue

type Queue[T comparable] interface {
	Enqueue(val T) error
	Dequeue() (T, error)
	Empty() bool
}
