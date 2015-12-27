// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	// Add 0 to words according to the length of word
	// e.g. [0] if x=[0-63]
	//      [0, 0] if x=[64-127]
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	// Bitwise OR, then left bit shift
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 { // word is not set.
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len return the number of elements
func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				count++
			}
		}
	}
	return count
}

// Remove removes x from the set
func (s *IntSet) Remove(x int) {
	if !s.Has(x) {
		fmt.Println("No such an element: ", x)
		return
	}

	word, bit := x/64, uint(x%64)

	// Take the opposite way of Add()
	// s.words[word] |= 1 << bit
	s.words[word] &^= 1 << bit
}

// Clear removes all elements from the set
func (s *IntSet) Clear() {
	// just assign a new slice uint64
	s.words = []uint64{}
}

// Copy is a copy of the set
func (s *IntSet) Copy() IntSet {
	w := make([]uint64, len(s.words))
	copy(w, s.words)
	return IntSet{w}
}

// AddAll allows a list of values to be added.
func (s *IntSet) AddAll(x ...int) {
	for _, v := range x {
		s.Add(v)
	}
}

// ex 6.3: See the sample in p. 54

// IntersectWith returns the intersection s ∩ x.
func (s *IntSet) IntersectWith(t *IntSet) {}

// DifferenceWith returns the difference s ∖ x.
func (s *IntSet) DifferenceWith() {}

// SymmetricDifference returns the symmetric difference s ∆ x.
func (s *IntSet) SymmetricDifference() {}

// ex 6.4

// Elems returns a slice containing the elements of the set.
func (s *IntSet) Elems() []int {
	var elems []int
	for i, word := range s.words {
		if word == 0 { // word is not set.
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, 64*i+j)
			}
		}
	}
	return elems
}
