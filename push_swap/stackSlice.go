package main

import (
	"errors"
	"sort"
)

type stack []float64

func prepend(s stack, n float64) stack {
	s = append(s, 0)
	copy(s[1:], s)
	s[0] = n
	return s
}

func (s *stack) swap() {
	if len(*s) >= 2 {
		var tmp float64
		tmp = (*s)[0]
		(*s)[0] = (*s)[1]
		(*s)[1] = tmp
	}
}

func (s *stack) pop() (float64, stack) {
	if len(*s) != 0 {
		x, new := (*s)[0], (*s)[1:]
		*s = new
		return x, new
	}
	return 0, nil
}

func (s *stack) popBot() (float64, stack) {
	if len(*s) != 0 {
		x, new := (*s)[len((*s))-1], (*s)[:len((*s))-1]
		*s = new
		return x, new
	}
	return 0, nil
}

func (s *stack) rotateDown() {
	num, new := (*s).popBot()
	if new != nil { //if new = nil means nothing was poped so nothing to prepend
		(*s) = prepend(new, num)
	}
}

func (s *stack) rotateUp() {
	num, new := (*s).pop()
	if new != nil {
		(*s) = append(new, num)
	}
}

func (s *stack) isSorted() bool {
	for i := (len(*s) - 1); i > 0; i-- {
		if (*s)[i] < (*s)[i-1] {
			return false
		}
	}
	return true
}

func (s *stack) isEmpty() bool {
	if len((*s)) == 0 {
		return true
	}
	return false
}

// func (s *stack) Peak() bool {
// 	if len((*s)) == 0 {
// 		return (*s)[0]
// 	}
// 	return false
// }

func (s *stack) getPivot() (float64, error) {
	var pivot float64
	if len(*s) == 0 {
		return 0, errors.New("stack B is empty")
	}
	tmp := make([]float64, len(*s))
	copy(tmp, *s)
	sort.Float64s(tmp)
	pivot = tmp[len(*s)/2]
	return pivot, nil
}
