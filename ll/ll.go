package ll

import (
	"fmt"
	"strings"
)

type Generic interface{}

var Zero Generic = nil

type GenericLL struct {
	head *llNode
	tail *llNode
	len  int
}

type llNode struct {
	data Generic
	next *llNode
	prev *llNode
}

func NewGenericLL(data ...Generic) *GenericLL {
	l := &GenericLL{}

	for _, d := range data {
		l.Append(d)
	}

	return l
}

func newLLNode(data Generic) *llNode {
	return &llNode{data: data}
}

func (l *GenericLL) String() string {
	if l.len == 0 {
		return "[]"
	}

	sb := strings.Builder{}
	sb.WriteString("[ ")

	n := l.head
	for n != nil {
		sb.WriteString(fmt.Sprintf("%v ", n.data))
		n = n.next
	}
	sb.WriteString("]")

	return sb.String()
}

func (l *GenericLL) Append(data Generic) {
	l.len++

	n := newLLNode(data)
	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}

	l.tail.next = n
	n.prev = l.tail
	l.tail = n
}

func (l *GenericLL) Prepend(data Generic) {
	l.len++

	n := newLLNode(data)
	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}

	l.head.prev = n
	n.next = l.head
	l.head = n
}

func (l *GenericLL) Get(index int) Generic {
	if index < 0 || index > l.len-1 {
		panic("invalid index")
	}

	n := l.head
	for i := 0; i < index; i++ {
		n = n.next
	}

	return n.data
}

func (l *GenericLL) PopFront() Generic {
	if l.len == 0 {
		return Zero
	}

	l.len--
	node := l.head

	l.head = l.head.next

	// If the head is nil, the last element was removed and the tail
	// needs to be cleaned up
	if l.head != nil {
		l.head.prev = nil
	} else {
		l.tail = nil
	}

	return node.data
}

func (l *GenericLL) PopBack() Generic {
	if l.len == 0 {
		return Zero
	}

	l.len--
	node := l.tail

	l.tail = l.tail.prev

	// If the tail is nil, the last element was removed and the head
	// needs to be cleaned up
	if l.tail != nil {
		l.tail.next = nil
	} else {
		l.head = nil
	}

	return node.data
}
