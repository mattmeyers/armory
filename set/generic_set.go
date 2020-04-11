package set

import (
	"fmt"

	"github.com/cheekybits/genny/generic"
)

//go:generate genny -in=generic_set.go -out=set.go gen "Generic=NUMBERS,string,rune"

// Generic represents a generic type. It is a alias of interface{}.
type Generic generic.Type

// GenericSet implements the set data structure implemented with a map. The set
// has an optional capacity. By default, this cap is set to zero. Setting a
//positive cap will cause silent failures whenever trying to add elements to a
// full set.
type GenericSet struct {
	vals map[Generic]bool
	cap  int
}

// NewGenericSet creates a new set of Generic values.
func NewGenericSet(vals ...Generic) *GenericSet {
	s := GenericSet{vals: map[Generic]bool{}}
	for _, v := range vals {
		s.vals[v] = true
	}
	return &s
}

func (s *GenericSet) String() string {
	return fmt.Sprintf("%v", s.Enumerate())
}

// Contains checks if the provided value is in the set.
func (s *GenericSet) Contains(val Generic) bool {
	return s.vals[val]
}

// IsSubset checks if the set is a subset of b i,e, every element of the set s
// is contained in b.
func (s *GenericSet) IsSubset(b GenericSet) bool {
	if len(s.vals) > len(b.vals) {
		return false
	}

	for k := range s.vals {
		if !b.vals[k] {
			return false
		}
	}

	return true
}

// IsEmpty checks of the set does not contain any elements.
func (s *GenericSet) IsEmpty() bool {
	return len(s.vals) == 0
}

// Len returns the number of elements in the set.
func (s *GenericSet) Len() int {
	return len(s.vals)
}

// Cap returns the capacity of the set.
func (s *GenericSet) Cap() int {
	return s.cap
}

// SetCap sets the capacity of the set. If a negative integer is passed, then
// the capacity is set to 0. A capacity of 0 indicates no capacity.
func (s *GenericSet) SetCap(c int) *GenericSet {
	if c < 0 {
		c = 0
	}

	s.cap = c
	return s
}

// Enumerate returns a slice of the elements in the set that can be used in a
// for ... range loop. Depending on what the data is being used for, consider
// using the Map, Filter, and Fold methods instead as they act on the set
// itself.
func (s *GenericSet) Enumerate() []Generic {
	keys := make([]Generic, len(s.vals))
	i := 0
	for k := range s.vals {
		keys[i] = k
		i++
	}
	return keys
}

// Add adds the provided val to the set. If the set has already reached
// capacity, then the addition silently fails.
func (s *GenericSet) Add(val Generic) *GenericSet {
	if s.cap > 0 && len(s.vals) == s.cap {
		return s
	}

	s.vals[val] = true
	return s
}

// Remove deletes the provided values from the set. If the value is not in the
// set, then this is a no-op.
func (s *GenericSet) Remove(val Generic) *GenericSet {
	delete(s.vals, val)
	return s
}

// Clear removes all elements from the set. This does not clear the capcity.
func (s *GenericSet) Clear() *GenericSet {
	for k := range s.vals {
		delete(s.vals, k)
	}
	return s
}

// Equals checks if the provided set is equal to this set. Sets A and B are
// equal if A is a subset of B and B is a subset of A.
func (s *GenericSet) Equals(s2 GenericSet) bool {
	if len(s.vals) != len(s2.vals) {
		return false
	}

	return s.IsSubset(s2)
}

// Map applies the provided function to every element in the set. Because the
// set must only contain unique elements, the size of the set may change
// after applying the transformation function.
func (s *GenericSet) Map(f func(Generic) Generic) *GenericSet {
	out := NewGenericSet().SetCap(s.cap)
	for k := range s.vals {
		out.vals[f(k)] = true
	}
	return out
}

// Filter applies the provided function to every element in the set. The
// resulting set contains only the elements that caused the function to
// return true.
func (s *GenericSet) Filter(f func(Generic) bool) *GenericSet {
	out := NewGenericSet().SetCap(s.cap)
	for k := range s.vals {
		if f(k) {
			out.vals[k] = true
		}
	}
	return out
}

// Fold reduces the set to a single value by successively applying the provided
// function to elements of the set beginning with the base. The operation
// performed in the provided function must be associative and communitive to
// guarantee a deterministic result.
func (s *GenericSet) Fold(base Generic, f func(Generic, Generic) Generic) Generic {
	for k := range s.vals {
		base = f(base, k)
	}
	return base
}

// Union returns the union of this set with the provided set. The union of two
// sets A and B is defined by:
//	A union B = {x | (x in A) or (x in B)}
func (s *GenericSet) Union(s2 GenericSet) *GenericSet {
	out := NewGenericSet()

	for k := range s.vals {
		out.vals[k] = true
	}

	for k := range s2.vals {
		out.vals[k] = true
	}

	return out
}

// Intersect returns the intersection of this set with the provided set. The
// intersection of two sets A and B is defined by:
//	A intersect B = {x | (x in A) and (x in B)}
func (s *GenericSet) Intersect(s2 GenericSet) *GenericSet {
	out := NewGenericSet()

	for k := range s.vals {
		if s2.vals[k] {
			out.vals[k] = true
		}
	}

	for k := range s2.vals {
		if s.vals[k] {
			out.vals[k] = true
		}
	}

	return out
}

// Diff returns the set difference of this set with the provided set. The
// set difference of two sets A and B is defined as:
//	A \ B = {x | (x in A) and (x not in B)}
func (s *GenericSet) Diff(s2 GenericSet) *GenericSet {
	out := NewGenericSet()

	for k := range s.vals {
		if !s2.vals[k] {
			out.vals[k] = true
		}
	}

	return out
}

// SymDiff returns the symmetric difference of this set with the provided set.
// The symmetric difference of two sets A and B is defined as:
//	A (+) B = {x | (x in A) xor (x in B)}
func (s *GenericSet) SymDiff(s2 GenericSet) *GenericSet {
	out := NewGenericSet()

	for k := range s.vals {
		if !s2.vals[k] {
			out.vals[k] = true
		}
	}

	for k := range s2.vals {
		if !s.vals[k] {
			out.vals[k] = true
		}
	}

	return out
}
