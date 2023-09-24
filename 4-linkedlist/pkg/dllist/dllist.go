package dllist

type DoublyLinkedListItem[T comparable] struct {
	Head T
	next *DoublyLinkedListItem[T]
	prev *DoublyLinkedListItem[T]
}

type DoublyLinkedList[T comparable] struct {
	first *DoublyLinkedListItem[T]
	last  *DoublyLinkedListItem[T]
}

func New[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

func (l *DoublyLinkedList[T]) First() *DoublyLinkedListItem[T] {
	return l.first
}

func (l *DoublyLinkedList[T]) Last() *DoublyLinkedListItem[T] {
	return l.last
}

func (l *DoublyLinkedListItem[T]) Next() *DoublyLinkedListItem[T] {
	return l.next
}

func (l *DoublyLinkedListItem[T]) Prev() *DoublyLinkedListItem[T] {
	return l.prev
}

func (l *DoublyLinkedList[T]) Add(value T) *DoublyLinkedListItem[T] {
	next := l.first
	item := &DoublyLinkedListItem[T]{Head: value, next: next}

	if next != nil {
		next.prev = item
	}

	l.first = item

	if l.last == nil {
		l.last = item
	}

	return item
}

func (l *DoublyLinkedList[T]) Length() int {
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

func (l *DoublyLinkedList[T]) Find(value T) *DoublyLinkedListItem[T] {
	if l.first == nil {
		return nil
	}

	item := l.first

	for item.next != nil {
		if item.Head == value {
			return item
		}

		item = item.next
	}

	return nil
}

func (l *DoublyLinkedList[T]) Remove(value T) {
	if l.first == nil {
		return
	}

	item := l.Find(value)

	if item == nil {
		return
	}

	l.Unlink(item)
}

func (l *DoublyLinkedList[T]) Unlink(item *DoublyLinkedListItem[T]) {
	if item.prev != nil {
		item.prev.next = item.next
	} else {
		l.first = item.next
	}

	if item.next != nil {
		item.next.prev = item.prev
	} else {
		l.last = item.prev
	}
}

func (l *DoublyLinkedList[T]) Insert(item *DoublyLinkedListItem[T]) {
	item.next = l.first
	l.first.prev = item
	l.first = item
}
