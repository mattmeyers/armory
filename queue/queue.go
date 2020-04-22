package queue

import (
	"fmt"
	"strings"
)

// Generic represents a generic value.
type Generic interface{}

// Zero is the default zero value for a queue.
var Zero Generic = nil

// GenericQueue is a generic queue that is implemented as a linked list.
type GenericQueue struct {
	head *dNode
	tail *dNode
	len  int
}

// dNode is a node in a queue.
type dNode struct {
	data Generic
	next *dNode
	prev *dNode
}

// NewGenericQueue creates a new queue of generic values.
func NewGenericQueue(vals ...Generic) *GenericQueue {
	q := &GenericQueue{}

	for _, v := range vals {
		q.Push(v)
	}

	return q
}

func newDNode(data Generic) *dNode {
	return &dNode{data: data}
}

func (q *GenericQueue) String() string {
	if q.IsEmpty() {
		return "[]"
	}

	sb := strings.Builder{}
	sb.WriteString("[ ")
	n := q.head
	for n != nil {
		sb.WriteString(fmt.Sprintf("%v ", n.data))
		n = n.next
	}
	sb.WriteString("]")

	return sb.String()
}

// Len returns the number of elements in the queue.
func (q *GenericQueue) Len() int {
	return q.len
}

// IsEmpty checks if the queue does not contain any elements.
func (q *GenericQueue) IsEmpty() bool {
	return q.Len() == 0
}

// Push appends an element to the queue.
func (q *GenericQueue) Push(val Generic) {
	q.len++

	n := newDNode(val)
	if q.head == nil {
		q.head = n
		q.tail = n
		return
	}

	q.tail.next = n
	n.prev = q.tail
	q.tail = n
}

// Pop removes the element in the front of the queue and returns the value.
// This method will panic when attempting to pop an empty queue.
func (q *GenericQueue) Pop() Generic {
	if q.IsEmpty() {
		panic("queue is empty")
	}

	n := q.head
	q.len--

	if n.prev == nil {
		q.head = n.next
		if q.head != nil {
			q.head.prev = nil
		}
	} else {
		n.prev.next = n.next
	}

	if n.next == nil {
		q.tail = n.prev
		if q.tail != nil {
			q.tail.next = nil
		}
	} else {
		n.next.prev = n.prev
	}

	return n.data
}

// Peek returns the value in the front of the queue without removing the
// element. This method will panic when attempting to peek an empty queue.
func (q *GenericQueue) Peek() Generic {
	if q.IsEmpty() {
		panic("empty queue")
	}

	return q.head.data
}
