# Armory

Armory is a collection of data structures implemented for the builtin data types.

## Installation

```
go get -u github.com/mattmeyers/armory
```

## Set

A set is a collection of unique elements. An Armory set provides the methods

```go
String() string
Contains(val T) bool
IsSubset(b *Set) bool
IsEmpty() bool
Len() int
Cap() int
SetCap(c int) *Set
Enumerate() []T
Add(val T) *Set
Remove(val T) *Set
Clear() *Set
Equals(s2 *Set) bool
Map(f func(T) T) *Set
Filter(f func(T) bool) *Set
Fold(base T, f func(T, T) T) T
Union(s2 *Set) *Set
Intersect(s2 *Set) *Set
Diff(s2 *Set) *Set
SymDiff(s2 *Set) *Set
```

Set methods can be chained together for more complicated behavior.

```go
v := set.NewIntSet(1,2,3).
  Add(4).
  Add(5).
  Map(func(a int) int { return 2 * a }).
  Remove(10).
  Filter(func(a int) bool { return a <= 6 }).
  Remove(8).
  Fold(0, func(a int, b int) int { return a + b })

fmt.Println(v) // 12
```
