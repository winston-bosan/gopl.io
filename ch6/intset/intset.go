// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

const intType int = 32

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []int
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/intType, uint(x%intType)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/intType, uint(x%intType)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
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

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < intType; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", intType*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string

// return the number of elements
func (s *IntSet) Len() (counter int) {
	for _, word := range s.words {
		if word == 0 { continue }
		for j := 0; j < intType; j++ {
			if word&(1<<j) != 0 {
				counter++
			}
		}
	}
	return counter
}

// remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/intType, uint(x%intType)
	s.words[word] &^= 1 << bit
}

// remove all elements from the set
func (s *IntSet) Clear() {
	x := make([]int, 1)
	s.words = x
}

// return a copy of the set
func (s *IntSet) Copy() *IntSet {
	x := func(s IntSet) IntSet { return s }(*s)
	return &x
}

// Exercise 6.2: Define a variadic (*IntSet).AddAll(...int) method 
// that allows a list of values to be added, such as s.AddAll(1, 2, 3).
func (s *IntSet) AddAll(ints ...int) {
	for i := 0; i < len(ints); i++ {
		s.Add(ints[i])
	}
}

// Exercise 6.3: IntersectWith
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}

// Exercise 6.3: DifferenceWith
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Exercise 6.3: SymmetricDifference
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Exercise 6.4: Add a method Elems that returns a slice 
// containing the elements of the set, suitable for iterating over with a range loop.
func (s *IntSet) Elems() []int {
	t := make([]int, 0)
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < intType; j++ {
			if word&(1<<uint(j)) != 0 {
				t = append(t, intType*i+j)
			}
		}
	}
	return t
}

// Exercise 6.5: The type of each word used by IntSet is int, 
// but intType-bit arithmetic may be inefficient on a 32-bit platform. 
// Modify the program to use the uint type, which is the most efficient 
// unsigned integer type for the platform. Instead of dividing by 64, 
// define a constant holding the effective size of uint in bits, 32 or 64. 
// You can use the perhaps too-clever expression 32 << (^uint(0) >> 63) for 
// this purpose.



func main() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"
	fmt.Println(x.Len()) // Pred: 3
	x.Remove(9)
	fmt.Println(x.String()) // "{1 144}"
	fmt.Println(x.Len()) // Pred: 2
	x.Clear()
	fmt.Println(x.String()) // "{}"
	fmt.Println(x.Len()) // Pred: 0

	y := x.Copy()
	y.AddAll(1,2,3,4,5,6,7,8)
	fmt.Println(y.String()) // "{1,2,3,4,5,6,7,8}"
	fmt.Println(y.Len()) // Pred: 8

	z := y.Copy()
	z.Remove(5)
	z.AddAll(231,43253456546,234,2,1,99)
	y.IntersectWith(z)
	fmt.Println(y.String()) // "{1,2,3,4,6,7,8}"
	fmt.Println(y.Len()) // Pred: 7

}

