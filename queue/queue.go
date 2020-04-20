package queue

import (
	"fmt"
	"strings"
)

type Generic interface{}

var Zero Generic = nil

type GenericQueue struct {
	vals []Generic
}

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

func (q *GenericQueue) IsEmpty() bool {
	return len(q.vals) == 0
}

func (q *GenericQueue) Push(val Generic) {
	q.vals = append(q.vals, val)
}

func (q *GenericQueue) Pop() Generic {
	if q.IsEmpty() {
		return Zero
	}

	front := q.vals[0]
	copy(q.vals, q.vals[1:])
	q.vals = q.vals[:len(q.vals)-1]

	return front
}
