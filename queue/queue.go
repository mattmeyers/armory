package queue

import (
	"fmt"
	"strings"
)

// Generic aliases the empty interface.
type Generic interface{}

// Zero is the default zero value for a queue.
var Zero Generic = nil

// GenericQueue implements a queue data structure with a slice.
type GenericQueue struct {
	vals []Generic
}

// NewGenericQueue creates a new queue of generic values.
func NewGenericQueue(vals ...Generic) *GenericQueue {
	q := &GenericQueue{vals: []Generic{}}

	for _, v := range vals {
		q.Push(v)
	}

	return q
}

func (q *GenericQueue) String() string {
	if q.IsEmpty() {
		return "[]"
	}

	sb := strings.Builder{}
	sb.WriteString("[ ")
	for _, v := range q.vals {
		sb.WriteString(fmt.Sprintf("%v ", v))
	}
	sb.WriteString("]")
	return sb.String()
}

// IsEmpty checks if the queue does not contain any elements.
func (q *GenericQueue) IsEmpty() bool {
	return len(q.vals) == 0
}

// Push appends an element to the queue.
func (q *GenericQueue) Push(val Generic) {
	q.vals = append(q.vals, val)
}

// Pop removes the element in the front of the queue and returns the value.
func (q *GenericQueue) Pop() Generic {
	if q.IsEmpty() {
		return Zero
	}

	front := q.vals[0]
	q.vals = q.vals[1:]

	return front
}

// Peek returns the value in the front of the queue without removing the
// element.
func (q *GenericQueue) Peek() Generic {
	if q.IsEmpty() {
		return Zero
	}

	return q.vals[0]
}
