package llist

type LinkedListItem[T comparable] struct {
	Head T
	next *LinkedListItem[T]
}

type LinkedList[T comparable] struct {
	first *LinkedListItem[T]
}

func New[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedListItem[T]) Next() *LinkedListItem[T] {
	return l.next
}

func (l *LinkedList[T]) First() *LinkedListItem[T] {
	return l.first
}

func (l *LinkedList[T]) Last() *LinkedListItem[T] {
	if l.first == nil {
		return nil
	}

	item := l.first

	for item.next != nil {
		item = item.next
	}

	return item
}

func (l *LinkedList[T]) Add(value T) {
	next := l.first
	item := &LinkedListItem[T]{Head: value, next: next}

	l.first = item
}

func (l *LinkedList[T]) Append(value T) *LinkedListItem[T] {
	item := &LinkedListItem[T]{Head: value}

	last := l.Last()

	if last == nil {
		l.first = item
		return item
	}

	last.next = item
	return item
}

func (l *LinkedList[T]) Length() int {
	if l.first == nil {
		return 0
	}

	item := l.first
	length := 1

	for item.next != nil {
		item = item.next
		length++
	}

	return length
}

func (l *LinkedList[T]) Find(value T) *LinkedListItem[T] {
	if l.first == nil {
		return nil
	}

	item := l.first

	for item != nil {
		if item.Head == value {
			return item
		}

		item = item.next
	}

	return nil
}

func (l *LinkedList[T]) Remove(value T) {
	if l.first == nil {
		return
	}

	item := l.first

	if item.Head == value {
		l.first = item.next
		return
	}

	for item.next != nil {
		if item.next.Head == value {
			item.next = item.next.next
			return
		}

		item = item.next
	}
}

func (l *LinkedList[T]) Unlink(item *LinkedListItem[T]) {
	if l.first == nil {
		return
	}

	if l.first == item {
		l.first = item.next
		return
	}

	prev := l.first

	for prev.next != nil {
		if prev.next == item {
			prev.next = item.next
			return
		}

		prev = prev.next
	}
}

func (l *LinkedList[T]) Insert(item *LinkedListItem[T]) {
	item.next = l.first
	l.first = item
}
