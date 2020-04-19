package stack

import (
	"fmt"
	"strings"
)

type Generic interface{}

var Zero Generic = nil

type GenericStack struct {
	top  int
	vals []Generic
}

func NewGenericStack(vals ...Generic) *GenericStack {
	s := &GenericStack{top: -1, vals: []Generic{}}

	for _, v := range vals {
		s.Push(v)
	}

	return s
}

func (s *GenericStack) String() string {
	if s.top < 0 {
		return "[]"
	}

	sb := strings.Builder{}
	sb.WriteString("[ ")
	for i := s.top; i >= 0; i-- {
		sb.WriteString(fmt.Sprintf("%v ", s.vals[i]))
	}
	sb.WriteString("]")
	return sb.String()
}

func (s *GenericStack) IsEmpty() bool {
	return s.top < 0
}

func (s *GenericStack) Push(val Generic) {
	s.vals = append(s.vals, val)
	s.top++
}

func (s *GenericStack) Pop() Generic {
	if s.top < 0 {
		return Zero
	}

	n := s.vals[s.top]
	s.vals = s.vals[0:s.top]
	s.top--
	return n
}

func (s *GenericStack) Clear() {
	s.top = -1
	s.vals = s.vals[:0]
}

func (s *GenericStack) Peek() Generic {
	if len(s.vals) == 0 {
		return Zero
	}
	return s.vals[s.top]
}
