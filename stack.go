package main

import (
	"bytes"
	"fmt"
)

type node struct {
	num int
	above *node
	below *node
}

type Stack struct {
	top *node
	bot *node
}


func (s *Stack) addNew(num int) {
	tmp := &node{num, nil, nil}
	s.push(tmp)
}

func swapNodes(first *node, second *node) {
	tmp := first.above
	first.above = second
	first.below = second.below
	second.above = tmp
	second.below = first
}

func (s *Stack) swap() {
	if s.top != nil && s.top.below != nil {
		swapNodes(s.top, s.top.below)
		s.top = s.top.above
	}
}
func (s *Stack) rotate() {
	if s.top != nil && s.top.below != nil {
		var tmp *node
		tmp = s.top
		s.top = s.top.below
		tmp.above = s.bot
		tmp.below = nil
		s.bot.below = tmp
		s.bot = tmp
	}
}

func (s *Stack) push(node *node) {
	if (node != nil) {
		if (s.top == nil) {
			s.top = node
			s.bot = node
		} else {
			node.below = s.top
			s.top.above = node
			s.top = node
		}
	}
}

func (s *Stack) pop() *node {
	if (s.top != nil) {
		tmp := s.top
		s.top = s.top.below
		tmp.above = nil
		tmp.below = nil
		return tmp
	}
	return nil
}

func (s Stack) String() string {
	buf := bytes.Buffer{}
	for snode := s.top ;snode != nil; snode = snode.below {
		buf.WriteString(fmt.Sprintf("%d - ", snode.num))		
	}
	return	buf.String()
}
