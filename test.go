package main

import  (
	"os"
	"fmt"
	"strconv"
	"bytes"
	// "strings"
	"bufio"
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


func scanner(stackA Stack) []string {
	commands := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		
		//Add caseswitch to check for input as it's being inputed
		commands = append(commands, input)	
	}
	return commands
  }

func	sa(A, B *Stack) {
	A.swap()
}


func visualizer(commands []string, A, B *Stack) {
	function := map[string]func (A, B *Stack) { "sa" : sa}
	for _, key := range commands {
		if fun, ok := function[key]; ok {
			fun(A, B)
			fmt.Println(A)
		} else {
			//quit here?
			fmt.Println("that's not a function I can use") // ?? I dont need to check because alll my indexs are already okay and i quit early before
		}
		//wait someamount of time to see based on size of input?
		// update(stackA, stackB)
	}
}

func main() {
	argv := os.Args[1:]
	stackA := Stack{}
	stackB := Stack{}
	for _, n := range argv {
		num, err := strconv.Atoi(n)
		if err != nil {
			fmt.Println("Bad input")
			os.Exit(0)
		}
		stackA.addNew(num)
	}
	commands := scanner(stackA)
	visualizer(commands, &stackA, &stackB)
	fmt.Println(commands)
	// if -V pass comands to visualizer
	fmt.Println(stackA)
	fmt.Println(stackB)
}

/* 
sa
sb
ss : sa and sb at the same time.
pa : push a - take the first element at the top of b and put it at the top of a. Do
nothing if b is empty.
pb : push b - take the first element at the top of a and put it at the top of b. Do
nothing if a is empty.
ra : rotate a - shift up all elements of stack a by 1. The first element becomes
the last one.
rb : rotate b - shift up all elements of stack b by 1. The first element becomes
the last one.
rr : ra and rb at the same time.
rra : reverse rotate a - shift down all elements of stack a by 1. The flast element
becomes the first one.
8
Push_swap Because Swap_push isn’t as natural
rrb : reverse rotate b - shift down all elements of stack b by 1. The flast element
becomes the first one.
rrr : rra and rrb at the same time.
*/
