package main

import  (
	"os"
	"fmt"
	"strconv"
	"bufio"
)


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
