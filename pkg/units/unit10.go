package units

import (
	"fmt"
	"math"
)

func JoiningGroups() {
	s := NewSet()
	s.Insert(-25.4)
	s.Insert(-27.0)
	s.Insert(13.0)
	s.Insert(19.0)
	s.Insert(15.5)
	s.Insert(24.5)
	s.Insert(21.0)
	s.Insert(32.5)
	fmt.Println(s.GetAll())
	fmt.Println(s.GetSubsetFor(-10))

}

type measurement int64
type subsetId int

type Set struct {
	subsets map[subsetId]*Subset
}

type Subset struct {
	elements map[measurement]int
}

func NewSet() *Set {
	return &Set{
		subsets: map[subsetId]*Subset{},
	}
}

const roundFactor = 10

func toStorage(val float64) (subsetId, measurement) {
	m := measurement(math.Round(val * roundFactor))
	id := subsetId((int(val) / 10) * 10)
	return id, m
}

func fromStorage(m measurement, count int) []float64 {
	val := float64(m) / float64(roundFactor)
	var vals []float64
	for i := 0; i < count; i++ {
		vals = append(vals, val)
	}
	return vals
}

func (s *Set) Insert(val float64) {
	id, m := toStorage(val)
	if ss, found := s.subsets[id]; !found {
		s.subsets[id] = &Subset{
			elements: map[measurement]int{m: 1},
		}
	} else {
		ss.elements[m]++
	}
}

func (s *Set) DeleteOne(val float64) {
	id, m := toStorage(val)
	if ss, found := s.subsets[id]; found {
		ss.elements[m]--
	}
}

func (s *Set) DeleteAll(val float64) {
	id, m := toStorage(val)
	if ss, found := s.subsets[id]; found {
		delete(ss.elements, m)
	}
}

func (s *Set) Contains(val float64) bool {
	id, m := toStorage(val)
	if ss, found := s.subsets[id]; found {
		if c, found := ss.elements[m]; found && c > 0 {
			return true
		}
	}
	return false
}

func (s *Set) GetSubsetFor(val float64) []float64 {
	id, _ := toStorage(val)
	if ss, found := s.subsets[id]; found {
		var vals []float64
		for m, count := range ss.elements {
			vals = append(vals, fromStorage(m, count)...)
		}
		return vals
	}
	return nil
}

func (s *Set) GetAll() []float64 {
	var vals []float64
	for _, ss := range s.subsets {
		for m, count := range ss.elements {
			vals = append(vals, fromStorage(m, count)...)
		}
	}
	return vals
}
