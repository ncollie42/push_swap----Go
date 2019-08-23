package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func scanner(stackA, stackB Stack) []string {
	commands := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	function := map[string]func(A, B *Stack){
		"sa": sa, "sb": sb, "ss": ss, "pa": pa, "pb": pb, "ra": ra, "rb": rb, "rr": rr, "rra": rra, "rrb": rrb, "rrr": rrr}
	for scanner.Scan() {
		input := scanner.Text()
		if fun, ok := function[input]; ok {
			fun(&stackA, &stackB)
			fmt.Print("A: ")
			fmt.Println(stackA)
			fmt.Print("B: ")
			fmt.Println(stackB)
		} else {
			//quit here?
			fmt.Println("that's not a function I can use") // ?? I dont need to check because alll my indexs are already okay and i quit early before
		}
		commands = append(commands, input)
	}
	return commands
}

func visualizer(commands []string, A, B *Stack) {
	function := map[string]func(A, B *Stack){
		"sa": sa, "sb": sb, "ss": ss, "pa": pa, "pb": pb, "ra": ra, "rb": rb, "rr": rr, "rra": rra, "rrb": rrb, "rrr": rrr}
	for _, key := range commands {
		if fun, ok := function[key]; ok {
			fun(A, B)
			fmt.Print("A: ")
			fmt.Println(A)
			fmt.Print("B: ")
			fmt.Println(B)
		} else {
			//quit here?
			fmt.Println("that's not a function I can use") // ?? I dont need to check because alll my indexs are already okay and i quit early before
		}
		//wait someamount of time to see based on size of input?
		// update(stackA, stackB)
	}
}

func main() {
	runGocui()

	argv := os.Args[1:]
	stackA := Stack{}
	stackB := Stack{}
	duplicate := map[int]bool{}
	for _, n := range argv { // First argument should be the top thing on the stack, need to update this
		num, err := strconv.Atoi(n)
		if err != nil {
			fmt.Println("Bad input")
			os.Exit(0)
		}
		if _, ok := duplicate[num]; !ok {
			duplicate[num] = true
		} else {
			fmt.Println("Duplicate number", num)
			os.Exit(0)
		}
		stackA.addNew(num)
	}
	commands := scanner(stackA, stackB)
	// visualizer(commands, &stackA, &stackB)
	fmt.Println(commands)
	// if -V pass comands to visualizer
}

/*
	Todo:
		reverse input order
		checker Ok Ko
*/
