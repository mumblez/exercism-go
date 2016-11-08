package stringset

import "strings"

const testVersion = 3

// the test for Add and Delete seems broken, definition says accept one argument but test supplies two!!!

// Set is a []string
type Set []string

// New returns a new set
func New() Set {
	return make(Set, 0)
}

// NewFromSlice returns a Set
func NewFromSlice(in []string) Set {
	if len(in) == 0 {
		return Set(in)
	}
	s := make(Set, 0)
	var duplicate bool
	for k, v := range in {
		for x, y := range in {
			if k != x && v == y && s.Has(v) {
				// skip duplicate
				duplicate = true
				break
			}
		}
		if !duplicate {
			s = append(s, v)
		}
		duplicate = false
	}
	return s
}

// Add ...
func (s *Set) Add(str string) {
	if !s.Has(str) {
		*s = append(*s, str)
	}
}

// Delete ...
func (s *Set) Delete(str string) {
	for i, v := range *s {
		if v == str {
			*s = append((*s)[:i], (*s)[i+1:]...)
			break
		}
	}
}

// Len ...
func (s Set) Len() int {
	return len(s)
}

// IsEmpty ...
func (s Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

// Slice ...
func (s Set) Slice() []string {
	return s
}

// String ...
func (s Set) String() string {
	if s.IsEmpty() {
		return "{}"
	}
	str := "{\""
	str += strings.Join(s, "\", \"")
	str += "\"}"
	return str
}

// Has ..
func (s Set) Has(str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// Equal ...
func Equal(s1, s2 Set) bool {
	if s1.IsEmpty() {
		if s2.IsEmpty() {
			return true
		}
		return false
	}
	if s2.IsEmpty() {
		return false
	}
	var counter int
	for _, v2 := range s2 {
		if s1.Has(v2) {
			counter++
		}
	}
	if counter != len(s1) {
		return false
	}
	return true
}

// Subset ...
func Subset(s1, s2 Set) bool { // return s1 ⊆ s2
	if s1.IsEmpty() {
		return true
	}
	if s2.IsEmpty() || len(s1) > len(s2) {
		return false
	}
	for _, v := range s1 {
		if !s2.Has(v) {
			return false
		}
	}
	return true
}

// Disjoint ...
func Disjoint(s1, s2 Set) bool {
	if s1.IsEmpty() || s2.IsEmpty() {
		return true
	}
	for _, v := range s1 {
		if s2.Has(v) {
			return false
		}
	}
	return true
}

// Intersection ...
func Intersection(s1, s2 Set) Set {
	iset := make(Set, 0)
	for _, v := range s1 {
		if s2.Has(v) {
			iset = append(iset, v)
		}
	}
	return iset
}

// Union ...
func Union(s1, s2 Set) Set {
	if s1.IsEmpty() && s2.IsEmpty() {
		return s1
	}
	if s1.IsEmpty() && !s2.IsEmpty() {
		return s2
	}
	if s2.IsEmpty() && !s1.IsEmpty() {
		return s1
	}
	unionSet := make(Set, len(s1))
	copy(unionSet, s1)
	for _, v := range s2 {
		if !unionSet.Has(v) {
			unionSet = append(unionSet, v)
		}
	}
	return unionSet
}

// Difference ...
func Difference(s1, s2 Set) Set { // return s1 ∖ s2
	if s1.IsEmpty() {
		return s1
	}
	if s2.IsEmpty() {
		return s1
	}
	diffset := make(Set, len(s1))
	copy(diffset, s1)
	for _, v := range s2 {
		if diffset.Has(v) {
			diffset.Delete(v)
		}
	}
	return diffset
}

// SymmetricDifference ...
func SymmetricDifference(s1, s2 Set) Set {
	if s1.IsEmpty() {
		return s2
	}
	if s2.IsEmpty() {
		return s1
	}
	sset := make(Set, len(s1))
	copy(sset, s1)
	for _, v := range s2 {
		if sset.Has(v) {
			sset.Delete(v)
			continue
		}
		sset.Add(v)
	}
	return sset
}
