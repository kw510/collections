package treeset

type Node[T comparable] struct {
	// Nil if root of a tree
	parent *Node[T]
	// Nil if empty
	child map[any]*Node[T]
}

// Makes a new, empty `treeset`.
func New[T comparable]() *Node[T] {
	return &Node[T]{
		child: map[any]*Node[T]{},
	}
}

// Makes a new `treeset` inserting values.
func NewFromValues[T comparable](values ...T) *Node[T] {
	n := &Node[T]{
		child: map[any]*Node[T]{},
	}

	for _, v := range values {
		n.InsertValue(v)
	}

	return n
}

// Makes a new `treeset` inserting nodes.
func NewFromNodes[T comparable](nodes ...*Node[T]) *Node[T] {
	n := &Node[T]{
		child: map[any]*Node[T]{},
	}

	for _, node := range nodes {
		n.InsertNode(node)
	}

	return n
}

// Returns direct children
func (n *Node[T]) Children() map[any]*Node[T] {
	return n.child
}

// Returns `true` if the set contains an element equal to the value.
func (n *Node[T]) Contains(value any) bool {
	_, ok := n.child[value]
	return ok
}

// Returns node from set
func (n *Node[T]) Get(value any) *Node[T] {
	return n.child[value]
}

// Adds a node to the set.
//
// Returns whether the node was newly inserted. That is:
//
//   - If the set did not previously contain an equal node, `true` is
//     returned.
//   - If the set already contained an equal node, `false` is returned, and
//     the entry is not updated.

func (n *Node[T]) InsertNode(node *Node[T]) bool {
	if n.Contains(node) {
		return false
	}

	node.parent = n
	n.child[node] = node

	return true
}

// Adds a value to the set.
//
// Returns whether the value was newly inserted. That is:
//
//   - If the set did not previously contain an equal value, `true` is
//     returned.
//   - If the set already contained an equal value, `false` is returned, and
//     the entry is not updated.
func (n *Node[T]) InsertValue(value T) bool {
	if n.Contains(value) {
		return false
	}

	node := New[T]()
	node.parent = n
	n.child[value] = node

	return true
}

// Returns `true` if the set contains no elements.
func (n *Node[T]) IsEmpty() bool {
	return n.Len() == 0
}

// Returns `true` if the set is a subset of another,
// i.e., `other` contains at least all the elements in `n`.
func (n *Node[T]) IsSubset(other *Node[T]) bool {
	if n == other {
		return true
	}

	if other.parent != nil && n.IsSubset(other.parent) {
		return true
	}

	return false
}

// Returns `true` if the set is a superset of another,
// i.e., `n` contains at least all the elements in `other`.
func (n *Node[T]) IsSuperSet(other *Node[T]) bool {
	if n == other {
		return true
	}

	if n.parent != nil && n.parent.IsSuperSet(other) {
		return true
	}

	return false
}

// Returns the number of elements in the set.
func (n *Node[T]) Len() int {
	return len(n.child)
}

// Returns direct superset
func (n *Node[T]) Parent() *Node[T] {
	return n.parent
}

// If the set contains an element equal to the value, removes it from the
// set and drops it. Returns whether such an element was present.
func (n *Node[T]) Remove(value any) bool {
	if !n.Contains(value) {
		return false
	}

	delete(n.child, value)
	return true
}
