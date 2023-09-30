package tree

import (
	"cmp"
)

type BinaryTreeNode[K cmp.Ordered, V any] struct {
	Key   K
	Val   V
	Left  *BinaryTreeNode[K, V]
	Right *BinaryTreeNode[K, V]
}

type BinaryTree[K cmp.Ordered, V any] struct {
	Root *BinaryTreeNode[K, V]
}

func NewBinaryTreeNode[K cmp.Ordered, V any](key K, val V) *BinaryTreeNode[K, V] {
	return &BinaryTreeNode[K, V]{key, val, nil, nil}
}

func NewBinaryTree[K cmp.Ordered, V any]() *BinaryTree[K, V] {
	return &BinaryTree[K, V]{nil}
}

func (t *BinaryTree[K, V]) Lookup(key K) (V, bool) {
	return t.Root.lookup(key)
}

func (node *BinaryTreeNode[K, V]) lookup(key K) (V, bool) {
	if node == nil {
		var zero V
		return zero, false
	}

	if node.Key == key {
		return node.Val, true
	}

	if key < node.Key {
		return node.Left.lookup(key)
	}

	return node.Right.lookup(key)
}

func (t *BinaryTree[K, V]) Add(key K, val V) {
	if t.Root == nil {
		t.Root = NewBinaryTreeNode(key, val)
		return
	}

	t.Root.add(key, val)
}

// Returns true if the key was added, false if it was updated.
func (node *BinaryTreeNode[K, V]) add(key K, val V) bool {
	if node.Key == key {
		node.Val = val
		return false
	}

	if key < node.Key {
		if node.Left == nil {
			node.Left = NewBinaryTreeNode(key, val)
			return true
		}

		return node.Left.add(key, val)
	}

	if node.Right == nil {
		node.Right = NewBinaryTreeNode(key, val)
		return true
	}

	return node.Right.add(key, val)
}

func (t *BinaryTree[K, V]) DepthFirstList() []V {
	return t.Root.depthFirstList()
}

func (node *BinaryTreeNode[K, V]) depthFirstList() []V {
	if node == nil {
		return []V{}
	}

	return append(append(node.Left.depthFirstList(), node.Val), node.Right.depthFirstList()...)
}

func (t *BinaryTree[K, V]) DepthFirstIterator() <-chan *BinaryTreeNode[K, V] {
	ch := make(chan *BinaryTreeNode[K, V])

	go func() {
		t.Root.depthFirstIterator(ch)
		close(ch)
	}()

	return ch
}

func (node *BinaryTreeNode[K, V]) depthFirstIterator(ch chan<- *BinaryTreeNode[K, V]) {
	if node == nil {
		return
	}

	node.Left.depthFirstIterator(ch)
	ch <- node
	node.Right.depthFirstIterator(ch)
}
