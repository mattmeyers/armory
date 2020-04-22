package ll

import (
	"fmt"
	"strings"
)

// Generic represents a generic value.
type Generic interface{}

// GenericDLL is a generic doubly linked list with references to the head and
// the tail.
type GenericDLL struct {
	head *dNode
	tail *dNode
	len  int
}

// dNode is a node in a DLL.
type dNode struct {
	data Generic
	next *dNode
	prev *dNode
}

// NewGenericDLL creates a new doubly linked list from the provided data.
func NewGenericDLL(data ...Generic) *GenericDLL {
	l := &GenericDLL{}

	for _, d := range data {
		l.Append(d)
	}

	return l
}

func newDNode(data Generic) *dNode {
	return &dNode{data: data}
}

func (l *GenericDLL) String() string {
	if l.len == 0 {
		return "[]"
	}

	sb := strings.Builder{}
	sb.WriteString("[")
	sb.WriteString(fmt.Sprintf("%v", l.head.data))
	n := l.head.next
	for n != nil {
		sb.WriteString(fmt.Sprintf(" %v", n.data))
		n = n.next
	}
	sb.WriteString("]")

	return sb.String()
}

// Len returns the length of the list.
func (l *GenericDLL) Len() int {
	return l.len
}

// IsEmpty returns if the list is empty. A linked list is empty when there are
// no nodes in it.
func (l *GenericDLL) IsEmpty() bool {
	return l.Len() == 0
}

// DeepCopy creates a new GenericDLL from this GenericDLL with no references to
// the original.
func (l *GenericDLL) DeepCopy() *GenericDLL {
	nl := NewGenericDLL()
	n := l.head
	for n != nil {
		nl.Append(n.data)
		n = n.next
	}
	return nl
}

// Enumerate returns a slice of the values in the list.
func (l *GenericDLL) Enumerate() []Generic {
	s := make([]Generic, l.Len())
	i := 0
	n := l.head
	for n != nil {
		s[i] = n.data
		i++
		n = n.next
	}
	return s
}

// Append adds a new node to the end of the list.
func (l *GenericDLL) Append(data Generic) {
	l.len++

	n := newDNode(data)
	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}

	l.tail.next = n
	n.prev = l.tail
	l.tail = n
}

// Prepend adds a new node to the front of the list.
func (l *GenericDLL) Prepend(data Generic) {
	l.len++

	n := newDNode(data)
	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}

	l.head.prev = n
	n.next = l.head
	l.head = n
}

// getNode retrieves the node with the provided index. The algorithm takes
// advantage of the head and tail pointers to start the search at the
// beginning or end of the list depending on the index.
func (l *GenericDLL) getNode(index int) *dNode {
	var m *dNode
	// If before the midway point, start from the head
	if l.len-index > index {
		i := 0
		m = l.head
		for i < index {
			m = m.next
			i++
		}
	} else {
		i := l.len - 1
		m = l.tail
		for i > index {
			m = m.prev
			i--
		}
	}
	return m
}

// Get returns the data at the provided index. This method panics if an out of
// range index is provided.
func (l *GenericDLL) Get(index int) Generic {
	if index < 0 || index > l.len-1 {
		panic("invalid index")
	}

	return l.getNode(index).data
}

// InsertAfter creates a new node in the list after the provided index. There
// is no way to add a node to the front of the list with this method. Use
// Prepend to achieve this behavior. This method panics if an out of range
// index is provided.
func (l *GenericDLL) InsertAfter(index int, data Generic) {
	if index < 0 || index > l.len-1 {
		panic("invalid index")
	}

	if index == l.len-1 {
		l.Append(data)
		return
	}

	n := newDNode(data)
	m := l.getNode(index)
	l.len++

	n.prev = m
	n.next = m.next
	m.next = n
	n.next.prev = n
}

// Remove deletes the node at the provided index and returns its value. This
// method panics if an out of range index is provided.
func (l *GenericDLL) Remove(index int) Generic {
	if index < 0 || index > l.len-1 {
		panic("invalid index")
	}

	n := l.getNode(index)
	l.len--

	if n.prev == nil {
		l.head = n.next
		if l.head != nil {
			l.head.prev = nil
		}
	} else {
		n.prev.next = n.next
	}

	if n.next == nil {
		l.tail = n.prev
		if l.tail != nil {
			l.tail.next = nil
		}
	} else {
		n.next.prev = n.prev
	}

	return n.data
}
