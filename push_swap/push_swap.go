package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dupChecker() func(int) bool {
	duplicate := map[int]bool{}
	return func(num int) bool {
		if _, ok := duplicate[num]; !ok {
			duplicate[num] = true
			return false
		} else {
			return true
		}
	}
}

func stackFromArgNums(argv []string) stack {
	isDup := dupChecker()
	stackA := stack{}
	var smallest int
	for _, str := range argv {
		split := strings.Split(str, " ")
		for _, x := range split {
			num, err := strconv.Atoi(x)
			if num < smallest {
				smallest = num
			}
			if err != nil {
				fmt.Println("Bad input")
				os.Exit(2)
			}
			if isDup(num) {
				fmt.Println(num, "is duplicated, Error")
				os.Exit(2)
			}
			stackA = append(stackA, float64(num))
		}
	}
	return stackA
}

func main() {
	// functions := map[string]func(A, B *stack){
	// 	"sa": sa, "sb": sb, "ss": ss, "pa": pa, "pb": pb, "ra": ra, "rb": rb, "rr": rr, "rra": rra, "rrb": rrb, "rrr": rrr}

	if len(os.Args) == 1 {
		os.Exit(2)
	}
	stackA := stackFromArgNums(os.Args[1:])
	stackB := stack{}
	solve(&stackA, &stackB)
	// stackA.getPivot()
	// fmt.Println(stackA, stackB)
	// solve(stackA, stackB)
}

//get pivot function
//push or rotate, // double check if B can do it too // and swap as well if both works

func test(A, B *stack) {
	if len(*A) <= 5 {
		return
	}
	pivotA, _ := A.getPivot()
	pivotB, errB := B.getPivot()
	for (*A)[0] != pivotA {
		if (*A)[0] > pivotA && errB != nil && (*B)[0] > pivotB {
			rr(A, B)
		} else if (*A)[0] > pivotA {
			ra(A, B)
		} else {
			pb(A, B)
		}
	}
	test(A, B)
}

func solve(A, B *stack) {
	test(A, B)
}
