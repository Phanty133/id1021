package stacks

type Stack[T any] interface {
	Push(val T) error
	Pop() (T, error)
	Empty() bool
}